package Delivery

import (
	"CleanArch/app/Models"
	"fmt"
	"net/http"
	"time"
)

func (yaFood *Delivery)SendLetter(w http.ResponseWriter, r *http.Request){
	fmt.Print("Send Letter: ")
	fmt.Print("\n\n")
	if r.Method != http.MethodPost {
		return
	}
	var letter Models.Letter
	letter.Sender = r.FormValue("from")
	letter.Receiver = r.FormValue("to")
	letter.Theme = r.FormValue("title")
	letter.Text = r.FormValue("text")
	letter.DateTime=time.Now().Unix()
	err:=yaFood.Uc.SaveLetter(letter)
	w.Write(SendLetterError(uint16(err)))
}

func (yaFood *Delivery)GetLetters(w http.ResponseWriter, r *http.Request){
	user, _, code:=yaFood.getUserByRequest(r)
	if code != 200{
		w.Write(getErrorNoCockyAns())
		return
	}
	err, letters:=yaFood.Uc.GetLetters(user.Email)
	w.Write(GetLettersError(uint16(err), letters))
}