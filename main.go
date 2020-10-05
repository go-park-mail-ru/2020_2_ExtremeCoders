package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)


const (
	sizeSID=32
)
var sidRunes=[]rune("1234567890_qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

type PostGetter struct {
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
		Description: "Do not require requests expect of POST",
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
	if r.Method != http.MethodPost {
		w.Write(getErrorNotPostAns())
		return
	}
	user, err:=db.users[r.FormValue("login")]
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
	http.SetCookie(w, cocky)
	w.Write(getOkAns(sid))
}

func (db *loggedIn)signup(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		w.Write(getErrorNotPostAns())
		return
	}
	_, err:=db.users[r.FormValue("login")]
	if err{
		w.Write(getErrorLoginExistAns())
		return
	}

	user:=User{
		id: generateUID(db),
		name: r.FormValue("name"),
		surname: r.FormValue("surname"),
		login: r.FormValue("login"),
		password: r.FormValue("password"),
		//birthday: r.FormValue("birthday"),
		//imgPath: r.FormValue("img"),
	}

	sid:=string(generateSID(db))
	db.sessions[sid]= user.id

	cocky:=&http.Cookie{
		Name: "session_id",
		Value: sid,
		Expires: time.Now().Add(24*7*4*time.Hour),
	}
	http.SetCookie(w, cocky)
	w.Write(getOkAns(sid))
}

func (db *loggedIn)updateProfile(changes *PostGetter, uid string) uint16 {
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
	if r.Method != http.MethodPost {
		w.Write(getErrorNotPostAns())
		return
	}
	jsonData:=r.Form
	var change =&PostGetter{}

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
}

func main() {
	var db=loggedIn{
		sessions: make(map[string] uint),
		users: make(map[string] *User),
	}
	http.HandleFunc("/signup", db.signin)
	http.HandleFunc("/signin", db.signup)
	http.HandleFunc("/profile", db.profile)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
