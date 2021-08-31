package amazon_advertising

import "fmt"

func (c *Client) SdListCampaignsEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/sd/campaigns/extended", query...)
}

func (c *Client) SdListAdGroupsEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/sd/adGroups/extended", query...)
}

func (c *Client) SdListProductAdsEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/sd/productAds/extended", query...)
}

func (c *Client) SdListTargetsEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/sd/targets/extended", query...)
}

func (c *Client) SdListNegativeTargetsEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/sd/negativeTargets/extended", query...)
}

func (c *Client) SdRequestReport(recordType, params interface{}) (*Response, error) {
	url := fmt.Sprintf("/sd/%s/report", recordType)
	return c.HttpPost(url, params)
}