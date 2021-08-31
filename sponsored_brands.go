package amazon_advertising

import "fmt"

func (c *Client) SbListBrands(query ...interface{}) (*Response, error) {
	return c.HttpGet("/brands", query...)
}

func (c *Client) SbListCampaigns(query ...interface{}) (*Response, error) {
	return c.HttpGet("/sb/campaigns", query...)
}

func (c *Client) SbListDrafts(query ...interface{}) (*Response, error) {
	return c.HttpGet("/sb/drafts/campaigns", query...)
}

func (c *Client) SbListAdGroups(query ...interface{}) (*Response, error) {
	return c.HttpGet("/sb/adGroups", query...)
}

func (c *Client) SbListKeywords(query ...interface{}) (*Response, error) {
	return c.HttpGet("/sb/keywords", query...)
}

func (c *Client) SbListNegativeKeywords(query ...interface{}) (*Response, error) {
	return c.HttpGet("/sb/negativeKeywords", query...)
}

func (c *Client) SbListProductTargets(json interface{}) (*Response, error) {
	return c.HttpPost("/sb/targets/list", json)
}

func (c *Client) SbListNegativeProductTargets(json interface{}) (*Response, error) {
	return c.HttpPost("/sb/negativeTargets/list", json)
}

func (c *Client) SbRequestReport(recordType, params interface{}) (*Response, error) {
	url := fmt.Sprintf("/v2/hsa/%s/report", recordType)
	return c.HttpPost(url, params)
}