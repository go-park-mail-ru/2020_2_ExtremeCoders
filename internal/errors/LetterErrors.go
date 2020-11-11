package errors

import (
	"CleanArch/internal/Letter/LetterModel"
	"encoding/json"
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

func GetErrorReceivedLetterAns() []byte {
	err := &AnswerGet{
		Code:        408,
		Description: "Could not get letters",
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
