package errors

import (
	"MainApplication/internal/Letter/LetterModel"
	"MainApplication/internal/User/UserModel"
)

type LetterAns struct {
	Code    int
	Lid     uint64
	Letters []LetterModel.Letter
}

type AnswerGet struct {
	Code        uint16
	Description string
	sid         string
	User        UserModel.User
}