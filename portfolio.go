// list Portfolios
package amazon_advertising

func (c *Client) ListPortfoliosEx(query ...interface{}) (*Response, error) {
	return c.HttpGet("/v2/portfolios/extended", query...)
}
