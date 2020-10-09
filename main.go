package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	sizeSID = 32
)

var sidRunes = []rune("1234567890_qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

type Profile struct {
	method string
	value  string
}

type User struct {
	id       uint64
	Name     string
	Surname  string
	Email    string
	Password string
	//Date     uint64
	Img string
}

type Answer struct {
	Code        uint16
	Description string
	sid         string
	User        User
}

type loggedIn struct {
	sessions map[string]uint64
	users    map[string]*User
}

func getErrorBadJsonAns() []byte {
	err := &Answer{
		Code:        400,
		Description: "Bad Json",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getErrorUnexpectedAns() []byte {
	err := &Answer{
		Code:        500,
		Description: "Unexpected error",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getErrorNoCockyAns() []byte {
	err := &Answer{
		Code:        401,
		Description: "not authorized user",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getErrorWrongCookieAns() []byte {
	err := &Answer{
		Code:        401,
		Description: "wrong session id",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getErrorNotPostAns() []byte {
	err := &Answer{
		Code:        400,
		Description: "Do not require request's method, expected POST",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getErrorNotNumberAns() []byte {
	err := &Answer{
		Code:        400,
		Description: "Not number",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getErrorNoUserAns() []byte {
	err := &Answer{
		Code:        404,
		Description: "Do not find this user in db",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getErrorLoginExistAns() []byte {
	err := &Answer{
		Code:        401,
		Description: "This Email has already exists",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getErrorBadPasswordAns() []byte {
	err := &Answer{
		Code:        401,
		Description: "Wrong Password",
	}
	ans, _ := json.Marshal(err)
	return ans
}

func getOkAns(cocky string) []byte {
	ok := &Answer{
		Code:        200,
		Description: "ok",
		sid:         cocky,
	}
	ans, _ := json.Marshal(ok)
	return ans
}

func getOkAnsData(cocky string, data User) []byte {
	fmt.Println("DATA::::::::::", data.Email, data.Name, data.Password)
	ok := &Answer{
		Code:        200,
		Description: "ok",
		sid:         cocky,
		User:        data,
	}
	ans, _ := json.Marshal(ok)
	return ans
}

func generateSID(db *loggedIn) []rune {
	var sid = make([]rune, sizeSID)
	for {
		for i := 0; i < sizeSID; i++ {
			sid[i] = sidRunes[rand.Intn(len(sidRunes))]
		}
		_, exist := db.sessions[string(sid)]
		if !exist {
			break
		}
	}
	return sid
}

func generateUID(db *loggedIn) uint64 {
	var uid uint64
	for {
		for i := 0; i < sizeSID; i++ {
			uid = rand.Uint64()
		}
		var _, exist = db.users[strconv.FormatUint(uid, 10)]
		if !exist {
			break
		}
	}
	return uid
}

func (db *loggedIn) signin(w http.ResponseWriter, r *http.Request) {
	setHeader(w, r)
	fmt.Println("SIGNIN GOT: ", r.URL, r.Body)
	fmt.Println("USER", r.FormValue("Email"), r.FormValue("Password"))
	if r.Method == http.MethodOptions {
		w.Write([]byte(""))
		return
	}
	if r.Method != http.MethodPost {
		w.Write(getErrorNotPostAns())
		return
	}

	type Uinfo struct {
		Email    string
		Password string
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var user Uinfo
	err := dec.Decode(&user)
	if err != nil {
		w.Write(getErrorBadJsonAns())
		return
	}
	userEx, erro := db.users[user.Email]
	if !erro {
		w.Write(getErrorNoUserAns())
		return
	}
	if userEx.Password != user.Password {
		w.Write(getErrorBadPasswordAns())
		return
	}
	sid := string(generateSID(db))
	db.sessions[sid] = userEx.id
	cocky := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(24 * 7 * 4 * time.Hour),
	}
	cocky.Path = "/"
	http.SetCookie(w, cocky)
	w.Write(getOkAns(sid))
}

func (db *loggedIn) signup(w http.ResponseWriter, r *http.Request) {

	setHeader(w, r)
	fmt.Println("SIGNUP GOT: ", r.URL, r.Body, r.Method)
	if r.Method == http.MethodOptions {
		w.Write([]byte(""))
		return
	}
	if r.Method != http.MethodPost {
		w.Write(getErrorNotPostAns())
		return
	}


	var user User
	user.Name = r.FormValue("name")
	user.Surname = r.FormValue("surname")
	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password1")

	_, erro := db.users[user.Email]
	if erro {
		w.Write(getErrorLoginExistAns())
		return
	}

	user.id = generateUID(db)
	sid := string(generateSID(db))
	db.sessions[sid] = user.id
	db.users[user.Email] = &user

	fmt.Println("USER", user.Name, user.Password)

	file, fileHeader, err := r.FormFile("avatar")
	if file==nil{
		fmt.Println("FILE IS EMPTY")
		cocky := &http.Cookie{
			Name:    "session_id",
			Value:   sid,
			Expires: time.Now().Add(24 * 7 * 4 * time.Hour),
		}
		cocky.Path = "/"
		http.SetCookie(w, cocky)
		w.Write(getOkAns(sid))
		return
	}
	user.Img = fileHeader.Filename
	fmt.Println("FILLLLLLLLLLLLLLLLLLLLLLLE", fileHeader.Filename, err, r.FormValue("Name"))
	f, err := os.Create(fileHeader.Filename)
	if err != nil {
		fmt.Println("sendImg GOT ERROR1: ", err)
		cocky := &http.Cookie{
			Name:    "session_id",
			Value:   sid,
			Expires: time.Now().Add(24 * 7 * 4 * time.Hour),
		}
		cocky.Path = "/"
		http.SetCookie(w, cocky)
		w.Write(getOkAns(sid))
		http.Error(w, err.Error(), 500)
		return
	}
	defer f.Close()
	io.Copy(f, file)


	cocky := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(24 * 7 * 4 * time.Hour),
	}
	cocky.Path = "/"
	http.SetCookie(w, cocky)
	w.Write(getOkAns(sid))
}

func (db *loggedIn) updateProfile(changes *Profile, uid string) uint16 {
	if changes.method == "change Password" {
		db.users[uid].Password = changes.value
	} else if changes.method == "change Email" {
		db.users[uid].Email = changes.value
	} else if changes.method == "change Name" {
		db.users[uid].Name = changes.value
	} else if changes.method == "change Surname" {
		db.users[uid].Surname = changes.value
	} else if changes.method == "change Date" {
		//date, err := strconv.ParseUint(changes.value, 10, 64)
		//if err != nil{
		//	return 400
		//}
		//db.users[uid].Date =date
	} else if changes.method == "change Img" {
		db.users[uid].Img = changes.value
	}
	return 200
}

func (db *loggedIn) profile(w http.ResponseWriter, r *http.Request) {
	setHeader(w, r)
	fmt.Println("PROFILE GOT: ", r.URL, r.Form, r.Method)
	if r.Method == http.MethodOptions {
		w.Write([]byte(""))
		return
	}
	if r.Method == http.MethodGet {
		session, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			fmt.Println("NO COOKIE")
			w.Write(getErrorNoCockyAns())
			return
		}
		fmt.Println("COOKIE!!!!!!!!!!!!!!!!!!!")

		uid, ok := db.sessions[session.Value]
		if !ok {
			w.Write(getErrorWrongCookieAns())
			return
		}
		for _, val := range db.users {
			if (*val).id == uid {
				fmt.Println("If_DATA::::::", (*val).Password, (*val).Name)
				w.Write(getOkAnsData(session.Value, *val))
				return
			}
		}

	} else if r.Method != http.MethodPost {
		w.Write(getErrorNotPostAns())
		return
	} else {
		session, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			fmt.Println("NO COOKIE")
			w.Write(getErrorNoCockyAns())
			return
		}
		fmt.Println("COOKIE!!!!!!!!!!!!!!!!!!!")
		fmt.Println("NAME", r.FormValue("profile_firstName"), "SURNAME", r.FormValue("profile_lastName"))
		newName := r.FormValue("profile_firstName")
		newSurname := r.FormValue("profile_lastName")
		uid, ok := db.sessions[session.Value]
		if !ok {
			w.Write(getErrorWrongCookieAns())
			return
		}
		var currentUser *User
		for _, val := range db.users {
			if (*val).id == uid {
				currentUser = val
				break
			}
		}
		if currentUser==nil{
			w.Write(getErrorUnexpectedAns())
			return
		}
		fmt.Println("If_DATA::::::", (*currentUser).Password, (*currentUser).Name)
		(*currentUser).Name = newName
		(*currentUser).Surname = newSurname

		file, fileHeader, err := r.FormFile("avatar")
		if file==nil{
			fmt.Println("FILE IS EMPTY")
			w.Write(getOkAns(session.Value))
			return
		}
		(*currentUser).Img = fileHeader.Filename
		fmt.Println("FILLLLLLLLLLLLLLLLLLLLLLLE", fileHeader.Filename, err, r.FormValue("Name"))
		f, err := os.Create(fileHeader.Filename)
		if err != nil {
			fmt.Println("sendImg GOT ERROR1: ", err)
			http.Error(w, err.Error(), 500)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		w.Write(getOkAns(session.Value))
		return
	}
	w.Write(getErrorUnexpectedAns())
}

func (db *loggedIn) getAvatar(w http.ResponseWriter, r *http.Request){
	setHeader(w, r)
	fmt.Println("getAvatar GOT: ", r.URL, r.Form, r.Method)
	if r.Method == http.MethodOptions {
		w.Write([]byte(""))
		return
	}
	if r.Method == http.MethodGet {
		session, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			fmt.Println("NO COOKIE")
			w.Write(getErrorNoCockyAns())
			return
		}
		fmt.Println("COOKIE!!!!!!!!!!!!!!!!!!!")

		uid, ok := db.sessions[session.Value]
		if !ok {
			w.Write(getErrorWrongCookieAns())
			return
		}
		var currentUser *User
		for _, val := range db.users {
			if (*val).id == uid {
				currentUser = val

			}
		}
		fmt.Println("USER::::::", (*currentUser).Password, (*currentUser).Name)
		if (*currentUser).Img == ""{
			fmt.Println("USER HAVE NOT AVATAR")
			w.Write([]byte("USER HAVE NOT AVATAR"))
			return
		}

		file, err := os.Open((*currentUser).Img) // path to image file
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

func setHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	fmt.Println("Origin:::::::::::::::::::::::::::::::::", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Version, Authorization, Content-Type")
	//w.Header().Set("Access-Control-Expose-Headers", "Content-Length, API-Key, Content-Disposition")
}

func main() {
	var db = loggedIn{
		sessions: make(map[string]uint64),
		users:    make(map[string]*User),
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", db.signup)
	mux.HandleFunc("/signin", db.signin)
	mux.HandleFunc("/profile", db.profile)
	mux.HandleFunc("/getAvatar", db.getAvatar)


	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("starting server at :8080")
	server.ListenAndServe()
}
