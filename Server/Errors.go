package Server

import (
	"encoding/json"
	"fmt"
)

func getErrorBadJsonAns() []byte {
	err := &Answer{
		Code:        400,
		Description: "Bad Json",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getErrorUnexpectedAns() []byte {
	err := &Answer{
		Code:        500,
		Description: "Unexpected error",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getErrorNoCockyAns() []byte {
	err := &Answer{
		Code:        401,
		Description: "not authorized user",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getErrorWrongCookieAns() []byte {
	err := &Answer{
		Code:        401,
		Description: "wrong session id",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getErrorNotPostAns() []byte {
	err := &Answer{
		Code:        400,
		Description: "Do not require request's method, expected POST",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getErrorNotNumberAns() []byte {
	err := &Answer{
		Code:        400,
		Description: "Not number",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getErrorNoUserAns() []byte {
	err := &Answer{
		Code:        404,
		Description: "Do not find this user in db",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getErrorLoginExistAns() []byte {
	err := &Answer{
		Code:        401,
		Description: "This Email has already exists",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getErrorBadPasswordAns() []byte {
	err := &Answer{
		Code:        401,
		Description: "Wrong Password",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getOkAns(cocky string) []byte {
	ok := &Answer{
		Code:        200,
		Description: "ok",
		sid:         cocky,
	}
	ans, _ := json.Marshal(ok)
	return ans
}

func getOkAnsData(cocky string, data User) []byte {
	fmt.Println("DATA::::::::::", data.Email, data.Name, data.Password)
	ok := &Answer{
		Code:        200,
		Description: "ok",
		sid:         cocky,
		User:        data,
	}
	ans, _ := json.Marshal(ok)
	return ans
}
