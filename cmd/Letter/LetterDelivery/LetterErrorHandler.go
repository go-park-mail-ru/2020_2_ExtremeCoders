package LetterDelivery

import (
	"CleanArch/cmd/Letter/LetterModel"
	"CleanArch/cmd/errors"
)

func SendLetterError(code uint16, letter LetterModel.Letter) []byte{
	switch code {
	case 200:
		return errors.GetSendOkAns(letter)
	case 408:
		return errors.GetErrorNoRecieverAns()
	case 409:
		return errors.GetErrorSaveErrorAns()
	}
	return nil
}

func GetLettersError(code uint16, letters []LetterModel.Letter) []byte{
	switch code {
	case 200:
		return errors.GetGetLettersOkAns(letters)
	case 400:
		return errors.GetErrorSaveErrorAns()
	}
	return nil
}