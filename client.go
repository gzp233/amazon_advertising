// base client
package amazon_advertising

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/guonaihong/gout/filter"

	"github.com/guonaihong/gout/dataflow"

	"github.com/tidwall/gjson"

	"github.com/guonaihong/gout"
)

const (
	apiTokenUrl = "https://api.amazon.com/auth/o2/token"
	MAX_RETRIES = 3
)

type Response struct {
	StatusCode int
	Body       string
	Header http.Header
}

type ClientOption struct {
	Region       string
	ClientId     string
	ClientSecret string
	RefreshToken string
	AccessToken  string
}

// 请求客户端
type Client struct {
	Debug     bool
	Option    ClientOption
	Endpoint  string
	ProfileId int64
}

func (c *Client) getUri(url string) string {
	return fmt.Sprintf("%s%s", c.Endpoint, url)
}

func (c *Client) getHeader() gout.H {
	headers := gout.H{
		"Authorization":                   fmt.Sprintf("bearer %s", c.Option.AccessToken),
		"Amazon-Advertising-API-ClientId": c.Option.ClientId,
	}
	if c.ProfileId != 0 {
		headers["Amazon-Advertising-API-Scope"] = c.ProfileId
	}

	return headers
}

func (c *Client) SetDebug(mode bool) {
	c.Debug = mode
}

func (c *Client) SetProfileId(profileId int64) {
	c.ProfileId = profileId
}

func (c *Client) HttpGet(url string, query ...interface{}) (*Response, error) {
	return c.handleRequest(gout.GET(c.getUri(url)).SetQuery(query...))
}

func (c *Client) HttpPost(url string, params interface{}) (*Response, error) {
	if params == nil {
		params = gout.H{}
	}
	return c.handleRequest(gout.POST(c.getUri(url)).SetJSON(params))
}

func (c *Client) GetReport(reportId string) (*Response, error) {
	url := fmt.Sprintf("/v2/reports/%s", reportId)
	return c.HttpGet(url)
}

func (c *Client) DownloadReportData(reportId string) (*Response, error) {
	r := &Response{}
	var bs []byte
	url := fmt.Sprintf("/v2/reports/%s/download", reportId)
	err := gout.GET(c.getUri(url)).Debug(c.Debug).SetHeader(c.getHeader()).
		BindHeader(&r.Header).BindBody(&bs).Code(&r.StatusCode).Do()
	if err != nil {
		return nil, err
	}
	if r.StatusCode >= 400 {
		return r, fmt.Errorf("request error:%s", bs)
	}

	// 解析gz二进制数据
	b := new(bytes.Buffer)
	err = binary.Write(b, binary.BigEndian, bs)
	if err != nil {
		return nil, fmt.Errorf("write binary error:%s", err)
	}
	g, err := gzip.NewReader(b)
	if err != nil {
		return nil, fmt.Errorf("open gzip file error:%s", err)
	}
	defer g.Close()
	d, err := ioutil.ReadAll(g)
	if err != nil {
		return nil, fmt.Errorf("read gzip file error:%s", err)
	}
	r.Body = string(d)

	return c.handleResponse(r, err)
}

// 处理请求,增加重试条件
func (c *Client) handleRequest(d *dataflow.DataFlow) (*Response, error) {
	r := &Response{}
	err := d.Debug(c.Debug).SetHeader(c.getHeader()).BindHeader(&r.Header).BindBody(&r.Body).Code(&r.StatusCode).F().
		Retry().Attempt(MAX_RETRIES).WaitTime(2 * time.Second).MaxWaitTime(5 * time.Second).
		Func(func(ctx *dataflow.Context) error {
			if ctx.Error != nil {
				if ctx.Code == 401 {
					refreshErr := c.RefreshToken()
					if refreshErr != nil {
						return nil
					}
				}
				if ctx.Code == http.StatusTooManyRequests || ctx.Code >= 500 {
					return filter.ErrRetry
				}
			}

			return nil
		}).Do()
	return c.handleResponse(r, err)
}

// 处理响应
func (c *Client) handleResponse(r *Response, err error) (*Response, error) {
	if err != nil {
		return nil, err
	}
	if r.StatusCode >= 400 {
		return r, errors.New(r.Body)
	}

	return r, nil
}

// 获取accessToken
func (c *Client) RefreshToken() error {
	headers := gout.H{
		"User-Agent": "AdvertisingAPI Golang Client Library v1.0",
	}
	params := gout.H{
		"grant_type":    "refresh_token",
		"refresh_token": c.Option.RefreshToken,
		"client_id":     c.Option.ClientId,
		"client_secret": c.Option.ClientSecret,
	}
	resp := ""
	err := gout.POST(apiTokenUrl).SetHeader(headers).SetWWWForm(params).Debug(c.Debug).BindBody(&resp).Do()
	if err != nil {
		return err
	}

	c.Option.AccessToken = gjson.Get(resp, "access_token").String()
	if c.Option.AccessToken == "" {
		return errors.New("get access_token failed")
	}

	return nil
}

func NewClient(option ClientOption) (*Client, error) {
	c := &Client{
		Option: option,
	}
	if option.Region == "NA" {
		c.Endpoint = "https://advertising-api.amazon.com"
	} else if option.Region == "FE" {
		c.Endpoint = "https://advertising-api-fe.amazon.com"
	} else if option.Region == "EU" {
		c.Endpoint = "https://advertising-api-eu.amazon.com"
	} else {
		return nil, errors.New("the region is not correct")
	}
	if option.AccessToken == "" {
		err := c.RefreshToken()
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}
