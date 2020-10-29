package Delivery

import (
	"CleanArch/app/Models"
	"encoding/json"
)

type LetterAns struct {
	code int
	letters []Models.Letter
}

func getErrorSaveErrorAns() []byte {
	err := AnswerGet{
		Code:        400,
		Description: "Could not save letter",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getSaveOkAns(letters []Models.Letter) []byte{
	ok:=LetterAns{
		code: 200,
		letters: letters,
	}
	ans, _ := json.Marshal(ok)
	return ans
}