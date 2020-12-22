package UserDelivery

import (
	"Mailer/MainApplication/internal/User/UserModel"
	"Mailer/MainApplication/internal/User/UserRepository"
	"Mailer/MainApplication/internal/User/UserUseCase"
	"Mailer/MainApplication/internal/errors"
	"Mailer/MainApplication/internal/pkg/context"
	FileServise "Mailer/FileService/proto"

	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Interface interface {

	Session(w http.ResponseWriter, r *http.Request)
	Signup(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	GetUserByRequest(r *http.Request) (*UserModel.User, *http.Cookie, uint16)
	Profile(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	LoadFile(user *UserModel.User, r *http.Request)
	GetAvatar(w http.ResponseWriter, r *http.Request)
}

type delivery struct {
	Uc          UserUseCase.UserUseCase
	FileManager FileServise.FileServiceClient
}

func New(usecase UserUseCase.UserUseCase, fileManager FileServise.FileServiceClient) Interface {
	return delivery{Uc: usecase, FileManager: fileManager}
}

func (de delivery) Session(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		de.SignIn(w, r)
	}
	if r.Method == http.MethodDelete {
		de.Logout(w, r)
	}
}

func (de delivery) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	var user UserModel.User
	user.Name = context.GetStrFormValueSafety(r, "name")
	user.Surname = context.GetStrFormValueSafety(r, "surname")
	user.Email = context.GetStrFormValueSafety(r, "email")
	user.Password = context.GetStrFormValueSafety(r, "password1")
	de.LoadFile(&user, r)
	err, sid := de.Uc.Signup(user)
	var response []byte
	if err == nil {
		cookie := &http.Cookie{
			Name:    "session_id",
			Value:   sid,
			Expires: time.Now().Add(15 * 10000 * time.Hour),
		}
		cookie.Path = "/"
		http.SetCookie(w, cookie)
		response = SignUpError(err, cookie)
	} else {
		response = SignUpError(UserRepository.CantAddUser, nil)
	}

	w.Write(response)
}

func (de delivery) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write(errors.GetErrorNotPostAns())
		return
	}
	var user UserModel.User

	user.Email = context.GetStrFormValueSafety(r, "email")
	user.Password = context.GetStrFormValueSafety(r, "password")
	err, sid := de.Uc.SignIn(user)
	var response []byte
	if err == nil {
		cookie := &http.Cookie{
			Name:    "session_id",
			Value:   sid,
			Expires: time.Now().Add(15 * 10000 * time.Hour),
		}
		cookie.Path = "/"
		http.SetCookie(w, cookie)
		response = SignInError(err, cookie)
	} else {
		response = SignInError(err, nil)
	}
	w.Write(response)
}

func (de delivery) GetUserByRequest(r *http.Request) (*UserModel.User, *http.Cookie, uint16) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		return nil, nil, 401
	}
	uid, ok := de.Uc.GetDB().IsOkSession(session.Value)
	if ok != nil {
		return nil, nil, 402
	}
	user, err := de.Uc.GetDB().GetUserByUID(uid)
	if err != nil {
		return nil, nil, 402
	}
	return user, session, 200
}

func (de delivery) Profile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		de.Signup(w, r)
		return
	}
	user, session, err := de.GetUserByRequest(r)
	if err != 200 {
		w.Write(CookieError(err))
		return
	}
	if r.Method == http.MethodGet {
		w.Write(errors.GetOkAnsData(session.Value, *user))
		return
	} else if r.Method == http.MethodPut {
		var up UserModel.User
		up.Email = user.Email
		up.Name = context.GetStrFormValueSafety(r, "profile_firstName")
		up.Surname = context.GetStrFormValueSafety(r, "profile_lastName")
		de.LoadFile(&up, r)
		err := de.Uc.Profile(up)
		w.Write(ProfileError(err, session))
		return
	}
	w.Write(errors.GetErrorUnexpectedAns())
}

func (de delivery) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Write(errors.GetErrorNotPostAns())
		return
	} else {
		_, session, err := de.GetUserByRequest(r)
		if err != 200 {
			w.Write(CookieError(err))
			return
		}

		e := de.Uc.Logout(session.Value)
		if e == nil {
			session.Expires = time.Now().AddDate(0, 0, -1)
			http.SetCookie(w, session)
		}
		w.Write(LogoutError(e))
		return
	}
}

func (de delivery) LoadFile(user *UserModel.User, r *http.Request) {
	file, fileHeader, err := r.FormFile("avatar")
	if file == nil {
		return
	}
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Println("EEERR", err)
	}
	avatar := FileServise.Avatar{
		Email:    (*user).Email,
		FileName: fileHeader.Filename,
		Content:  buf.Bytes(),
	}
	_, _ = de.FileManager.SetAvatar(r.Context(), &avatar)
}

func (de delivery) GetAvatar(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Write([]byte(""))
		return
	}
	if r.Method == http.MethodGet {
		user, _, Err := de.GetUserByRequest(r)
		if Err != 200 {
			CookieError(Err)
			return
		}
		avatar, err := de.FileManager.GetAvatar(r.Context(), &FileServise.User{Email: user.Email})
		if err != nil {
			fmt.Println("GET AVATAR ERROR ", err)
		}
		w.Header().Set("Content-Type", "image")
		w.Header().Set("Content-Length", strconv.Itoa(len(avatar.Content)))
		if _, err := w.Write(avatar.Content); err != nil {
			w.Write(errors.GetErrorUnexpectedAns())
			return
		}
		return
	}
}
