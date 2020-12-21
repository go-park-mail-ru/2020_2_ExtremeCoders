package main

import (
	"Mailer/MailService/Postgres"
	"Mailer/MailService/config"
	"Mailer/MailService/internal/Delivery"
	"Mailer/MailService/internal/Repository/LetterPostgres"
	"Mailer/MailService/internal/UseCase"
	letterProto "Mailer/MailService/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

)

func main() {
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalln("cant listen port", err)
	}

	server := grpc.NewServer()
	db := Postgres.DataBase{}
	_, err = db.Init(config.DbUser, config.DbPassword, config.DbDB)
	if err != nil {
		fmt.Println(err)
		return
	}

	repo := LetterPostgres.New(db.DB)
	uc := UseCase.New(repo)
	letterProto.RegisterLetterServiceServer(server, Delivery.New(uc))
	fmt.Println("starting File at :8083")
	server.Serve(lis)
}
