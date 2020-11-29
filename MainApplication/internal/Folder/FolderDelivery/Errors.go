package FolderDelivery

import (
	"encoding/json"
)

getFolderListError = errors.New("getErrorListError")

type LetterErr struct {
	Code    int
	Description string
}


func GetFoldersError(err error) []byte {
	ans := &LetterErr{
		Code:        400,
		Description: err.Error(),
	}
	res, _ := json.Marshal(ans)
	return res
}

