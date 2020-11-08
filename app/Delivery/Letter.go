package Delivery

import (
	"CleanArch/app/Models"
	"github.com/golang/glog"
	"net/http"
	"time"

)

func (yaFood *Delivery)SendLetter(w http.ResponseWriter, r *http.Request){
	glog.Info("REQUEST: ", r.URL.Path, r.Method, r.Form)
	if r.Method != http.MethodPost {
		glog.Info("RESPONSE: ",getErrorNotPostAns())
		w.Write(getErrorNotPostAns())
		return
	}
	var letter Models.Letter
	user, _, code:=yaFood.getUserByRequest(r)
	if code !=200{
		w.Write(getErrorUnexpectedAns())
		glog.Info("RESPONSE: ",getErrorUnexpectedAns())
		return
	}
	letter.Sender = user.Email
	letter.Receiver = getStrFormValueSafety(r,"to")
	letter.Theme = getStrFormValueSafety(r,"theme")
	letter.Text = getStrFormValueSafety(r,"text")
	letter.DateTime=time.Now().Unix()
	err:=yaFood.Uc.SaveLetter(&letter)
	w.Write(SendLetterError(uint16(err), letter))
	glog.Info("RESPONSE: ",SendLetterError(uint16(err), letter))
}

func (yaFood *Delivery) GetRecvLetters(w http.ResponseWriter, r *http.Request){
	glog.Info("REQUEST: ", r.URL.Path, r.Method, r.Form)
	user, _, code:=yaFood.getUserByRequest(r)
	if code != 200{
		w.Write(getErrorNoCockyAns())
		glog.Info("RESPONSE: ",getErrorNoCockyAns())
		return
	}
	err, letters:=yaFood.Uc.GetRecievedLetters(user.Email)
	w.Write(GetLettersError(uint16(err), letters))
	glog.Info("RESPONSE: ",GetLettersError(uint16(err), letters))
}

func (yaFood *Delivery) GetSendLetters(w http.ResponseWriter, r *http.Request){
	glog.Info("REQUEST: ", r.URL.Path, r.Method, r.Form)
	user, _, code:=yaFood.getUserByRequest(r)
	if code != 200{
		w.Write(getErrorNoCockyAns())
		glog.Info("RESPONSE: ",getErrorNoCockyAns())
		return
	}
	err, letters:=yaFood.Uc.GetSendedLetters(user.Email)
	w.Write(GetLettersError(uint16(err), letters))
	glog.Info("RESPONSE: ",GetLettersError(uint16(err), letters))
}