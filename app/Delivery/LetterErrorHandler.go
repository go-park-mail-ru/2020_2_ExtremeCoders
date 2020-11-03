package Delivery

import (
	"CleanArch/app/Models"
)

func SendLetterError(code uint16, letter Models.Letter) []byte{
	switch code {
	case 200:
		return getSendOkAns(letter)
	case 400:
		return getErrorSaveErrorAns()
	}
	return nil
}

func GetLettersError(code uint16, letters []Models.Letter) []byte{
	switch code {
	case 200:
		return getGetLettersOkAns(letters)
	case 400:
		return getErrorSaveErrorAns()
	}
	return nil
}