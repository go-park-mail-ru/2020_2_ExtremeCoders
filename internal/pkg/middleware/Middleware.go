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
				fmt.Println("recovered", err)
				http.Error(w, "Internal server error", 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

type authMiddleware struct {
	sessions UserRepository.UserDB
}

type userKey struct {}

const(
	cookieName="session_id"
)

func (a authMiddleware) auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cookieName)
		if err != http.ErrNoCookie {
			// 401
		}
		uid, ok := a.sessions.IsOkSession(cookie.Value)
		if !ok {
			// херово
		}
		user := a.sessions.GetUserByUID(uid)
		ctx := r.Context()
		context.SaveUserToContext(ctx, *user)
		r.WithContext(ctx)
		next.ServeHTTP(w, r)

	})
}