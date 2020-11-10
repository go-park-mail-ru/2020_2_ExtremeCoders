package UserDelivery

import (
	"CleanArch/internal/User/UserRepository"
	"CleanArch/internal/User/UserUseCase"
	"CleanArch/internal/errors"
	"net/http"
)

func SignUpError(err error, cookie *http.Cookie) []byte{
	switch err {
	case nil:
		return errors.GetOkAns(cookie.Value)
	case UserRepository.EmailAlreadyExists:
		return errors.GetErrorLoginExistAns()
	case UserRepository.CantAddSession:
		return errors.GetAddSessionError()
	case UserRepository.CantAddUser:
		return errors.AddUserError()
	}
	return nil
}

func SignInError(err error, cookie *http.Cookie) []byte{
	switch err {
	case nil:
		return errors.GetOkAns(cookie.Value)
	case UserRepository.CantAddSession:
		return errors.GetAddSessionError()
	case UserRepository.CantGetUserByEmail:
		return errors.GetErrorNoUserAns()
	case UserRepository.GetSessionError:
		errors.GetErrorNoCockyAns()
	case UserUseCase.WrongPasswordError:
		errors.GetErrorBadPasswordAns()
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

func LogoutError(err error)[]byte{
	switch err {
	case nil:
		return errors.GetOkAns("")
	case UserRepository.InvalidSession:
		return errors.GetErrorNoCockyAns()
	case UserRepository.RemoveSessionError:
		return errors.RemoveSessionError()
	}
	return nil
}

func ProfileError(err error, cookie *http.Cookie) []byte{
	switch err {
	case nil:
		return errors.GetOkAns(cookie.Value)
	case UserRepository.CantUpdateUser:
		return errors.UpdateProfileError()
	case UserRepository.CantGetUserOnUpdate:
		return errors.GetUserOnUpdateError()
	}
	return nil
}