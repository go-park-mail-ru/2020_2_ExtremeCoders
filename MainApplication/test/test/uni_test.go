package test

import (
	"2020_2_ExtremeCoders/internal/Letter/LetterDelivery"
	"CleanArch/internal/Letter/LetterModel"
	"CleanArch/internal/Letter/LetterRepository"
	"CleanArch/internal/User/UserDelivery"
	"CleanArch/internal/User/UserModel"
	"CleanArch/internal/User/UserRepository"
	"CleanArch/internal/User/UserRepository/UserPostgres"
	"CleanArch/internal/User/UserUseCase"
	"CleanArch/internal/errors"
	"CleanArch/internal/pkg/context"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"testing"
)

func TestPasswordBcrypt(t *testing.T) {
	raw := "1538"
	enc := UserPostgres.PasswordBcrypt([]byte(raw))
	if bcrypt.CompareHashAndPassword(enc, []byte(raw)) != nil {
		log.Fatalf("error in PasswordBcrypt")
	}
}

func TestGenerateCSRF(t *testing.T) {
	csrf := context.GenerateCSRF()
	if csrf == "" {
		log.Fatalf("error in GenerateCSRF")
	}
}

func TestGetUserOnUpdateError(t *testing.T) {
	err := errors.GetUserOnUpdateError()
	if len(err) == 0 {
		log.Fatalf("error in GetUserOnUpdateError")
	}
}

func TestUpdateProfileError(t *testing.T) {
	err := errors.UpdateProfileError()
	if len(err) == 0 {
		log.Fatalf("error in UpdateProfileError")
	}
}

func TestAddUserError(t *testing.T) {
	err := errors.AddUserError()
	if len(err) == 0 {
		log.Fatalf("error in AddUserError")
	}
}

func TestRemoveSessionError(t *testing.T) {
	err := errors.RemoveSessionError()
	if len(err) == 0 {
		log.Fatalf("error in RemoveSessionError")
	}
}

func TestGetAddSessionError(t *testing.T) {
	err := errors.GetAddSessionError()
	if len(err) == 0 {
		log.Fatalf("error in GetAddSessionError")
	}
}

func TestGetOkAnsData(t *testing.T) {
	err := errors.GetOkAnsData("", UserModel.User{})
	if len(err) == 0 {
		log.Fatalf("error in GetOkAnsData")
	}
}

func TestGetOkAns(t *testing.T) {
	err := errors.GetOkAns("")
	if len(err) == 0 {
		log.Fatalf("error in GetOkAns")
	}
}

func TestGetErrorBadPasswordAns(t *testing.T) {
	err := errors.GetErrorBadPasswordAns()
	if len(err) == 0 {
		log.Fatalf("error in GetErrorBadPasswordAns")
	}
}

func TestGetErrorNoUserAns(t *testing.T) {
	err := errors.GetErrorNoUserAns()
	if len(err) == 0 {
		log.Fatalf("error in GetErrorNoUserAns")
	}
}

func TestGetErrorNotNumberAns(t *testing.T) {
	err := errors.GetErrorNotNumberAns()
	if len(err) == 0 {
		log.Fatalf("error in GetErrorNotNumberAns")
	}
}

func TestGetErrorNotPostAns(t *testing.T) {
	err := errors.GetErrorNotPostAns()
	if len(err) == 0 {
		log.Fatalf("error in GetErrorNotPostAns")
	}
}

func TestGetErrorWrongCookieAns(t *testing.T) {
	err := errors.GetErrorWrongCookieAns()
	if len(err) == 0 {
		log.Fatalf("error in GetErrorWrongCookieAns")
	}
}

func TestGetErrorNoCockyAns(t *testing.T) {
	err := errors.GetErrorNoCockyAns()
	if len(err) == 0 {
		log.Fatalf("error in GetErrorNoCockyAns")
	}
}

func TestGetErrorLoginExistAns(t *testing.T) {
	err := errors.GetErrorLoginExistAns()
	if len(err) == 0 {
		log.Fatalf("error in GetErrorLoginExistAns")
	}
}

func TestGetErrorBadJsonAns(t *testing.T) {
	err := errors.GetErrorBadJsonAns()
	if len(err) == 0 {
		log.Fatalf("error in GetErrorBadJsonAns")
	}
}

func TestGetErrorBadCsrfAns(t *testing.T) {
	err := errors.GetErrorBadCsrfAns(UserRepository.GetSessionError)
	if len(err) == 0 {
		log.Fatalf("error in GetErrorBadCsrfAns")
	}
}

func TestGetSendOkAns(t *testing.T) {
	err := errors.GetSendOkAns(LetterModel.Letter{})
	if len(err) == 0 {
		log.Fatalf("error in GetSendOkAns")
	}
}

func TestGetGetLettersOkAns(t *testing.T) {
	err := errors.GetGetLettersOkAns([]LetterModel.Letter{})
	if len(err) == 0 {
		log.Fatalf("error in GetGetLettersOkAns")
	}
}

func TestGetErrorReceivedLetterAns(t *testing.T) {
	err := errors.GetErrorReceivedLetterAns()
	if len(err) == 0 {
		log.Fatalf("error in GetErrorReceivedLetterAns")
	}
}

