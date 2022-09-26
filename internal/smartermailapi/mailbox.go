package smartermailapi

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

type GetMailboxForwardListResponse struct {
	List struct {
		SpamForwardOption string   `json:"spamForwardOption"`
		ForwardList       []string `json:"forwardList"`
		KeepRecipients    bool     `json:"keepRecipients"`
		DeleteOnForward   bool     `json:"deleteOnForward"`
	} `json:"mailboxForwardList"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (c *Client) UserGetMailboxForwardList() (*GetMailboxForwardListResponse, error) {
	req := c.includeAccessToken(c.defaultGetRequest("api/v1/settings/mailbox-forward-list"))

	var httpRes fasthttp.Response
	fasthttp.Do(req, &httpRes)

	if err := HandleError(&httpRes); err != nil {
		return nil, err
	}

	var res GetMailboxForwardListResponse
	json.Unmarshal(httpRes.Body(), &res)

	return &res, nil
}

type AdminGetMailboxForwardListResponse = []AdminForwardItem

type AdminForwardItem struct {
	Email      string
	ForwardsTo []string
	Success    bool
}

func (c *Client) AdminGetMailboxForwardList(domain string) (*AdminGetMailboxForwardListResponse, error) {

	data, err := c.AdminListUsernames(domain)

	if err != nil {
		return nil, err
	}

	listResponse := AdminGetMailboxForwardListResponse{}

	for idx, username := range data.Usernames {
		email := fmt.Sprintf("%s@%s", username, domain)

		listResponse = append(listResponse, AdminForwardItem{
			Email:      email,
			ForwardsTo: []string{},
			Success:    false,
		})

		user, err := c.AdminGetUser(email)

		if err != nil {
			listResponse[idx].Success = false
			continue
		}

		hasTempPassword := false
		tempPassword := user.UserData.TempPassword

		if user.UserData.TempPassword != "" {
			hasTempPassword = true
		}

		if !hasTempPassword {
			tp, err := c.CreateTempPassword(email)

			if err != nil {
				listResponse[idx].Success = false
				continue
			}

			defer c.RevokeTempPassword(email)

			tempPassword = tp.TempPassword
		}

		userClient := NewClient(c.host, username, tempPassword)

		_, err = userClient.Authenticate()

		if err != nil {
			listResponse[idx].Success = false
			continue
		}

		list, err := userClient.UserGetMailboxForwardList()

		if err != nil {
			listResponse[idx].Success = false
			continue
		}

		listResponse[idx].ForwardsTo = list.List.ForwardList
	}

	return &listResponse, nil
}
