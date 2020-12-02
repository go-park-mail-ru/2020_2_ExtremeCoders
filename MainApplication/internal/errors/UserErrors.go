package errors

import (
	"MainApplication/internal/User/UserModel"
	"fmt"
)



func GetErrorBadJsonAns() []byte {
	err := &AnswerGet{
		Code:        400,
		Description: "Bad Json",
	}
	ans, _ := err.MarshalJSON()
	return ans
}

func GetErrorLoginExistAns() []byte {
	err := &AnswerGet{
		Code:        401,
		Description: "This Email has already exists",
	}
	ans, _ := err.MarshalJSON()
	return ans
}

func GetErrorUnexpectedAns() []byte {
	err := &AnswerGet{
		Code:        500,
		Description: "Unexpected error",
	}
	ans, _ := err.MarshalJSON()
	return ans
}

func GetErrorNoCockyAns() []byte {
	err := &AnswerGet{
		Code:        401,
		Description: "not authorized user",
	}
	ans, _ := err.MarshalJSON()
	return ans
}

func GetErrorWrongCookieAns() []byte {
	err := &AnswerGet{
		Code:        403,
		Description: "wrong session id",
	}
	ans, _ := err.MarshalJSON()
	return ans
}

func GetErrorNotPostAns() []byte {
	err := &AnswerGet{
		Code:        400,
		Description: "Do not require request's method, expected POST",
	}
	ans, _ := err.MarshalJSON()
	return ans
}

func GetErrorNotNumberAns() []byte {
	err := &AnswerGet{
		Code:        400,
		Description: "Not number",
	}
	ans, _ := err.MarshalJSON()
	return ans
}

func GetErrorNoUserAns() []byte {
	err := &AnswerGet{
		Code:        404,
		Description: "Do not find this user in db",
	}
	ans, _ := err.MarshalJSON()
	return ans
}

func GetErrorBadPasswordAns() []byte {
	err := &AnswerGet{
		Code:        401,
		Description: "Wrong Password",
	}
	ans, _ := err.MarshalJSON()
	return ans
}

func GetOkAns(cocky string) []byte {
	ok := &AnswerGet{
		Code:        200,
		Description: "ok",
		sid:         cocky,
	}
	ans, _ := ok.MarshalJSON()
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
	ans, _ := ok.MarshalJSON()
	return ans
}

func GetAddSessionError() []byte {
	ok := &AnswerGet{
		Code:        401,
		Description: "Could not add session",
	}
	ans, _ := ok.MarshalJSON()
	return ans
}

func RemoveSessionError() []byte {
	ok := &AnswerGet{
		Code:        408,
		Description: "Could not remove session",
	}
	ans, _ := ok.MarshalJSON()
	return ans
}

func AddUserError() []byte {
	ok := &AnswerGet{
		Code:        407,
		Description: "Could not add user",
	}
	ans, _ := ok.MarshalJSON()
	return ans
}

func UpdateProfileError() []byte {
	ok := &AnswerGet{
		Code:        407,
		Description: "Could not update profile",
	}
	ans, _ := ok.MarshalJSON()
	return ans
}

func GetUserOnUpdateError() []byte {
	ok := &AnswerGet{
		Code:        407,
		Description: "Could not get user on update",
	}
	ans, _ := ok.MarshalJSON()
	return ans
}