func TestGetErrorNoRecieverAns(t *testing.T) {
	err := errors.GetErrorNoRecieverAns()
	if len(err) == 0 {
		log.Fatalf("error in GetErrorNoRecieverAns")
	}
}

func TestGetErrorSaveErrorAns(t *testing.T) {
	err := errors.GetErrorSaveErrorAns()
	if len(err) == 0 {
		log.Fatalf("error in GetErrorSaveErrorAns")
	}
}

func TestProfileError(t *testing.T) {
	err := UserDelivery.ProfileError(nil, &http.Cookie{})
	if len(err) == 0 {
		log.Fatalf("error in ProfileError")
	}
	err = UserDelivery.ProfileError(UserRepository.CantUpdateUser, &http.Cookie{})
	if len(err) == 0 {
		log.Fatalf("error in ProfileError")
	}
	err = UserDelivery.ProfileError(UserRepository.CantGetUserOnUpdate, &http.Cookie{})
	if len(err) == 0 {
		log.Fatalf("error in ProfileError")
	}
}

func TestLogoutError(t *testing.T) {
	err := UserDelivery.LogoutError(nil)
	if len(err) == 0 {
		log.Fatalf("error in LogoutError")
	}
	err = UserDelivery.LogoutError(UserRepository.InvalidSession)
	if len(err) == 0 {
		log.Fatalf("error in LogoutError")
	}
	err = UserDelivery.LogoutError(UserRepository.RemoveSessionError)
	if len(err) == 0 {
		log.Fatalf("error in LogoutError")
	}
}

func TestCookieError(t *testing.T) {
	err := UserDelivery.CookieError(401)
	if len(err) == 0 {
		log.Fatalf("error in CookieError")
	}
	err = UserDelivery.CookieError(402)
	if len(err) == 0 {
		log.Fatalf("error in CookieError")
	}
}

func TestSignInError(t *testing.T) {
	err := UserDelivery.SignInError(nil, &http.Cookie{})
	if len(err) == 0 {
		log.Fatalf("error in SignInError")
	}
	err = UserDelivery.SignInError(UserRepository.CantAddSession, &http.Cookie{})
	if len(err) == 0 {
		log.Fatalf("error in SignInError")
	}
	err = UserDelivery.SignInError(UserRepository.CantGetUserByEmail, &http.Cookie{})
	if len(err) == 0 {
		log.Fatalf("error in SignInError")
	}
	err = UserDelivery.SignInError(UserRepository.GetSessionError, &http.Cookie{})
	if len(err) == 0 {
		log.Fatalf("error in SignInError")
	}
	err = UserDelivery.SignInError(UserUseCase.WrongPasswordError, &http.Cookie{})
	if len(err) == 0 {
		log.Fatalf("error in SignInError")
	}
	err = UserDelivery.SignInError(UserRepository.RemoveSessionError, &http.Cookie{})
	if len(err) == 0 {
		log.Fatalf("error in SignInError")
	}
}

func TestSignUpError(t *testing.T) {
	err := UserDelivery.SignUpError(nil, &http.Cookie{})
	if len(err) == 0 {
		log.Fatalf("error in SignUpError")
	}
	err = UserDelivery.SignUpError(UserRepository.EmailAlreadyExists, &http.Cookie{})
	if len(err) == 0 {
		log.Fatalf("error in SignUpError")
	}
	err = UserDelivery.SignUpError(UserRepository.CantAddSession, &http.Cookie{})
	if len(err) == 0 {
		log.Fatalf("error in SignUpError")
	}
	err = UserDelivery.SignUpError(UserRepository.CantAddUser, &http.Cookie{})
	if len(err) == 0 {
		log.Fatalf("error in SignUpError")
	}
}

func TestGetLettersError(t *testing.T) {
	err := LetterDelivery.GetLettersError(nil, []LetterModel.Letter{})
	if len(err) == 0 {
		log.Fatalf("error in GetLettersError")
	}
	err = LetterDelivery.GetLettersError(LetterRepository.ReceivedLetterError, []LetterModel.Letter{})
	if len(err) == 0 {
		log.Fatalf("error in GetLettersError")
	}
	err = LetterDelivery.GetLettersError(LetterRepository.DbError, []LetterModel.Letter{})
	if len(err) == 0 {
		log.Fatalf("error in GetLettersError")
	}
	err = LetterDelivery.GetLettersError(context.UserFromContextError, []LetterModel.Letter{})
	if len(err) == 0 {
		log.Fatalf("error in GetLettersError")
	}
}

func TestSendLetterError(t *testing.T) {
	err := LetterDelivery.SendLetterError(nil, LetterModel.Letter{})
	if len(err) == 0 {
		log.Fatalf("error in SendLetterError")
	}
	err = LetterDelivery.SendLetterError(LetterRepository.ReceiverNotFound, LetterModel.Letter{})
	if len(err) == 0 {
		log.Fatalf("error in SendLetterError")
	}
	err = LetterDelivery.SendLetterError(LetterRepository.DbError, LetterModel.Letter{})
	if len(err) == 0 {
		log.Fatalf("error in SendLetterError")
	}
	err = LetterDelivery.SendLetterError(LetterRepository.SaveLetterError, LetterModel.Letter{})
	if len(err) == 0 {
		log.Fatalf("error in SendLetterError")
	}
}
