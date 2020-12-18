package errors

import (
	"Mailer/MainApplication/internal/Letter/LetterModel"
	"Mailer/MainApplication/internal/User/UserModel"
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
