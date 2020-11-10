package LetterDelivery

import (
	"CleanArch/internal/Letter/LetterModel"
	"CleanArch/internal/Letter/LetterUseCase"
	"CleanArch/internal/User/UserDelivery"
	errors "CleanArch/internal/errors"
	"github.com/golang/glog"
	"net/http"
	"time"
	"CleanArch/internal/pkg/context"
)

type Delivery struct{
	Uc LetterUseCase.UseCase
}

func (de *Delivery)SendLetter(w http.ResponseWriter, r *http.Request){
	glog.Info("hui")
	if r.Method != http.MethodPost {
		//glog.Info("RESPONSE: ",getErrorNotPostAns())
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
	//glog.Info("RESPONSE: ",SendLetterError(uint16(err), letter))
}

func (de *Delivery) GetRecvLetters(w http.ResponseWriter, r *http.Request){
	//glog.Info("REQUEST: ", r.URL.Path, r.Method, r.Form)
	user:=context.GetUserFromCtx(r.Context())
	err, letters:=de.Uc.GetRecievedLetters(user.Email)
	w.Write(GetLettersError(uint16(err), letters))
	//glog.Info("RESPONSE: ",GetLettersError(uint16(err), letters))
}

func (de *Delivery) GetSendLetters(w http.ResponseWriter, r *http.Request){
	//glog.Info("REQUEST: ", r.URL.Path, r.Method, r.Form)
	user:=context.GetUserFromCtx(r.Context())
	err, letters:=de.Uc.GetSendedLetters(user.Email)
	w.Write(GetLettersError(uint16(err), letters))
	//glog.Info("RESPONSE: ",GetLettersError(uint16(err), letters))
}