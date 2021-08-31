// profile api
package amazon_advertising

func (c *Client) GetProfiles() (*Response, error) {
	return c.HttpGet("/v2/profiles")
}
