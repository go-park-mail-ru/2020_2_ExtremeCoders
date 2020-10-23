package Delivery

import (
	"CleanArch/User/Models"
	"CleanArch/User/UseCase"
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Delivery struct{
	Uc UseCase.UseCase
}

func (yaFood *Delivery)Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Print("SIGNUP: ")
	fmt.Print("\n\n")
	if r.Method != http.MethodPost {
		return
	}
	var user Models.User
	user.Name = r.FormValue("name")
	user.Surname = r.FormValue("surname")
	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password1")

	code, cookie:=yaFood.Uc.Signup(user)
	yaFood.LoadFile(&user,r)
	fmt.Print("\n\n")
	if cookie!=nil{
		http.SetCookie(w, cookie)
	}
	response:=SignUpError(code, cookie)
	w.Write(response)
}

func (yaFood *Delivery)SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Print("SIGNIN: ")
	yaFood.Uc.Db.ShowAll()
	fmt.Print("\n\n")
	fmt.Println("SIGNIN GOT: ", r.URL, r.Body)
	if r.Method != http.MethodPost {
		w.Write(getErrorNotPostAns())
		return
	}
	var user Models.User
	user.Email=r.FormValue("email")
	user.Password=r.FormValue("password")
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
	fmt.Print("PROFILE: ")

	fmt.Print("\n\n")
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
		var up Models.User
		up.Name=r.FormValue("profile_firstName")
		up.Surname=r.FormValue("profile_lastName")
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

func (yaFood *Delivery)LoadFile(user *Models.User, r *http.Request){
	file, fileHeader, err := r.FormFile("avatar")
	if file == nil {
		fmt.Println("FILE IS EMPTY")
		return
	}
	(*user).Img = fileHeader.Filename
	fmt.Println("FILLLLLLLLLLLLLLLLLLLLLLLE", fileHeader.Filename, err, r.FormValue("Name"))
	f, err := os.Create(fileHeader.Filename)
	if err != nil {
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

func (yaFood *Delivery)GetAvatar(w http.ResponseWriter, r *http.Request){
	fmt.Println("getAvatar GOT: ", r.URL, r.Form, r.Method)
	if r.Method == http.MethodOptions {
		w.Write([]byte(""))
		return
	}
	if r.Method == http.MethodGet {
		user, _, Err:=yaFood.getUserByRequest(r)
		if Err!=200{
			CookieError(Err)
			return
		}
		if (*user).Img == "" {
			fmt.Println("USER HAVE NOT AVATAR")
			w.Write([]byte("USER HAVE NOT AVATAR"))
			return
		}

		file, err := os.Open((*user).Img) // path to image file
		if err != nil {
			fmt.Println("ERROR", err)
			return
		}

		img, fmtName, err := image.Decode(file)
		fmt.Println("FMT NAME", fmtName)
		if err != nil {
			fmt.Println(err)
			w.Write(getErrorUnexpectedAns())
		}

		buffer := new(bytes.Buffer)
		if err := jpeg.Encode(buffer, img, nil); err != nil {
			fmt.Println("unable to encode image.")
		}

		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
		if _, err := w.Write(buffer.Bytes()); err != nil {
			w.Write(getErrorUnexpectedAns())
			fmt.Println("unable to write image.")
			return
		}
		return
	}
}