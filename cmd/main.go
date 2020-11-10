package main

import (
	"CleanArch/internal/Letter/LetterDelivery"
	"CleanArch/internal/Letter/LetterRepository/LetterPostgres"
	"CleanArch/internal/Letter/LetterUseCase"
	"CleanArch/internal/Postgres"
	"CleanArch/internal/User/UserDelivery"
	"CleanArch/internal/User/UserRepository/UserPostgres"
	"CleanArch/internal/User/UserUseCase"
	"CleanArch/internal/pkg/middleware"
	"fmt"
	"github.com/rs/cors"
	"net/http"
	"time"
)

func main() {
	var db = Postgres.DataBase{}
	DataBase, err := db.Init("postgres", "1538", "maila")
	if err != nil {
		fmt.Println(err)
		return
	}
	var uDB = UserPostgres.DataBase{DB: DataBase}
	var uUC = UserUseCase.UseCase{Db: uDB}
	var uDE = UserDelivery.Delivery{Uc: uUC}

	var lDB = LetterPostgres.DataBase{DB: DataBase}
	var lUC = LetterUseCase.UseCase{Db: lDB}
	var lDE = LetterDelivery.Delivery{Uc: lUC}
	mux := http.NewServeMux()

	mux.HandleFunc("/session", uDE.Session)
	mux.HandleFunc("/user", uDE.Profile)
	mux.HandleFunc("/user/avatar", uDE.GetAvatar)
	mux.HandleFunc("/letter", lDE.SendLetter)
	mux.HandleFunc("/user/letter/sent", lDE.GetSendLetters)
	mux.HandleFunc("/user/letter/received", lDE.GetRecvLetters)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000", "http://95.163.209.195:3000"},
		AllowedHeaders:   []string{"Version", "Authorization", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}).Handler(mux)
	siteHandler := middleware.AccessLogMiddleware(handler)
	siteHandler = middleware.PanicMiddleware(siteHandler)
	server := http.Server{
		Addr:         ":8080",
		Handler:      siteHandler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("starting server at :8080")
	server.ListenAndServe()
}
