package main

import (
	fileProto "Mailer/FileService/proto"
	"Mailer/MainApplication/config"
	"Mailer/MainApplication/internal/Letter/LetterDelivery"
	"Mailer/MainApplication/internal/Letter/LetterRepository/LetterPostgres"
	"Mailer/MainApplication/internal/Letter/LetterUseCase"
	"Mailer/MainApplication/internal/Postgres"
	"Mailer/MainApplication/internal/User/UserDelivery"
	"Mailer/MainApplication/internal/User/UserRepository/UserPostgres"
	"Mailer/MainApplication/internal/User/UserUseCase"
	"Mailer/MainApplication/internal/pkg/middleware"
	"fmt"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
)

func main() {
	var db = Postgres.DataBase{}
	DataBase, err := db.Init(config.DbUser, config.DbPassword, config.DbDB)
	if err != nil {
		fmt.Println(err)
		return
	}

	grcpConn, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConn.Close()

	fileManager := fileProto.NewFileServiceClient(grcpConn)

	var uDB = UserPostgres.New(DataBase)
	var uUC = UserUseCase.New(uDB)
	var uDE = UserDelivery.New(uUC, fileManager)

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

	//siteHandler := middleware.AccessLogMiddleware(mux)
	//siteHandler = middleware.PanicMiddleware(siteHandler)
	a := middleware.AuthMiddleware{Sessions: uDB}
	siteHandler := a.Auth(mux)
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
	fmt.Println("starting File at ", config.Port)
	err = server.ListenAndServe()
	fmt.Println(err.Error())
}
