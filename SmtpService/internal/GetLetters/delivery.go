package GetLetters

import (
	server "Mailer/MailService/proto"
	send "Mailer/SmtpService/internal/SendLetters"
	"context"
	"fmt"
	"github.com/emersion/go-smtp"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"strings"
)

// The Backend implements SMTP server methods.
type Backend struct{}

// Login handles a login command with username and password.
func (bkd *Backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	fmt.Println("USPEX")
	return &Session{}, nil
}

// AnonymousLogin requires clients to authenticate using SMTP AUTH before sending emails
func (bkd *Backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	fmt.Println("USPEX Anonymous ", state.Hostname, state.RemoteAddr, state.LocalAddr)
	fmt.Println(state.TLS)

	return &Session{}, nil
}

// A Session is returned after successful login.
type Session struct{}

func (s *Session) Mail(from string, opts smtp.MailOptions) error {
	fmt.Println("HUI")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in mail", r)
		}
	}()
	fmt.Println("EMail from:", from, opts.Auth)
	go send.SendAnswerCouldNotFindUser(from)
	return nil
}

func (s *Session) Rcpt(to string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in rcpt", r)
		}
	}()
	fmt.Println("Rcpt to:", to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	grcpMailService, _ := grpc.Dial(
		"95.163.209.195:8083",
		grpc.WithInsecure(),
	)
	defer grcpMailService.Close()
	mailManager :=server.NewLetterServiceClient(grcpMailService)
	var mail string
	fmt.Println("||||||||||||||||||||||||||||||||||||||||||||||||||||||||\n\n\n")
	if b, err := ioutil.ReadAll(r); err != nil {
		return err
	} else {
		fmt.Println("Data:", string(b))
		mail+=string(b)
	}
	fmt.Println("\n\n\n||||||||||||||||||||||||||||||||||||||||||||||||||||||||")
	ctx:=context.Background()

	resp, _:=mailManager.SaveLetter(ctx, &server.Letter{})
	if resp.Ok==false{
		_ = send.SendAnswerCouldNotFindUser(getEmailFromMail(mail))
	}
	return nil
}

func getEmailFromMail(mail string) string{
	from:="\nFrom:"
	pos:=strings.Index(mail, from)
	var flag bool
	var email string
	for ;mail[pos]!='>';pos++{
		if flag ==true{
			email+=string(mail[pos])
		}
		if mail[pos]=='<'{
			flag=true
		}
	}
	return mail
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}
