package lotus

type PingResponse struct {
	OrganizationId string `json:"organization_id"`
}

// Ping pings Lotus API to check if the API key is valid
// See: https://docs.uselotus.io/api-reference/api-overview
func (c *Client) Ping() (resp *PingResponse, err error) {
	resp = new(PingResponse)
	err = c.get("/api/ping/", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
