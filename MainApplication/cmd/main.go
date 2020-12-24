package main

import (
	protoFs "Mailer/FileService/proto"
	protoMail "Mailer/MailService/proto"
	"Mailer/config"
	"Mailer/MainApplication/internal/Folder/FolderDelivery"
	"Mailer/MainApplication/internal/Letter/LetterDelivery"
	"Mailer/MainApplication/internal/Letter/LetterRepository/LetterService"
	"Mailer/MainApplication/internal/Letter/LetterUseCase"
	"Mailer/MainApplication/internal/User/UserDelivery"
	"Mailer/MainApplication/internal/User/UserRepository/UserMicroservice"
	"Mailer/MainApplication/internal/User/UserUseCase"
	"Mailer/MainApplication/internal/pkg/middleware"
	protoUs "Mailer/UserService/proto"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/swaggo/http-swagger"
    _ "Mailer/docs"
	"google.golang.org/grpc"
	"net/http"
)


// @title Mailer Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email martin7.heinz@gmail.com

// @license.name MIT
// @license.url https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

// @BasePath /api/v1

func main() {
	grcpMailService, err := grpc.Dial(
		"127.0.0.1:8083",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc mail service")
	}
	defer grcpMailService.Close()
	mailManager := protoMail.NewLetterServiceClient(grcpMailService)

	grcpFileService, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc file service")
	}
	defer grcpFileService.Close()
	fileManager := protoFs.NewFileServiceClient(grcpFileService)

	grcpUserService, err := grpc.Dial(
		"127.0.0.1:8082",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc file service")
	}
	defer grcpUserService.Close()
	userManager := protoUs.NewUserServiceClient(grcpUserService)

	var uDB = UserMicroservice.New(userManager)
	var uUC = UserUseCase.New(uDB)
	var uDE = UserDelivery.New(uUC, fileManager)

	var lDB = LetterService.New(mailManager)
	var lUC = LetterUseCase.New(lDB)
	var lDE = LetterDelivery.New(lUC)

	var fDe = FolderDelivery.New(userManager, mailManager)

	mux := mux.NewRouter()

	//post session body {email:"email@mail.ru", password: "password"}- логин
	//delete session - логаут
	mux.HandleFunc("/api/session", uDE.Session)

	//get user - информация о пользователе
	//post user {name:"Sam", surname: "Potter", email: "img", password1: "password1", avatar: "img"}- регистрация
	//put user {profile_firstName:"Sam", profile_lastName: "Potter", avatar: "img"} - изменить профиль
	mux.HandleFunc("/api/user", uDE.Profile)

	//get user/avatar - получить аватарку
	mux.HandleFunc("/api/user/avatar", uDE.GetAvatar)

	//delete letter - удалить письмо
	//post letter {to:'receiver', theme:'theme', text:'letter content'}- отправить письмо
	mux.HandleFunc("/api/letter", lDE.SendLetter)

	//get user/letter/sent/{limit}/{offset} - получить отправленные письма
	mux.HandleFunc("/api/user/letter/sent/{limit}/{offset}", lDE.GetSendLetters)


	//get user/letter/sent/{limit}/{offset} - получить принятые письма
	mux.HandleFunc("/api/user/letter/received/{limit}/{offset}", lDE.GetRecvLetters)

	//get letter/{similar} - поиск по всем элементам
	mux.HandleFunc("/api/letter/{similar}", lDE.Search)

	//put watch/letter {id:10} - пометить письмо как прочитанное
	mux.HandleFunc("/api/watch/letter", lDE.WatchLetter)

	mux.HandleFunc("/api/letter/box/set", lDE.SetLetterInBox)

	mux.HandleFunc("/api/letter/spam/set", lDE.SetLetterInSpam)

	//get letter/by/{what}/{value} - what может быть равен
	//(id, sender, receiver, theme, text, date_time, directory_recv, directory_send) - поиск по письмам
	mux.HandleFunc("/api/letter/by/{what}/{value}", lDE.GetLetterBy)

	//get user/folders/received - список папок в отправленных
	mux.HandleFunc("/api/user/folders/received", fDe.GetFolderList)

	//get user/folders/received - список папок в полученных
	mux.HandleFunc("/api/user/folders/sended", fDe.GetFolderList)


	//get user/foders/{received/sended}/{folderName} - письма из папки в полученых, письма из папки в отправленнх
	mux.HandleFunc("/api/user/foders/received/{folderName}/{limit}/{offset}", fDe.GetLettersByFolder)

	//get user/foders/sended/{folderName} - письма из папки в отправленнх
	mux.HandleFunc("/api/user/foders/sended/{folderName}", fDe.GetLettersByFolder)

	//post user/folders/received/folderName {folderName:"folderName"} - добавить папку в полученные
	//put user/folders/received/folderName {oldName:"oldName", newName"":} - переименовать папку в полученных
	mux.HandleFunc("/api/user/folders/received/folderName", fDe.AddFolder)

	//post user/folders/sended/folderName {folderName:"folderName"} - добавить папку в отправленные
	//put user/folders/sended/folderName {oldName:"oldName", newName:"newName"} - переименовать папку в отправленных
	//delete user/folders/{sended, recived}/folderName {folderName:"folderName"} - удалить папку
	mux.HandleFunc("/api/user/folders/sended/folderName", fDe.AddFolder)

	//put user/folders/received/folderName/letter body{letterID: id} - добавить письмо в папку из полученных
	//delete /user/folders/received/folderName/letter body{letterID:Id} - удалить письмо из папки в полученных
	mux.HandleFunc("/api/user/folders/received/{folderName}/letter", fDe.AddLetterInFolder)

	//put user/folders/sended/folderName/letter body{letterID: id} - добавить письмо в папку из отправленных
	//delete /user/folders/sended/folderName/letter body{letterID:Id} - удалить письмо из папки в отправленных
	mux.HandleFunc("/api/user/folders/sended/folderName/letter", fDe.AddLetterInFolder)

	mux.PathPrefix("/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL("https://localhost:8080/docs/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
	//mux.Handle("/metrics", promhttp.Handler())
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
	fmt.Println("starting Main at ", config.Port)
	err = server.ListenAndServe()
	fmt.Println(err.Error())
}
