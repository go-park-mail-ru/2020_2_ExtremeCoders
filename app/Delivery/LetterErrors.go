package Delivery

import (
	"CleanArch/app/Models"
	"encoding/json"
)

type LetterAns struct {
	Code    int
	Lid     uint64
	Letters []Models.Letter
}

func getErrorSaveErrorAns() []byte {
	err := &AnswerGet{
		Code:        400,
		Description: "Could not save letter",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getGetLettersOkAns(letters []Models.Letter) []byte {
	ok := &LetterAns{
		Code:    200,
		Letters: letters,
	}
	ans, _ := json.Marshal(ok)
	return ans
}

func getSendOkAns(letters Models.Letter) []byte {
	ok := &LetterAns{
		Code: 200,
		Lid:  letters.Id,
	}
	ans, _ := json.Marshal(ok)
	return ans
}
