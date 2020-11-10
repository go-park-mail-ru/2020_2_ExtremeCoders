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

type Delivery struct{
	Uc LetterUseCase.UseCase
}

func (de *Delivery)SendLetter(w http.ResponseWriter, r *http.Request){
	glog.Info("hui")
	if r.Method != http.MethodPost {
		w.Write(errors.GetErrorNotPostAns())
		return
	}
	var letter LetterModel.Letter
	user:=context.GetUserFromCtx(r.Context())
	letter.Sender = user.Email
	letter.Receiver = UserDelivery.GetStrFormValueSafety(r,"to")
	letter.Theme = UserDelivery.GetStrFormValueSafety(r,"theme")
	letter.Text = UserDelivery.GetStrFormValueSafety(r,"text")
	letter.DateTime=time.Now().Unix()
	err:=de.Uc.SaveLetter(&letter)
	w.Write(SendLetterError(err, letter))
}

func (de *Delivery) GetRecvLetters(w http.ResponseWriter, r *http.Request){
	user:=context.GetUserFromCtx(r.Context())
	err, letters:=de.Uc.GetReceivedLetters(user.Email)
	w.Write(GetLettersError(err, letters))
}

func (de *Delivery) GetSendLetters(w http.ResponseWriter, r *http.Request){
	user:=context.GetUserFromCtx(r.Context())
	err, letters:=de.Uc.GetSendedLetters(user.Email)
	w.Write(GetLettersError(err, letters))
}