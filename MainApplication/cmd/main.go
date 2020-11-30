package main

import (
	"MainApplication/config"
	"MainApplication/internal/Folder/FolderDelivery"
	"MainApplication/internal/Letter/LetterDelivery"
	"MainApplication/internal/Letter/LetterRepository/LetterService"
	"MainApplication/internal/Letter/LetterUseCase"
	"MainApplication/internal/User/UserDelivery"
	"MainApplication/internal/User/UserRepository/UserMicroservice"
	"MainApplication/internal/User/UserUseCase"
	"MainApplication/internal/pkg/middleware"
	protoFs "MainApplication/proto/FileServise"
	protoMail "MainApplication/proto/MailService"
	protoUs "MainApplication/proto/UserServise"
	"fmt"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
)

func main() {
	//var db = Postgres.DataBase{}
	//DataBase, err := db.Init(config.DbUser, config.DbPassword, config.DbDB)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

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

	var fDe = FolderDelivery.New()
	mux := http.NewServeMux()

	mux.HandleFunc("/session", uDE.Session)
	mux.HandleFunc("/user", uDE.Profile)
	mux.HandleFunc("/user/avatar", uDE.GetAvatar)
	mux.HandleFunc("/letter", lDE.SendLetter)
	mux.HandleFunc("/user/letter/sent", lDE.GetSendLetters)
	mux.HandleFunc("/user/letter/received", lDE.GetRecvLetters)
	mux.HandleFunc("/watch/letter", lDE.WatchLetter)
	//get /user/folders/{recived/sended} - список папок
	mux.HandleFunc("/user/folders/recived", fDe.GetFolderList)
	mux.HandleFunc("/user/folders/sended", fDe.GetFolderList)
	//get /user/foders/{recived/sended}/folderName - письма
	mux.HandleFunc("/user/foders/recived/folderName", fDe.GetLettersByFolder)
	mux.HandleFunc("/user/foders/sended/folderName", fDe.GetLettersByFolder)
	//post /user/folders/{recived/sended}/folderName - добавить папку
	mux.HandleFunc("/user/folders/recived/folderName", fDe.AddFolder)
	mux.HandleFunc("/user/folders/sended/folderName", fDe.AddFolder)
	//put /user/folders/{recived/sended}/folderName/letter body{letterID: id} - добавить письмо в папку
	mux.HandleFunc("/user/folders/recived/folderName/letter", fDe.RenameFolder)
	mux.HandleFunc("/user/folders/sended/folderName/letter", fDe.AddLetterInFolder)
	//put /user/folders/{recived/sended}/folderName body:{ name: newName} - переименовать папку
	mux.HandleFunc("/user/folders/recived/folderName ", fDe.AddLetterInFolder)
	mux.HandleFunc("/user/folders/sended/folderName ", fDe.RenameFolder)
	//delete /user/folders/{recived/sended}/folderName/letter body{letterID:Id} - удалить письмо из папки
	mux.HandleFunc("/user/folders/recived/folderName/letter  ", fDe.RemoveLetterInFolder)
	mux.HandleFunc("/user/folders/sended/folderName/letter  ", fDe.RemoveLetterInFolder)
	//delete /user/folders/{recived/sended}/folderName  - удалить папку
	//mux.HandleFunc("/user/folders/recived/folderName " , fDe.RemoveFolder)
	//mux.HandleFunc("/user/folders/sended/folderName/letter" , fDe.RemoveFolder)

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
