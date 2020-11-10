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
		return errors.GetErrorUnexpectedAns()
	case LetterRepository.SaveLetterError:
		return errors.GetErrorSaveErrorAns()
	}
	return nil
}

func GetLettersError(err error, letters []LetterModel.Letter) []byte{
	switch err {
	case nil:
		return errors.GetGetLettersOkAns(letters)
	case LetterRepository.ReceivedLetterError:
		return errors.GetErrorReceivedLetterAns()
	case LetterRepository.DbError:
		return errors.GetErrorUnexpectedAns()
	}
	return nil
}