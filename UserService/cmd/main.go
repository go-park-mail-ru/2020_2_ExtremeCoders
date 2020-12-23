package main

import (
	"UserService/Postgres"
	"UserService/config"
	"UserService/internal/UserDelivery"
	"UserService/internal/UserRepository/UserPostgres"
	"UserService/internal/UserUseCase"
	proto "UserService/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalln("cant listen port", err)
	}

	var db = Postgres.DataBase{}
	DataBase, err := db.Init(config.DbUser, config.DbPassword, config.DbDB)
	if err != nil {
		fmt.Println(err)
		return
	}

	server := grpc.NewServer()
	repo := UserPostgres.New(DataBase)
	uc := UserUseCase.New(repo)
	proto.RegisterUserServiceServer(server, UserDelivery.New(uc))

	fmt.Println("starting File at :8082")
	server.Serve(lis)
}
