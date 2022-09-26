package smartermailapi

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type ListDomainNamesResponse struct {
	Data    []string `json:"data"`
	Success bool     `json:"success"`
	Message string   `json:"message"`
}

func (c *Client) AdminListDomainNames() (*ListDomainNamesResponse, error) {
	req := c.includeAccessToken(c.defaultGetRequest("api/v1/settings/sysadmin/domain-names"))

	var htttpResponse fasthttp.Response
	fasthttp.Do(req, &htttpResponse)

	if err := HandleError(&htttpResponse); err != nil {
		return nil, err
	}

	var res ListDomainNamesResponse
	json.Unmarshal(htttpResponse.Body(), &res)

	return &res, nil
}
