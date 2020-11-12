package middleware

import (
	"CleanArch/internal/User/UserRepository"
	"CleanArch/internal/pkg/context"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type RequestLog struct {
	addr string
	method string
	url string
}

var csrfError=errors.New("Sorry but your csrf token is over and someone can still your account. Please restart the page!")

func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.WithFields(log.Fields{
			"REQUEST": RequestLog{
				addr: r.RemoteAddr,
				method: r.Method,
				url: r.URL.Path,
			},
		}).Info("got")

	})
}

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.WithFields(log.Fields{
					"RECOVERED": err,
				}).Error("got")
				http.Error(w, "Internal server error", 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

type AuthMiddleware struct {
	Sessions UserRepository.UserDB
}

const (
	cookieName     = "session_id"
	csrfCookieName = "token"
)

func (a AuthMiddleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cookieName)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		csrf, errC := r.Cookie(csrfCookieName)
		if csrf.Value == r.Header.Get("csrf_token") {
			cookie := &http.Cookie{
				Name:    csrfCookieName,
				Value:   context.GenerateCSRF(),
				Expires: time.Now().Add(15 * time.Minute),
			}
			cookie.Path = "/"
			http.SetCookie(w, cookie)
		}
		if errC != nil {
			w.Write(authError(csrfError))
			log.WithFields(log.Fields{
				"RECOVERED": csrfError,
			}).Error("got")
			return
		}

		uid, er := a.Sessions.IsOkSession(cookie.Value)
		if er != nil {
			next.ServeHTTP(w, r)
			return
		}
		user, e := a.Sessions.GetUserByUID(uid)
		if e != nil {
			next.ServeHTTP(w, r)
			return
		}
		ctx := r.Context()
		ctx = context.SaveUserToContext(ctx, *user)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}