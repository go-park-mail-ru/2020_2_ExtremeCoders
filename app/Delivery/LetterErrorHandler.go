package Delivery

import (
	"CleanArch/app/Models"
)

func SendLetterError(code uint16) []byte{
	switch code {
	case 200:
		return getOkAns("")
	case 400:
		return getErrorSaveErrorAns()
	}
	return nil
}

func GetLettersError(code uint16, letters []Models.Letter) []byte{
	switch code {
	case 200:
		return getSaveOkAns(letters)
	case 400:
		return getErrorSaveErrorAns()
	}
	return nil
}