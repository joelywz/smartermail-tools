package smartermailapi

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

type ListUsernamesResponse struct {
	Usernames []string `json:"usernames"`
	Success   bool     `json:"success"`
	Message   string   `json:"message"`
}

func (c *Client) AdminListUsernames(domain string) (*ListUsernamesResponse, error) {
	req := c.includeAccessToken(
		c.defaultGetRequest(
			fmt.Sprintf("api/v1/settings/sysadmin/list-usernames/%s", domain),
		),
	)

	var httpResponse fasthttp.Response
	fasthttp.Do(req, &httpResponse)

	if err := HandleError(&httpResponse); err != nil {
		return nil, err
	}

	var res ListUsernamesResponse
	json.Unmarshal(httpResponse.Body(), &res)

	return &res, nil
}

// api/v1/settings/sysadmin/user
type GetUserResponse struct {
	UserData struct {
		AdminID            string `json:"adminId"`
		CreateDate         string `json:"createDate"`
		Description        string `json:"description"`
		Domain             string `json:"domain"`
		EnableSMTPAccounts bool   `json:"enableSmtpAccounts"`
		FullName           string `json:"fullName"`
		IsListModerator    bool   `json:"isListModerator"`
		// LastAcceptedPolicyVersion interface{} `json:"lastAcceptedPolicyVersion"`
		Loaded             string    `json:"loaded"`
		TempPassword       string    `json:"tempPassword"`
		TempPasswordExpUtc time.Time `json:"tempPasswordExpUtc"`
		PasswordLocked     bool      `json:"passwordLocked"`
		Permissions        int       `json:"permissions"`
		SecurityFlags      struct {
			AuthType                    int    `json:"authType"`
			AuthenticatingWindowsDomain string `json:"authenticatingWindowsDomain"`
			IsDisabled                  bool   `json:"isDisabled"`
			IsPrimarySystemAdmin        bool   `json:"isPrimarySystemAdmin"`
			IsDomainAdmin               bool   `json:"isDomainAdmin"`
			IsSystemAdmin               bool   `json:"isSystemAdmin"`
		} `json:"securityFlags"`
		SessionStr          string      `json:"sessionStr"`
		CsvFields           interface{} `json:"csvFields"`
		IsPasswordExpired   bool        `json:"isPasswordExpired"`
		IsPasswordCompliant int         `json:"isPasswordCompliant"`
		UserName            string      `json:"userName"`
		AdUsername          string      `json:"adUsername"`
		MaxMailboxSize      int         `json:"maxMailboxSize"`
		Roles               []string    `json:"roles"`
		EmailAddress        string      `json:"emailAddress"`
		CurrentMailboxSize  int         `json:"currentMailboxSize"`
	} `json:"userData"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (c *Client) AdminGetUser(email string) (*GetUserResponse, error) {
	req := c.includeAccessToken(c.defaultGetRequest(fmt.Sprintf("api/v1/settings/sysadmin/user/%s", email)))

	var httpRes fasthttp.Response
	fasthttp.Do(req, &httpRes)

	if err := HandleError(&httpRes); err != nil {
		return nil, err
	}

	var res GetUserResponse
	json.Unmarshal(httpRes.Body(), &res)

	return &res, nil

}
