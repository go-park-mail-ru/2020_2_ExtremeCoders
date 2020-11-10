package main

import (
	"CleanArch/app/User/UserDelivery"
	"CleanArch/app/User/UserRepository/UserPostgres"
	"CleanArch/app/User/UserUseCase"
	"fmt"
	"github.com/rs/cors"
	"net/http"
	"time"
)


func accessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("accessLogMiddleware", r.URL.Path)
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("REQUEST: [%s] %s, %s %s\n",
			r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
	})
}


func panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("panicMiddleware", r.URL.Path)
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered", err)
				http.Error(w, "Internal server error", 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	var db= UserPostgres.DataBase{ User: "postgres", Password: "1538", DataBaseName: "maila"}
	db.Init()
	var uc= UserUseCase.UseCase{Db: db}
	var yaFood= UserDelivery.Delivery{Uc: uc}
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", yaFood.Signup)
	mux.HandleFunc("/signin", yaFood.SignIn)
	mux.HandleFunc("/profile", yaFood.Profile)
	mux.HandleFunc("/logout", yaFood.Logout)
	mux.HandleFunc("/getAvatar", yaFood.GetAvatar)
	mux.HandleFunc("/sendMessage", yaFood.SendLetter)
	mux.HandleFunc("/getSendLetters", yaFood.GetSendLetters)
	mux.HandleFunc("/getRecvLetters", yaFood.GetRecvLetters)



	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://127.0.0.1:3000", "http://95.163.209.195:3000"},
		AllowedHeaders: []string{"Version", "Authorization", "Content-Type"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
	}).Handler(mux)
	siteHandler := accessLogMiddleware(handler)
	siteHandler = panicMiddleware(siteHandler)
	server := http.Server{
		Addr:         ":8080",
		Handler:      siteHandler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("starting server at :8080")
	server.ListenAndServe()
}
