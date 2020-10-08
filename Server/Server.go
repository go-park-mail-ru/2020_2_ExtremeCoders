package Server

import (
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	sizeSID = 32
)

var sidRunes = []rune("1234567890_qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

type Uinfo struct {
	Email    string
	Password string
}

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
	fmt.Println("SIGNIN GOT: ", r.URL, r.Body)
	if r.Method != http.MethodPost {
		w.Write(getErrorNotPostAns())
		return
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
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(24 * 7 * 4 * time.Hour),
	}
	cookie.Path = "/"
	http.SetCookie(w, cookie)
	w.Write(getOkAns(sid))
}

func (db *loggedIn) signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SIGNUP GOT: ", r.URL, r.Body, r.Method)
	if r.Method != http.MethodPost {
		w.Write(getErrorNotPostAns())
		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var user User
	err := dec.Decode(&user)
	if err != nil {
		w.Write(getErrorBadJsonAns())
		return
	}
	_, erro := db.users[user.Email]
	if erro {
		w.Write(getErrorLoginExistAns())
		return
	}

	user.id = generateUID(db)
	fmt.Println("USER", user.Name, user.Password)

	sid := string(generateSID(db))
	db.sessions[sid] = user.id
	db.users[user.Email] = &user

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Expires: time.Now().Add(24 * 7 * 4 * time.Hour),
	}
	cookie.Path = "/"
	http.SetCookie(w, cookie)
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
	} else if changes.method == "change Img" {
		db.users[uid].Img = changes.value
	}
	return 200
}

func (db *loggedIn) profile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PROFILE GOT: ", r.URL, r.Form, r.Method)
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

		uid, ok := db.sessions[session.Value]
		if !ok {
			w.Write(getErrorWrongCookieAns())
			return
		}
		for _, val := range db.users {
			if (*val).id == uid {
				fmt.Println("If_DATA::::::", (*val).Password, (*val).Name)
				dec := json.NewDecoder(r.Body)
				dec.DisallowUnknownFields()
				type update struct {
					Name    string
					Surname string
				}
				var up update
				err := dec.Decode(&up)
				if err != nil {
					w.Write(getErrorBadJsonAns())
					return
				}
				(*val).Name = up.Name
				(*val).Surname = up.Surname
				w.Write(getOkAns(session.Value))
				return
			}
		}
	}
	w.Write(getErrorUnexpectedAns())
}

func Start() {
	var db = loggedIn{
		sessions: make(map[string]uint64),
		users:    make(map[string]*User),
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", db.signup)
	mux.HandleFunc("/signin", db.signin)
	mux.HandleFunc("/profile", db.profile)
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://95.163.209.195:3000"},
		AllowedHeaders: []string{"Version", "Authorization", "Content-Type"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
	}).Handler(mux)
	server := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("starting server at :8080")
	server.ListenAndServe()
}
