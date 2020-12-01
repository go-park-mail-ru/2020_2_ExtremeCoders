package errors

func GetErrorBadCsrfAns(e error) []byte {
	err := &AnswerGet{
		Code:        500,
		Description: e.Error(),
	}
	ans, _ := err.MarshalJSON()
	return ans
}
