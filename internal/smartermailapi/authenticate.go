package smartermailapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

// api/v1/auth/authenticate-user
type AuthenticateBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthenticateResponse struct {
	Email                 string    `json:"emailAddress"`
	ChangePasswordNeeded  bool      `json:"changePasswordNeeded"`
	DisplayWelcomeWizard  bool      `json:"displayWelcomeWizard"`
	IsAdmin               bool      `json:"isAdmin"`
	IsDomainAdmin         bool      `json:"isDomainAdmin"`
	IsLicensed            bool      `json:"isLicensed"`
	AutoLoginToken        string    `json:"autoLoginToken"`
	AutoLoginURL          string    `json:"autoLoginUrl"`
	LocaleId              string    `json:"localeId"`
	IsImpersonating       bool      `json:"isImpersonating"`
	CanViewPasswords      bool      `json:"canViewPasswords"`
	AccessToken           string    `json:"accessToken"`
	RefreshToken          string    `json:"refreshToken"`
	AccessTokenExpiration time.Time `json:"accessTokenExpiration"`
	Username              string    `json:"username"`
	Success               bool      `json:"success"`
	ResultCode            int       `json:"resultCode"`
	Message               string    `json:"message"`
	DebugInfo             string    `json:"debugInfo"`
}

func (c *Client) Authenticate() (*AuthenticateResponse, error) {
	req := c.defaultPostRequest("api/v1/auth/authenticate-user")

	body, _ := json.Marshal(AuthenticateBody{
		Username: c.username,
		Password: c.password,
	})

	req.SetBody(body)

	var httpResponse fasthttp.Response
	fasthttp.Do(req, &httpResponse)

	if code := httpResponse.StatusCode(); code != 200 {
		if code == 400 {
			return nil, errors.New("")
		}

		return nil, errors.New("err")
	}

	var res AuthenticateResponse
	json.Unmarshal(httpResponse.Body(), &res)

	c.refreshToken = res.RefreshToken
	c.accessToken = res.AccessToken

	return &res, nil
}

// api/v1/auth/refresh-token
type RefreshTokenBody struct {
	Token string `json:"token"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (c *Client) RefreshToken() (*RefreshTokenResponse, error) {
	req := c.defaultPostRequest("api/v1/auth/refresh-token")
	body, _ := json.Marshal(RefreshTokenBody{
		Token: c.refreshToken,
	})

	req.SetBody(body)

	var httpResponse fasthttp.Response
	fasthttp.Do(req, &httpResponse)

	if err := HandleError(&httpResponse); err != nil {
		return nil, err
	}

	var res RefreshTokenResponse
	json.Unmarshal(httpResponse.Body(), &res)

	c.refreshToken = res.RefreshToken
	c.accessToken = res.AccessToken

	return &res, nil

}

// api/v1/auth/generate-temp-pass/
type GenerateTempPassResponse struct {
	TempPassword       string     `json:"tempPassword"`
	TempPasswordExpUtc time.Timer `json:"tempPasswordExpUtc"`
	Success            bool       `json:"success"`
	ResultCode         int        `json:"resultCode"`
	Message            string     `json:"message"`
	DebugInfo          string     `json:"debugInfo"`
}

func (c *Client) CreateTempPassword(email string) (*GenerateTempPassResponse, error) {
	req := c.defaultPostRequest(fmt.Sprintf("api/v1/auth/generate-temp-pass/%s", email))
	req = c.includeAccessToken(req)

	var httpResponse fasthttp.Response
	fasthttp.Do(req, &httpResponse)

	if err := HandleError(&httpResponse); err != nil {
		return nil, err
	}

	var res GenerateTempPassResponse
	json.Unmarshal(httpResponse.Body(), &res)

	return &res, nil
}

type RevokeTempPassResponse struct {
	TempPassword       string     `json:"tempPassword"`
	TempPasswordExpUtc time.Timer `json:"tempPasswordExpUtc"`
	Success            bool       `json:"success"`
	ResultCode         int        `json:"resultCode"`
	Message            string     `json:"message"`
	DebugInfo          string     `json:"debugInfo"`
}

// api/v1/auth/revoke-temp-pass
func (c *Client) RevokeTempPassword(email string) (*RevokeTempPassResponse, error) {
	req := c.includeAccessToken(c.defaultPostRequest(fmt.Sprintf("api/v1/auth/revoke-temp-pass/%s", email)))

	var httpRes fasthttp.Response
	fasthttp.Do(req, &httpRes)

	if err := HandleError(&httpRes); err != nil {
		return nil, err
	}

	var res RevokeTempPassResponse
	json.Unmarshal(httpRes.Body(), &res)

	return &res, nil
}
