package main

import (
	"CleanArch/config"
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
)

func main() {
	var db = Postgres.DataBase{}
	DataBase, err := db.Init(config.DbUser, config.DbPassword, config.DbDB)
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
		AllowedOrigins:   config.AllowedOriginsCORS,
		AllowedHeaders:   config.AllowedHeadersCORS,
		AllowedMethods:   config.AllowedMethodsCORS,
		AllowCredentials: true,
	}).Handler(mux)
	siteHandler := middleware.AccessLogMiddleware(handler)
	siteHandler = middleware.PanicMiddleware(siteHandler)
	server := http.Server{
		Addr:         config.Port,
		Handler:      siteHandler,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}
	fmt.Println("starting server at :8080")
	server.ListenAndServe()
}
