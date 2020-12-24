package main

import (
	"Mailer/SmtpService/internal/GetLetters"
	"Mailer/SmtpService/internal/SendLetters"
	pb "Mailer/SmtpService/proto"
	"fmt"
	"github.com/emersion/go-smtp"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

var wg sync.WaitGroup
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	be := &GetLetters.Backend{}
	s := smtp.NewServer(be)
	s.Addr = ":25"
	s.Domain = "mx.mailer.ru"
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	fmt.Println("Starting server at", s.Addr)
	wg.Add(1)
	go func() {
		_ = s.ListenAndServe()
	}()




	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("cant listen port", err)
	}
	server := grpc.NewServer()
	pb.RegisterLetterServiceServer(server, SendLetters.NewSMTPManager())
	fmt.Println("starting File at :8080")
	_ = server.Serve(lis)
	wg.Wait()
}
