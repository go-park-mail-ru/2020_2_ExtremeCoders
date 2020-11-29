package FolderDelivery

import (
	mailProto "MainApplication/proto/MailService"
	"encoding/json"
)

//var getFolderListError = errors.New("getErrorListError")
type LetterErr struct {
	Code        int
	Description string
}

type LetterList struct{
	Code        int
	Description string
	letter []*mailProto.Letter
}

func GetFoldersError(err error) []byte {
	ans := &LetterErr{
		Code:        400,
		Description: err.Error(),
	}
	res, _ := json.Marshal(ans)
	return res
}

func ProtoResponseAnswer(pbLetter *mailProto.Response)[]byte{
	code:=200
	if !pbLetter.Ok{
		code=409
	}
	ans:=LetterErr{
		Code: code,
		Description: pbLetter.Description,
	}
	res, _ := json.Marshal(ans)
	return res
}

func ProtoLetterListAnswer(pbLetter *mailProto.LetterListResponse)[]byte{
	code:=200
	if !pbLetter.Result.Ok{
		code=409
	}
	ans:=LetterList{
		Code: code,
		Description: pbLetter.Result.Description,
		letter: pbLetter.Letter,
	}
	res, _ := json.Marshal(ans)
	return res
}