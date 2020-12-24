package GetLetters

import (
	server "Mailer/MailService/proto"
	send "Mailer/SmtpService/internal/SendLetters"
	"context"
	"fmt"
	"github.com/emersion/go-smtp"
	"github.com/jhillyerd/enmime"
	"google.golang.org/grpc"
	"io"
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
	go func() {
		_ = send.SendAnswerCouldNotFindUser(from)
	}()
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
	ctx:=context.Background()

	env, _ := enmime.ReadEnvelope(r)
	// Headers can be retrieved via Envelope.GetHeader(name).
	fmt.Printf("From: %v\n", env.GetHeader("From"))
	// Address-type headers can be parsed into a list of decoded mail.Address structs.
	alist, _ := env.AddressList("To")
	for _, addr := range alist {
		fmt.Printf("To: %s <%s>\n", addr.Name, addr.Address)
	}
	fmt.Printf("Subject: %v\n", env.GetHeader("Subject"))


	// The plain text body is available as mime.Text.
	fmt.Printf("Text Body: %v chars\n", len(env.Text))

	// The HTML body is stored in mime.HTML.
	fmt.Printf("HTML Body: %v chars\n", len(env.HTML))

	// mime.Inlines is a slice of inlined attacments.
	fmt.Printf("Inlines: %v\n", len(env.Inlines))

	// mime.Attachments contains the non-inline attachments.
	fmt.Printf("Attachments: %v\n", len(env.Attachments))
	resp, _:=mailManager.SaveLetter(ctx, nil)
	if resp==nil || resp.Ok==false{
		fmt.Println("COULD NOT SAVE LETTER: ", resp.Description)
		_ = send.SendAnswerCouldNotFindUser(env.GetHeader("From"))
	}
	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}
