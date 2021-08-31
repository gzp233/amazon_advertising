// Sponsored products api
package amazon_advertising

import (
	"fmt"
)

func (c *Client) SpListCampaignsEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/v2/sp/campaigns/extended", query...)
}

func (c *Client) SpListCampaignNegativeKeywordsEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/v2/sp/campaignNegativeKeywords/extended", query...)
}

func (c *Client) SpListAdGroupsEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/v2/sp/adGroups/extended", query...)
}

func (c *Client) SpListBiddableKeywordsEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/v2/sp/keywords/extended", query...)
}

func (c *Client) SpListNegativeKeywordsEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/v2/sp/negativeKeywords/extended", query...)
}

func (c *Client) SpListProductAdsEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/v2/sp/productAds/extended", query...)
}

func (c *Client) SpListTargetingClausesEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/v2/sp/targets/extended", query...)
}

func (c *Client) SpListNegativeTargetingClausesEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/v2/sp/negativeTargets/extended", query...)
}

func (c *Client) SpRequestReport(recordType, params interface{}) (*Response, error) {
	url := fmt.Sprintf("/v2/sp/%s/report", recordType)
	return c.HttpPost(url, params)
}