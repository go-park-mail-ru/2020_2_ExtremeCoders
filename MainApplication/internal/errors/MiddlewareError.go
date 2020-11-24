package errors

import "encoding/json"

func GetErrorBadCsrfAns(e error) []byte {
	err := &AnswerGet{
		Code:        500,
		Description: e.Error(),
	}
	ans, _ := json.Marshal(err)
	return ans
}