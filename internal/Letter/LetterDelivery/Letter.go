package LetterDelivery

import (
	"CleanArch/internal/Letter/LetterModel"
	"CleanArch/internal/Letter/LetterUseCase"
	"CleanArch/internal/User/UserDelivery"
	"CleanArch/internal/errors"
	"CleanArch/internal/pkg/context"
	"github.com/golang/glog"
	"net/http"
	"time"
)

type Interface interface {
	SendLetter(w http.ResponseWriter, r *http.Request)
	GetRecvLetters(w http.ResponseWriter, r *http.Request)
	GetSendLetters(w http.ResponseWriter, r *http.Request)
}

type delivery struct {
	Uc LetterUseCase.LetterUseCase
}

func New(usecase LetterUseCase.LetterUseCase) Interface {
	return delivery{Uc: usecase}
}

func (de delivery) SendLetter(w http.ResponseWriter, r *http.Request) {
	glog.Info("hui")
	if r.Method != http.MethodPost {
		w.Write(errors.GetErrorNotPostAns())
		return
	}
	var letter LetterModel.Letter
	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetLettersError(er, nil))
		return
	}
	letter.Sender = user.Email
	letter.Receiver = UserDelivery.GetStrFormValueSafety(r, "to")
	letter.Theme = UserDelivery.GetStrFormValueSafety(r, "theme")
	letter.Text = UserDelivery.GetStrFormValueSafety(r, "text")
	letter.DateTime = time.Now().Unix()
	err := de.Uc.SaveLetter(&letter)
	w.Write(SendLetterError(err, letter))
}

func (de delivery) GetRecvLetters(w http.ResponseWriter, r *http.Request) {
	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetLettersError(er, nil))
		return
	}
	err, letters := de.Uc.GetReceivedLetters(user.Email)
	w.Write(GetLettersError(err, letters))
}

func (de delivery) GetSendLetters(w http.ResponseWriter, r *http.Request) {
	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetLettersError(er, nil))
		return
	}
	err, letters := de.Uc.GetSendedLetters(user.Email)
	w.Write(GetLettersError(err, letters))
}
