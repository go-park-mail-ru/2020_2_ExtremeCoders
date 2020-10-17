package Delivery

import (
	"CleanArch/User/Models"
	"CleanArch/User/UseCase"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Delivery struct{
	Uc UseCase.UseCase
}

func (yaFood *Delivery)Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SIGNUP GOT: ", r.URL, r.Body, r.Method)
	if r.Method != http.MethodPost {
		return
	}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var user Models.User
	err := dec.Decode(&user)
	if err != nil {
		w.Write(getErrorBadJsonAns())
		return
	}
	code, cookie:=yaFood.Uc.Signup(user)
	if cookie!=nil{
		http.SetCookie(w, cookie)
	}
	response:=SignUpError(code, cookie)
	w.Write(response)
}

func (yaFood *Delivery)SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SIGNIN GOT: ", r.URL, r.Body)
	if r.Method != http.MethodPost {
		w.Write(getErrorNotPostAns())
		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var user Models.User
	err := dec.Decode(&user)
	if err != nil {
		w.Write(getErrorBadJsonAns())
		return
	}
	code, cookie:=yaFood.Uc.SignIn(user)
	if cookie!=nil{http.SetCookie(w, cookie)}
	response:=SignInError(code, cookie)
	w.Write(response)
}

func (yaFood *Delivery)getUserByRequest(r *http.Request) (*Models.User, *http.Cookie, uint16) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		return nil,nil, 401
	}
	uid, ok := yaFood.Uc.Db.IsOkSession(session.Value)
	if !ok{
		return nil, nil, 402
	}
	user:=yaFood.Uc.Db.GetUserByUID(uid)
	return user, session, 200
}

func (yaFood *Delivery)Profile(w http.ResponseWriter, r *http.Request) {
	user, session, err:=yaFood.getUserByRequest(r)
	if err!=200{
		CookieError(err)
		return
	}
	if r.Method == http.MethodGet {
		w.Write(getOkAnsData(session.Value, *user))
		return
	} else if r.Method != http.MethodPost {
		w.Write(getErrorNotPostAns())
		return
	} else {
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var up Models.User
		err := dec.Decode(&up)
		if err != nil {
			w.Write(getErrorBadJsonAns())
			return
		}
		yaFood.Uc.Db.UpdateProfile(up, user.Email)
		w.Write(getOkAns(session.Value))
		return
	}
	w.Write(getErrorUnexpectedAns())
}

func (yaFood *Delivery)Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LOGOUT GOT: ", r.URL, r.Form, r.Method)
	if r.Method != http.MethodPost {
		w.Write(getErrorNotPostAns())
		return
	} else {
		_, session, err:=yaFood.getUserByRequest(r)
		if err!=200{
			CookieError(err)
		}
		uid, ok :=yaFood.Uc.Db.IsOkSession(session.Value)
		if !ok {
			w.Write(getErrorWrongCookieAns())
			return
		}
		yaFood.Uc.Db.RemoveSession(uid, session.Value)
		w.Write(getOkAns(session.Value))
		session.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, session)
		return
	}
	w.Write(getErrorUnexpectedAns())
}