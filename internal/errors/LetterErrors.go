package errors

import (
	"CleanArch/internal/Letter/LetterModel"
	"encoding/json"
	"errors"
)

type LetterAns struct {
	Code    int
	Lid     uint64
	Letters []LetterModel.Letter
}

func GetErrorSaveErrorAns() []byte {
	err := &AnswerGet{
		Code:        400,
		Description: "Could not save letter",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func GetErrorNoRecieverAns() []byte {
	err := &AnswerGet{
		Code:        408,
		Description: "No such user in DB",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func GetGetLettersOkAns(letters []LetterModel.Letter) []byte {
	ok := &LetterAns{
		Code:    200,
		Letters: letters,
	}
	ans, _ := json.Marshal(ok)
	return ans
}

func GetSendOkAns(letters LetterModel.Letter) []byte {
	ok := &LetterAns{
		Code: 200,
		Lid:  letters.Id,
	}
	ans, _ := json.Marshal(ok)
	return ans
}

func SaveLetterError(err error)[]byte{
	erro := &AnswerGet{
		Code:        408,
		Description: errors.Unwrap(err).Error(),
	}
	ans, _ := json.Marshal(erro)
	return ans
}