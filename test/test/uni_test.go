package test

import (
	"CleanArch/internal/User/UserRepository/UserPostgres"
	"CleanArch/internal/errors"
	"CleanArch/internal/pkg/context"
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
)


func TestPasswordBcrypt(t *testing.T) {
	raw:="1538"
	enc:=UserPostgres.PasswordBcrypt([]byte(raw))
	if bcrypt.CompareHashAndPassword(enc,[]byte(raw))!=nil{
		log.Fatalf("error in PasswordBcrypt")
	}
}

func TestGenerateCSRF(t *testing.T) {
	csrf:=context.GenerateCSRF()
	if csrf==""{
		log.Fatalf("error in GenerateCSRF")
	}
}

func TestGetUserOnUpdateError(t *testing.T) {
	err:=errors.GetUserOnUpdateError()
	if len(err)==0{
		log.Fatalf("error in GetUserOnUpdateError")
	}
}

func TestUpdateProfileError(t *testing.T) {
	err:=errors.UpdateProfileError()
	if len(err)==0{
		log.Fatalf("error in UpdateProfileError")
	}
}

func TestAddUserError(t *testing.T) {
	err:=errors.AddUserError()
	if len(err)==0{
		log.Fatalf("error in AddUserError")
	}
}

func TestRemoveSessionError(t *testing.T) {
	err:=errors.RemoveSessionError()
	if len(err)==0{
		log.Fatalf("error in RemoveSessionError")
	}
}

func TestGetAddSessionError(t *testing.T) {
	err:=errors.GetAddSessionError()
	if len(err)==0{
		log.Fatalf("error in GetAddSessionError")
	}
}

func TestGetOkAnsData(t *testing.T) {
	err:=errors.GetAddSessionError()
	if len(err)==0{
		log.Fatalf("error in GetOkAnsData")
	}
}

func TestGetOkAns(t *testing.T) {
	err:=errors.GetOkAns("")
	if len(err)==0{
		log.Fatalf("error in GetOkAns")
	}
}

func TestGetErrorBadPasswordAns(t *testing.T) {
	err:=errors.GetErrorBadPasswordAns()
	if len(err)==0{
		log.Fatalf("error in GetErrorBadPasswordAns")
	}
}