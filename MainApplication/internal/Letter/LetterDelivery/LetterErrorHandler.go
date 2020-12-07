package LetterDelivery

import (
	"MainApplication/internal/Letter/LetterModel"
	"MainApplication/internal/Letter/LetterRepository"
	"MainApplication/internal/errors"
	"MainApplication/internal/pkg/context"
	log "github.com/sirupsen/logrus"
)

func SendLetterError(err error, letter LetterModel.Letter) []byte {
	var returned []byte
	defer func() {
		log.WithFields(log.Fields{
			"RESPONSE": string(returned),
		}).Info("sent")
	}()
	switch err {
	case nil:
		returned = errors.GetSendOkAns(letter)
	case LetterRepository.ReceiverNotFound:
		returned = errors.GetErrorNoRecieverAns()
	case LetterRepository.DbError:
		returned = errors.GetErrorUnexpectedAns()
	case LetterRepository.SaveLetterError:
		returned = errors.GetErrorSaveErrorAns()
	}
	return returned
}

func GetLettersError(err error, letters []LetterModel.Letter) []byte {
	var returned []byte
	defer func() {
		log.WithFields(log.Fields{
			"RESPONSE": string(returned),
		}).Info("sent")
	}()
	switch err {
	case nil:
		returned = errors.GetGetLettersOkAns(letters)
	case LetterRepository.ReceivedLetterError:
		returned = errors.GetErrorReceivedLetterAns()
	case LetterRepository.DbError:
		returned = errors.GetErrorUnexpectedAns()
	case context.UserFromContextError:
		returned = errors.GetErrorUnexpectedAns()
	}
	return returned
}

func GetDeleteLetterError(err error) []byte {
	var returned []byte
	defer func() {
		log.WithFields(log.Fields{
			"RESPONSE": string(returned),
		}).Info("sent")
	}()
	switch err {
	case nil:
		returned = errors.GetOk()
	case LetterRepository.DeleteLetterError:
		returned = errors.GetDeleteLetterError(err)
	}
	return returned
}