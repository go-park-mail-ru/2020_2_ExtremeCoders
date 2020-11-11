package middleware

import (
	"CleanArch/internal/User/UserRepository"
	"CleanArch/internal/pkg/context"
	"fmt"
	"net/http"
	"time"
)

func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("REQUEST: [%s] %s, %s %s\n",
			r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
	})
}

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
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
	cookieName = "session_id"
)

func (a AuthMiddleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cookieName)
		if err==http.ErrNoCookie{
			var uid uint64
			cookie.
			if cookie!=nil {
				_, uid=a.Sessions.RemoveSession(cookie.Value)
				sid, _:=a.Sessions.GenerateSID()
				user, _:=a.Sessions.GetUserByUID(uid)
				a.Sessions.AddSession(string(sid), uid, user)
				cookie := &http.Cookie{
					Name:    "session_id",
					Value:   string(sid),
					//Expires: time.Now().Add(24 * 7 * 4 * time.Hour),
					Expires: time.Now().Add(1* time.Second),
				}
				cookie.Path = "/"
				http.SetCookie(w, cookie)
			}
		}
		if err != nil {
			next.ServeHTTP(w, r)
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
