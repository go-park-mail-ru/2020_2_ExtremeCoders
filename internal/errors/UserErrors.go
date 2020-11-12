package errors

import (
	"CleanArch/internal/User/UserModel"
	"encoding/json"
	"fmt"
)

type AnswerGet struct {
	Code        uint16
	Description string
	sid         string
	User        UserModel.User
}

func GetErrorBadJsonAns() []byte {
	err := &AnswerGet{
		Code:        400,
		Description: "Bad Json",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func GetErrorLoginExistAns() []byte {
	err := &AnswerGet{
		Code:        401,
		Description: "This Email has already exists",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func GetErrorUnexpectedAns() []byte {
	err := &AnswerGet{
		Code:        500,
		Description: "Unexpected error",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func GetErrorNoCockyAns() []byte {
	err := &AnswerGet{
		Code:        401,
		Description: "not authorized user",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func GetErrorWrongCookieAns() []byte {
	err := &AnswerGet{
		Code:        403,
		Description: "wrong session id",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func GetErrorNotPostAns() []byte {
	err := &AnswerGet{
		Code:        400,
		Description: "Do not require request's method, expected POST",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func GetErrorNotNumberAns() []byte {
	err := &AnswerGet{
		Code:        400,
		Description: "Not number",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func GetErrorNoUserAns() []byte {
	err := &AnswerGet{
		Code:        404,
		Description: "Do not find this user in db",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func GetErrorBadPasswordAns() []byte {
	err := &AnswerGet{
		Code:        401,
		Description: "Wrong Password",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func GetOkAns(cocky string) []byte {
	ok := &AnswerGet{
		Code:        200,
		Description: "ok",
		sid:         cocky,
	}
	ans, _ := json.Marshal(ok)
	return ans
}

func GetOkAnsData(cocky string, data UserModel.User) []byte {
	fmt.Println("DATA::::::::::", data.Email, data.Name, data.Password)
	ok := &AnswerGet{
		Code:        200,
		Description: "ok",
		sid:         cocky,
		User:        data,
	}
	ans, _ := json.Marshal(ok)
	return ans
}

func GetAddSessionError() []byte {
	ok := &AnswerGet{
		Code:        401,
		Description: "Could not add session",
	}
	ans, _ := json.Marshal(ok)
	return ans
}

func RemoveSessionError() []byte {
	ok := &AnswerGet{
		Code:        408,
		Description: "Could not remove session",
	}
	ans, _ := json.Marshal(ok)
	return ans
}

func AddUserError() []byte {
	ok := &AnswerGet{
		Code:        407,
		Description: "Could not add user",
	}
	ans, _ := json.Marshal(ok)
	return ans
}

func UpdateProfileError() []byte {
	ok := &AnswerGet{
		Code:        407,
		Description: "Could not update profile",
	}
	ans, _ := json.Marshal(ok)
	return ans
}

func GetUserOnUpdateError() []byte {
	ok := &AnswerGet{
		Code:        407,
		Description: "Could not get user on update",
	}
	ans, _ := json.Marshal(ok)
	return ans
}
