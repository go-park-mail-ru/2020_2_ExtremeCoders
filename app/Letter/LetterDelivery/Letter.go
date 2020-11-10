package LetterDelivery

import (
	"CleanArch/app/Letter/LetterUseCase"
	"CleanArch/app/User/UserModel"
	"github.com/golang/glog"
	"net/http"
	"time"
	errors "CleanArch/app/errors"
)

type Delivery struct{
	Uc LetterUseCase.UseCase
}

func (yaFood *Delivery)SendLetter(w http.ResponseWriter, r *http.Request){
	glog.Info("hui")
	if r.Method != http.MethodPost {
		//glog.Info("RESPONSE: ",getErrorNotPostAns())
		w.Write(errors.GetErrorNotPostAns())
		return
	}
	var letter UserModel.Letter
	user, _, code:=yaFood.getUserByRequest(r)
	if code !=200{
		w.Write(errors.GetErrorUnexpectedAns())
		//glog.Info("RESPONSE: ",getErrorUnexpectedAns())
		return
	}
	letter.Sender = user.Email
	letter.Receiver = getStrFormValueSafety(r,"to")
	letter.Theme = getStrFormValueSafety(r,"theme")
	letter.Text = getStrFormValueSafety(r,"text")
	letter.DateTime=time.Now().Unix()
	err:=yaFood.Uc.SaveLetter(&letter)
	w.Write(SendLetterError(uint16(err), letter))
	//glog.Info("RESPONSE: ",SendLetterError(uint16(err), letter))
}

func (yaFood *Delivery) GetRecvLetters(w http.ResponseWriter, r *http.Request){
	//glog.Info("REQUEST: ", r.URL.Path, r.Method, r.Form)
	user, _, code:=yaFood.getUserByRequest(r)
	if code != 200{
		w.Write(errors.GetErrorNoCockyAns())
		//glog.Info("RESPONSE: ",getErrorNoCockyAns())
		return
	}
	err, letters:=yaFood.Uc.GetRecievedLetters(user.Email)
	w.Write(GetLettersError(uint16(err), letters))
	//glog.Info("RESPONSE: ",GetLettersError(uint16(err), letters))
}

func (yaFood *Delivery) GetSendLetters(w http.ResponseWriter, r *http.Request){
	//glog.Info("REQUEST: ", r.URL.Path, r.Method, r.Form)
	user, _, code:=yaFood.getUserByRequest(r)
	if code != 200{
		w.Write(errors.GetErrorNoCockyAns())
		//glog.Info("RESPONSE: ",getErrorNoCockyAns())
		return
	}
	err, letters:=yaFood.Uc.GetSendedLetters(user.Email)
	w.Write(GetLettersError(uint16(err), letters))
	//glog.Info("RESPONSE: ",GetLettersError(uint16(err), letters))
}