package middleware

import (
	"CleanArch/internal/User/UserRepository"
	"CleanArch/internal/pkg/context"
	"net/http"
)

func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
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
		if err != http.ErrNoCookie {
		}
		uid, err := a.Sessions.IsOkSession(cookie.Value)
		if err != nil {
		}
		user, e := a.Sessions.GetUserByUID(uid)
		if e != nil {
		}
		ctx := r.Context()
		ctx = context.SaveUserToContext(ctx, *user)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
