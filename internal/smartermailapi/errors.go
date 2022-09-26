package smartermailapi

import (
	"errors"
	"fmt"

	"github.com/valyala/fasthttp"
)

var (
	ErrInvalidInput         = errors.New("invalid input parameters")
	ErrTooManyRefresh       = errors.New("too many refresh attempts")
	ErrRefreshTokenRejected = errors.New("refresh token rejected")
	ErrIpRestricted         = errors.New("ip restriction against a system administrator")
	ErrUserNotFound         = errors.New("user not found")
	ErrCommunicationFailed  = errors.New("communication error")
	ErrTokenExpired         = errors.New("token has expired")
	ErrUnhandled            = errors.New("unhandled error")

	ErrTooManyLoginAttempts  = errors.New("too many login attempts")
	ErrEmptyUsernamePassword = errors.New("empty username or password")
	ErrPasswordRequirements  = errors.New("user does not meet password requirements")

	ErrUnauthorized = errors.New("unauthorized")
)

func HandleError(httpResponse *fasthttp.Response) error {
	msg := string(httpResponse.Header.StatusMessage())

	if code := httpResponse.StatusCode(); code != 200 {
		switch code {
		case 400:
			switch msg {
			case "Invalid input parameters":
				return ErrInvalidInput
			case "User not found":
				return ErrUserNotFound
			case "User does not meet password requirements":
				return ErrPasswordRequirements
			}
		case 401:

			switch msg {
			case "Unauthorized":
				return ErrUnauthorized
			case "Too many refresh attempts":
				return ErrTooManyRefresh
			case "Refresh token rejected":
				return ErrRefreshTokenRejected
			case "IP restriction against a system administrator":
				return ErrIpRestricted
			case "Token has expired":
				return ErrTokenExpired
			case "Too many login attempts":
				return ErrTooManyLoginAttempts
			case "Username or password is empty":
				return ErrEmptyUsernamePassword
			}
		case 500:
			return ErrCommunicationFailed
		}

		return fmt.Errorf("unhandled error: %d %s", code, msg)
	}

	return nil
}
