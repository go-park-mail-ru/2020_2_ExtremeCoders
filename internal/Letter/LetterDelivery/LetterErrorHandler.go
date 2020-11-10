package LetterDelivery

import (
	"CleanArch/internal/Letter/LetterModel"
	"CleanArch/internal/Letter/LetterRepository"
	"CleanArch/internal/errors"
)

func SendLetterError(err error, letter LetterModel.Letter) []byte{
	switch err {
	case nil:
		return errors.GetSendOkAns(letter)
	case LetterRepository.ReceiverNotFound:
		return errors.GetErrorNoRecieverAns()
	case LetterRepository.DbError:
		errors.GetErrorUnexpectedAns()
	case LetterRepository.SaveLetterError:
		errors.GetErrorSaveErrorAns()
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