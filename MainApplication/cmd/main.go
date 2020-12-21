package main

import (
	protoFs "Mailer/FileService/proto"
	protoMail "Mailer/MailService/proto"
	"Mailer/MainApplication/internal/Folder/FolderDelivery"
	"Mailer/MainApplication/internal/Letter/LetterDelivery"
	"Mailer/MainApplication/internal/Letter/LetterRepository/LetterService"
	"Mailer/MainApplication/internal/Letter/LetterUseCase"
	"Mailer/MainApplication/internal/User/UserDelivery"
	"Mailer/MainApplication/internal/User/UserRepository/UserMicroservice"
	"Mailer/MainApplication/internal/User/UserUseCase"
	"Mailer/MainApplication/internal/pkg/middleware"
	protoUs "Mailer/UserService/proto"
	"Mailer/config"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

type tmp struct {
	ID int
}

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email martin7.heinz@gmail.com

// @license.name MIT
// @license.url https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

// @BasePath /api/v1

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

<<<<<<< HEAD
func main() {
=======
// ShowAccount godoc
// @Summary Show a account
// @Description get user by ID
// @ID get-user-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} tmp
// @Router /user/{id} [get]
func handleUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status": "ok"}`))
}
>>>>>>> CleanArch

func main() {
	grcpMailService, err := grpc.Dial(
		"127.0.0.1:8083",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc mail service")
	}
	defer func() { _ = grcpMailService.Close() }()

	mailManager := protoMail.NewLetterServiceClient(grcpMailService)

	grcpFileService, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc file service")
	}
	defer func() { _ = grcpFileService.Close() }()
	fileManager := protoFs.NewFileServiceClient(grcpFileService)

	grcpUserService, err := grpc.Dial(
		"127.0.0.1:8082",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc file service")
	}
	defer func() { _ = grcpUserService.Close() }()
	userManager := protoUs.NewUserServiceClient(grcpUserService)

	var uDB = UserMicroservice.New(userManager)
	var uUC = UserUseCase.New(uDB)
	var uDE = UserDelivery.New(uUC, fileManager)

	var lDB = LetterService.New(mailManager)
	var lUC = LetterUseCase.New(lDB)
	var lDE = LetterDelivery.New(lUC)

	var fDe = FolderDelivery.New(userManager, mailManager)

	mux := mux.NewRouter()

	//post session {email:"email@mail.ru", password: "password"}- логин
	//delete session - логаут
	mux.HandleFunc("/session", uDE.Session)

	//get user - информация о пользователе
	//post user {name:"Sam", surname: "Potter", email: "img", password1: "password1", avatar: "img"}- регистрация
	//put user {profile_firstName:"Sam", profile_lastName: "Potter", avatar: "img"} - изменить профиль
	mux.HandleFunc("/user", uDE.Profile)

	//get user/avatar - получить аватарку
	mux.HandleFunc("/user/avatar", uDE.GetAvatar)

	//delete letter - удалить письмо
	//post letter {to:'receiver', theme:'theme', text:'letter content'}- отправить письмо
	mux.HandleFunc("/letter", lDE.SendLetter)
<<<<<<< HEAD
	mux.HandleFunc("/user/letter/sent", lDE.GetSendLetters)
	mux.HandleFunc("/user/letter/received", lDE.GetRecvLetters)
=======

	//get user/letter/sent/{limit}/{offset} - получить отправленные письма
	mux.HandleFunc("/user/letter/sent/{limit}/{offset}", lDE.GetSendLetters)

	//get user/letter/sent/{limit}/{offset} - получить принятые письма
	mux.HandleFunc("/user/letter/received/{limit}/{offset}", lDE.GetRecvLetters)

	//get letter/{similar} - поиск по всем элементам
>>>>>>> CleanArch
	mux.HandleFunc("/letter/{similar}", lDE.Search)

	//put watch/letter {id:10} - пометить письмо как прочитанное
	mux.HandleFunc("/watch/letter", lDE.WatchLetter)

	//get letter/by/{what}/{value} - what может быть равен
	//(id, sender, receiver, theme, text, date_time, directory_recv, directory_send) - поиск по письмам
	mux.HandleFunc("/letter/by/{what}/{value}", lDE.GetLetterBy)

	//get user/folders/recived - список папок в отправленных
	mux.HandleFunc("/user/folders/recived", fDe.GetFolderList)

	//get user/folders/recived - список папок в полученных
	mux.HandleFunc("/user/folders/sended", fDe.GetFolderList)
<<<<<<< HEAD
	//get /user/foders/{recived/sended}/folderName - письма
	mux.HandleFunc("/user/foders/recived/{folderName}", fDe.GetLettersByFolder)
=======

	//get user/foders/{recived/sended}/{folderName} - письма из папки в полученых, письма из папки в отправленнх
	mux.HandleFunc("/user/foders/recived/{folderName}/{limit}/{offset}", fDe.GetLettersByFolder)

	//get user/foders/sended/{folderName} - письма из папки в отправленнх
>>>>>>> CleanArch
	mux.HandleFunc("/user/foders/sended/{folderName}", fDe.GetLettersByFolder)

	//post user/folders/recived/folderName {folderName:"folderName"} - добавить папку в полученные
	//put user/folders/recived/folderName {oldName:"oldName", newName"":} - переименовать папку в полученных
	mux.HandleFunc("/user/folders/recived/folderName", fDe.AddFolder)

	//post user/folders/sended/folderName {folderName:"folderName"} - добавить папку в отправленные
	//put user/folders/sended/folderName {oldName:"oldName", newName:"newName"} - переименовать папку в отправленных
	mux.HandleFunc("/user/folders/sended/folderName", fDe.AddFolder)

	//put user/folders/recived/folderName/letter body{letterID: id} - добавить письмо в папку из полученных
	//delete /user/folders/recived/folderName/letter body{letterID:Id} - удалить письмо из папки в полученных
	mux.HandleFunc("/user/folders/recived/folderName/letter", fDe.AddLetterInFolder)

	//put user/folders/sended/folderName/letter body{letterID: id} - добавить письмо в папку из отправленных
	//delete /user/folders/sended/folderName/letter body{letterID:Id} - удалить письмо из папки в отправленных
	mux.HandleFunc("/user/folders/sended/folderName/letter", fDe.AddLetterInFolder)
<<<<<<< HEAD
	//put /user/folders/{recived/sended}/folderName body:{ name: newName} - переименовать папку
	mux.HandleFunc("/user/folders/recived/folderName ", fDe.RenameFolder)
	mux.HandleFunc("/user/folders/sended/folderName ", fDe.RenameFolder)
	//delete /user/folders/{recived/sended}/folderName/letter body{letterID:Id} - удалить письмо из папки
	mux.HandleFunc("/user/folders/recived/folderName/letter  ", fDe.RemoveLetterInFolder)
	mux.HandleFunc("/user/folders/sended/folderName/letter  ", fDe.RemoveLetterInFolder)
=======

	mux.HandleFunc("/api", httpSwagger.WrapHandler)
>>>>>>> CleanArch

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
