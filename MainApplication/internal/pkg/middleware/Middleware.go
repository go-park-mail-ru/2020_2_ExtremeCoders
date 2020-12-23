package middleware

import (
	"MainApplication/internal/User/UserRepository"
	"MainApplication/internal/pkg/context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type RequestLog struct {
	addr   string
	method string
	url    string
}

var csrfError = errors.New("Sorry but your csrf token is over and someone can still your account. Please restart the page!")

func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.WithFields(log.Fields{
			"REQUEST": RequestLog{
				addr:   r.RemoteAddr,
				method: r.Method,
				url:    r.URL.Path,
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
				http.Error(w, "Internal File error", 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

type AuthMiddleware struct {
	Sessions UserRepository.UserDB
}

func (a AuthMiddleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(context.CookieName)
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
		if r.Method==http.MethodGet{
			http.SetCookie(w, context.CreateCsrfCookie())
			ctx := r.Context()
			ctx = context.SaveUserToContext(ctx, *user)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
			return
		} else{
			csrf, Error := r.Cookie(context.CsrfCookieName)
			fmt.Printf("REQ", r.URL.Path)
			//если пришли с нормальным csrf, то обновляем его, получаем юзера и прокидываем запрос дальше
			//fmt.Printf("%s == %s\n", csrf.Value, r.Header.Get("csrf_token"))
			if csrf != nil && csrf.Value == r.Header.Get("csrf_token") {
				if Error == nil {
					csrf.Expires = time.Now().AddDate(0, 0, -1)
					http.SetCookie(w, csrf)
				}
				http.SetCookie(w, context.CreateCsrfCookie())
				ctx := r.Context()
				ctx = context.SaveUserToContext(ctx, *user)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			} else {
				//если csrf не норм, то если это вход или регистрация, то надо отправить на них
				if (r.URL.Path == "/api/session" || r.URL.Path == "/api/user") && r.Method == http.MethodPost {
					if Error == nil {
						csrf.Expires = time.Now().AddDate(0, 0, -1)
						http.SetCookie(w, csrf)
					}
					http.SetCookie(w, context.CreateCsrfCookie())
					next.ServeHTTP(w, r)
					return
				}
				//либо злоумышленник либо что-то пошло совсем не так...
				w.Write(authError(csrfError))
				log.WithFields(log.Fields{
					"RECOVERED": csrfError,
				}).Error("got")
			}
		}


	})
}
