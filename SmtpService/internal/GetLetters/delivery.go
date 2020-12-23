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
	if b, err := ioutil.ReadAll(r); err != nil {
		return err
	} else {
		mail+=string(b)
	}
	ctx:=context.Background()
	letter:=parseEmail(mail)
	fmt.Println("||||||||||||||||||||||||||||||||||||||||||||||||||||||||\n\n\n")
	fmt.Println(letter)
	fmt.Println("\n\n\n||||||||||||||||||||||||||||||||||||||||||||||||||||||||")
	resp, _:=mailManager.SaveLetter(ctx, &letter)
	if resp.Ok==false{
		_ = send.SendAnswerCouldNotFindUser(letter.Sender)
	}
	return nil
}

func parseEmail(s string) server.Letter{
	letter :=server.Letter{}
	from := "\nFrom:"
	subj := "\nSubject: "
	text := "\n\n"
	to := "\nTo: "
	fmt.Println(strings.Index(s, from))
	pos := strings.Index(s, from)
	var flag bool
	var email string
	var emTo string
	var emtext string
	var emSubj string
	for ; s[pos] != '>'; pos++ {
		if flag == true {
			email += string(s[pos])
		}
		if s[pos] == '<' {
			flag = true
		}
	}
	pos = strings.Index(s, subj)
	pos += len(subj)
	for ; s[pos] != '\n'; pos++ {
		emSubj += string(s[pos])
	}
	pos = strings.Index(s, to)
	pos += len(to)
	for ; s[pos] != '\n'; pos++ {
		emTo += string(s[pos])
	}

	pos = strings.Index(s, text)
	pos += len(text)
	for ; pos < len(s); pos++ {
		emtext += string(s[pos])
	}
	letter.Sender=emTo
	letter.Receiver=email
	letter.Theme=emSubj
	letter.Text=emtext
	return letter
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}
