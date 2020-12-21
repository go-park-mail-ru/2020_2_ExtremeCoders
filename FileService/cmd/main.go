package main

import (
	"Mailer/FileService/internal/File/Delivery"
	fsRepo "Mailer/FileService/internal/File/Repository/FileSystem"
	"Mailer/FileService/internal/File/UseCase"
	fileProto "Mailer/FileService/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("cant listen port", err)
	}

	server := grpc.NewServer()
	repo := fsRepo.New()
	uc := UseCase.New(repo)
	fileProto.RegisterFileServiceServer(server, Delivery.NewFileManager(uc))

	fmt.Println("starting File at :8081")
	server.Serve(lis)
}
