// Sponsored products api
package amazon_advertising

import (
	"fmt"
)

func (c *Client) SpListCampaignsEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/sp/campaigns/extended", query...)
}

func (c *Client) SpListCampaignNegativeKeywordsEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/sp/campaignNegativeKeywords/extended", query...)
}

func (c *Client) SpListAdGroupsEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/sp/adGroups/extended", query...)
}

func (c *Client) SpListBiddableKeywordsEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/sp/keywords/extended", query...)
}

func (c *Client) SpListNegativeKeywordsEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/sp/negativeKeywords/extended", query...)
}

func (c *Client) SpListProductAdsEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/sp/productAds/extended", query...)
}

func (c *Client) SpListTargetingClausesEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/sp/targets/extended", query...)
}

func (c *Client) SpListNegativeTargetingClausesEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/sp/negativeTargets/extended", query...)
}

func (c *Client) SpRequestReport(recordType, params interface{}) (*Response, error) {
	url := fmt.Sprintf("/sp/%s/report", recordType)
	return c.HttpPost(url, params)
}

func (c *Client) SpGetReport(reportId string) (*Response, error) {
	url := fmt.Sprintf("/reports/%s", reportId)
	return c.HttpGet(url)
}

func (c *Client) SpDownloadReportData(location string) (*Response, error) {
	return c.HttpDownload(location)
}
