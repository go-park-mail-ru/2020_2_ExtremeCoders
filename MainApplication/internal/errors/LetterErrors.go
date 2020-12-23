package errors

import (
	"MainApplication/internal/Letter/LetterModel"
)



func GetErrorSaveErrorAns() []byte {
	err := &AnswerGet{
		Code:        400,
		Description: "Could not save letter",
	}
	ans, _ := err.MarshalJSON()
	return ans
}

func GetErrorNoRecieverAns() []byte {
	err := &AnswerGet{
		Code:        408,
		Description: "No such user in DB",
	}
	ans, _ := err.MarshalJSON()
	return ans
}

func GetErrorReceivedLetterAns() []byte {
	err := &AnswerGet{
		Code:        408,
		Description: "Could not get letters",
	}
	ans, _ := err.MarshalJSON()
	return ans
}

func GetGetLettersOkAns(letters []LetterModel.Letter) []byte {
	ok := &LetterAns{
		Code:    200,
		Letters: letters,
	}
	ans, _ := ok.MarshalJSON()
	return ans
}

func GetSendOkAns(letters LetterModel.Letter) []byte {
	ok := &LetterAns{
		Code: 200,
		Lid:  letters.Id,
	}
	ans, _ := ok.MarshalJSON()
	return ans
}

func GetDeleteLetterError(err error)[]byte{
	ans:= &AnswerGet{
		Code:        500,
		Description: err.Error(),
	}
	jsAns, _:=ans.MarshalJSON()
	return jsAns
}