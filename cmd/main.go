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
	var uDB = UserPostgres.New(DataBase)
	var uUC = UserUseCase.New(uDB)
	var uDE = UserDelivery.New(uUC)

	var lDB = LetterPostgres.New(DataBase)
	var lUC = LetterUseCase.New(lDB)
	var lDE = LetterDelivery.New(lUC)


	mux := http.NewServeMux()

	mux.HandleFunc("/session", uDE.Session)
	mux.HandleFunc("/user", uDE.Profile)
	mux.HandleFunc("/user/avatar", uDE.GetAvatar)
	mux.HandleFunc("/letter", lDE.SendLetter)
	mux.HandleFunc("/user/letter/sent", lDE.GetSendLetters)
	mux.HandleFunc("/user/letter/received", lDE.GetRecvLetters)

	siteHandler := middleware.AccessLogMiddleware(mux)
	siteHandler = middleware.PanicMiddleware(siteHandler)
	a := middleware.AuthMiddleware{Sessions: uDB}
	siteHandler = a.Auth(siteHandler)
	handler := cors.New(cors.Options{
		AllowedOrigins:   config.AllowedOriginsCORS,
		AllowedHeaders:   config.AllowedHeadersCORS,
		AllowedMethods:   config.AllowedMethodsCORS,
		AllowCredentials: true,
	}).Handler(siteHandler)
	server := http.Server{
		Addr:         config.Port,
		Handler:      handler,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}
	fmt.Println("starting server at ", config.Port)
	err = server.ListenAndServe()
	fmt.Println(err.Error())
}
