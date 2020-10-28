package Delivery

import "net/http"

func SignUpError(code uint16, cookie *http.Cookie) []byte{
	switch code {
	case 200:
		return getOkAns(cookie.Value)
	case 401:
		return getErrorLoginExistAns()
	}
	return nil
}

func SignInError(code uint16, cookie *http.Cookie) []byte{
	switch code {
	case 200:
		return getOkAns(cookie.Value)
	case 401:
		return getErrorBadPasswordAns()
	case 404:
		return getErrorNoUserAns()
	}
	return nil
}

func CookieError(code uint16) []byte{
	switch code {
	case 401:
		return getErrorNoCockyAns()
	case 402:
		return getErrorWrongCookieAns()
	}
	return nil
}