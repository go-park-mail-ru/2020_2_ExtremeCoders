package UserDelivery

import (
	"CleanArch/cmd/errors"
	"net/http"
)

func SignUpError(code uint16, cookie *http.Cookie) []byte{
	switch code {
	case 200:
		return errors.GetOkAns(cookie.Value)
	case 401:
		return errors.GetErrorLoginExistAns()
	}
	return nil
}

func SignInError(code uint16, cookie *http.Cookie) []byte{
	switch code {
	case 200:
		return errors.GetOkAns(cookie.Value)
	case 401:
		return errors.GetErrorBadPasswordAns()
	case 404:
		return errors.GetErrorNoUserAns()
	}
	return nil
}

func CookieError(code uint16) []byte{
	switch code {
	case 401:
		return errors.GetErrorNoCockyAns()
	case 402:
		return errors.GetErrorWrongCookieAns()
	}
	return nil
}