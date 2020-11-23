package UserDelivery

import (
	"Mailer/MainApplication/internal/User/UserRepository"
	"Mailer/MainApplication/internal/User/UserUseCase"
	"Mailer/MainApplication/internal/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func SignUpError(err error, cookie *http.Cookie) []byte {
	var returned []byte
	defer func() {
		log.WithFields(log.Fields{
			"RESPONSE": string(returned),
		}).Info("sent")
	}()
	switch err {
	case nil:
		returned= errors.GetOkAns(cookie.Value)
	case UserRepository.EmailAlreadyExists:
		returned= errors.GetErrorLoginExistAns()
	case UserRepository.CantAddSession:
		returned= errors.GetAddSessionError()
	case UserRepository.CantAddUser:
		returned= errors.AddUserError()
	}
	return returned
}

func SignInError(err error, cookie *http.Cookie) []byte {
	var returned []byte
	defer func() {
		log.WithFields(log.Fields{
			"RESPONSE": string(returned),
		}).Info("sent")
	}()
	switch err {
	case nil:
		returned= errors.GetOkAns(cookie.Value)
	case UserRepository.CantAddSession:
		returned= errors.GetAddSessionError()
	case UserRepository.CantGetUserByEmail:
		returned= errors.GetErrorNoUserAns()
	case UserRepository.GetSessionError:
		returned=errors.GetErrorNoCockyAns()
	case UserUseCase.WrongPasswordError:
		returned=errors.GetErrorBadPasswordAns()
	case UserRepository.RemoveSessionError:
		returned=errors.RemoveSessionError()
	}
	return returned
}

func CookieError(code uint16) []byte {
	var returned []byte
	defer func() {
		log.WithFields(log.Fields{
			"RESPONSE": string(returned),
		}).Info("sent")
	}()
	switch code {
	case 401:
		returned= errors.GetErrorNoCockyAns()
	case 402:
		returned= errors.GetErrorWrongCookieAns()
	}
	return returned
}

func LogoutError(err error) []byte {
	var returned []byte
	defer func() {
		log.WithFields(log.Fields{
			"RESPONSE": string(returned),
		}).Info("sent")
	}()
	switch err {
	case nil:
		returned= errors.GetOkAns("")
	case UserRepository.InvalidSession:
		returned= errors.GetErrorNoCockyAns()
	case UserRepository.RemoveSessionError:
		returned= errors.RemoveSessionError()
	}
	return returned
}

func ProfileError(err error, cookie *http.Cookie) []byte {
	var returned []byte
	defer func() {
		log.WithFields(log.Fields{
			"RESPONSE": string(returned),
		}).Info("sent")
	}()
	switch err {
	case nil:
		returned= errors.GetOkAns(cookie.Value)
	case UserRepository.CantUpdateUser:
		returned= errors.UpdateProfileError()
	case UserRepository.CantGetUserOnUpdate:
		returned= errors.GetUserOnUpdateError()
	}
	return returned
}
