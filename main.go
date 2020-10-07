package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)


const (
	sizeSID=32
)
var sidRunes=[]rune("1234567890_qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

type Profile struct {
	method string
	value string
}

type User struct {
	id       uint64
	name     string
	surname  string
	login    string
	password string
	birthday uint64
	imgPath  string
}

type Answer struct {
	Code        uint16
	Description string
	sid string
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

func getErrorWrongCookieAns() []byte{
	err:= &Answer{
		Code:401,
		Description: "wrong session id",
	}
	ans, _:=json.Marshal(err)
	return ans
}

func getErrorNotPostAns() []byte{
	err:= &Answer{
		Code:400,
		Description: "Do not require request's method, expected POST",
	}
	ans, _:=json.Marshal(err)
	return ans
}

func getErrorNotNumberAns() []byte{
	err:= &Answer{
		Code:400,
		Description: "Not number",
	}
	ans, _:=json.Marshal(err)
	return ans
}

func getErrorNoUserAns() []byte{
	err:= &Answer{
		Code:404,
		Description: "Do not find this user in db",
	}
	ans, _:=json.Marshal(err)
	return ans
}

func getErrorLoginExistAns() [] byte{
	err:= &Answer{
		Code:401,
		Description: "This login has already exists",
	}
	ans, _:=json.Marshal(err)
	return ans
}

func getErrorBadPasswordAns() []byte{
	err:= &Answer{
		Code:401,
		Description: "Wrong password",
	}
	ans, _:=json.Marshal(err)
	return ans
}

func getOkAns(cocky string) []byte{
	ok:= &Answer{
		Code:200,
		Description: "ok",
		sid: cocky,
	}
	ans, _:=json.Marshal(ok)
	return ans
}

func getOkAnsData(cocky string, data User) []byte{
	fmt.Println("DATA::::::::::", data.login, data.name, data.password)
	ok:= &Answer{
		Code:200,
		Description: "ok",
		sid: cocky,
		User: data,
	}
	ans, _:=json.Marshal(ok)
	return ans
}

func generateSID(db *loggedIn) []rune {
	var sid=make ([]rune, sizeSID)
	for{
		for i:=0;i<sizeSID;i++{
			sid[i]=sidRunes[rand.Intn(len(sidRunes))]
		}
		_,exist:=db.sessions[string(sid)]
		if !exist{
			break
		}
	}
	return sid
}

func generateUID(db *loggedIn) uint64{
	var uid uint64
	for{
		for i:=0;i<sizeSID;i++{
			uid=rand.Uint64()
		}
		var _, exist = db.users[strconv.FormatUint(uid, 10)]
		if !exist{
			break
		}
	}
	return uid
}

func (db *loggedIn)signin(w http.ResponseWriter, r *http.Request){
	setHeader(w, r)
	fmt.Println("SIGNIN GOT: ", r.URL, r.Body)
	fmt.Println("USER", r.FormValue("email"), r.FormValue("password"))
	if r.Method != http.MethodPost {
		w.Write(getErrorNotPostAns())
		return
	}
	user, err:=db.users[r.FormValue("email")]
	if !err{
		w.Write(getErrorNoUserAns())
		return
	}
	if user.password!=r.FormValue("password"){
		w.Write(getErrorBadPasswordAns())
		return
	}
	sid:=string(generateSID(db))
	db.sessions[sid]= user.id
	cocky:=&http.Cookie{
		Name: "session_id",
		Value: sid,
		Expires: time.Now().Add(24*7*4*time.Hour),
	}
	cocky.Path="/"
	http.SetCookie(w, cocky)
	w.Write(getOkAns(sid))
}



func (db *loggedIn)signup(w http.ResponseWriter, r *http.Request){
	body:=r.PostFormValue("body_form")
	setHeader(w, r)
	fmt.Println("SIGNUP GOT: ", r.URL, r.Body, r.Method)
	if r.Method==http.MethodOptions{
		w.Write([]byte(""))
		return
	}
	if r.Method != http.MethodPost {
		w.Write(getErrorNotPostAns())
		return
	}
	//body:=r.PostFormValue("body_form")
	//body=r.Form.Get("body_form")
	_, err:=db.users[gjson.Get(body,"email").String()]
	if err {
		w.Write(getErrorLoginExistAns())
		return
	}
	user:=User{
		id: generateUID(db),
		name: gjson.Get(body,"name").String(),
		surname: gjson.Get(body,"surname").String(),
		login: gjson.Get(body,"email").String(),
		password: gjson.Get(body,"password").String(),
		//birthday: gjson.Get(body,"date"),
		imgPath: gjson.Get(body,"img").String(),
	}
	fmt.Println("USER", user.name, user.password)

	sid:=string(generateSID(db))
	db.sessions[sid]= user.id
	db.users[user.login]=&user

	cocky:=&http.Cookie{
		Name: "session_id",
		Value: sid,
		Expires: time.Now().Add(24*7*4*time.Hour),
	}
	cocky.Path="/"
	http.SetCookie(w, cocky)
	w.Write(getOkAns(sid))
}

func (db *loggedIn)updateProfile(changes *Profile, uid string) uint16 {
	if changes.method=="change password" {
		db.users[uid].password=changes.value
	}else if changes.method=="change login"{
		db.users[uid].login=changes.value
	}else if changes.method=="change name"{
		db.users[uid].name=changes.value
	}else if changes.method=="change surname"{
		db.users[uid].surname=changes.value
	}else if changes.method=="change birthday"{
		date, err := strconv.ParseUint(changes.value, 10, 64)
		if err != nil{
			return 400
		}
		db.users[uid].birthday=date
	}else if changes.method=="change img"{
		db.users[uid].imgPath = changes.value
	}
	return 200
}

func (db *loggedIn)profile(w http.ResponseWriter, r *http.Request){
	setHeader(w, r)
	fmt.Println("PROFILE GOT: ", r.URL, r.Form, r.Method)
	if r.Method==http.MethodOptions{
		w.Write([]byte(""))
		return
	}
	if r.Method == http.MethodGet {
		session, err := r.Cookie("session_id")
		if err==http.ErrNoCookie {
			fmt.Println("NO COOKIE")
			w.Write(getErrorNoCockyAns())
			return
		}
		fmt.Println("COOKIE!!!!!!!!!!!!!!!!!!!")

		uid, ok := db.sessions[session.Value]
		if  !ok {
			w.Write(getErrorWrongCookieAns())
			return
		}
		for _, val:=range db.users{
			if (*val).id==uid{
				fmt.Println("If_DATA::::::", (*val).password, (*val).name)
				w.Write(getOkAnsData(session.Value, *val))
				return
			}
		}

	}else if r.Method != http.MethodPost{
		w.Write(getErrorNotPostAns())
		return
	}else{
		jsonData:=r.Form
		var change =&Profile{}

		session, err := r.Cookie("session_id")

		if err!=nil {
			w.Write(getErrorNoCockyAns())
			return
		}
		if _, ok := db.sessions[session.Value]; !ok {
			w.Write(getErrorWrongCookieAns())
			return
		}
		err=json.Unmarshal([]byte(jsonData.Encode()), change)

		if err!=nil{
			w.Write(getErrorBadJsonAns())
			return
		}
		ans:=db.updateProfile(change, session.Value)
		if ans==400{
			w.Write(getErrorNotNumberAns())
			return
		}
		w.Write(getOkAns(session.Value))
		return
	}
	w.Write(getErrorUnexpectedAns())
}


func setHeader(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Version, Authorization, Content-Type")
	//w.Header().Set("Access-Control-Expose-Headers", "Content-Length, API-Key, Content-Disposition")
}
func main() {
	var db=loggedIn{
		sessions: make(map[string]uint64),
		users: make(map[string]*User),
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", db.signup)
	mux.HandleFunc("/signin", db.signin)
	mux.HandleFunc("/profile", db.profile)
	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("starting server at :8080")
	server.ListenAndServe()
}
