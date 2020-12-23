package SendLetters

import (
	pb "Mailer/SmtpService/proto/smtp"
	smtp2 "Mailer/SmtpService/proto/smtp"
	"context"
	"fmt"
	"github.com/emersion/go-smtp"
	"net"
	"strings"
)

type SMTPManager struct {
}

func NewSMTPManager()  smtp2.LetterServiceServer{
	return &SMTPManager{}
}

func (fm *SMTPManager) SendLetter(ctx context.Context, mail *pb.Letter) (*pb.Response, error) {
	err:=SendLetter(mail)
	resp:=pb.Response{Ok: true, Description: "ok"}
	if err!=nil{
		resp.Description=err.Error()
		resp.Ok=false
	}
	return &resp, nil
}

func SendAnswerCouldNotFindUser(email string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in getanswer2", r)
		}
	}()
	fmt.Println("KEK_1")
	if email == "bot@mailer.ru.com" {
		return nil
	}
	//auth := sasl.NewPlainClient("", "bot@mailer.ru.com", "password")
	fmt.Println("KEK_2")
	servername := getHost(email) + ":25"
	to := []string{email}
	msg := strings.NewReader("To: " + email + "\r\n" +
		"From: " + "bot@mailer.ru.com\r\n" +
		"(((((((\r\n" +
		"\r\n" +
		"Sorry but we could not find out friend(\r\n")
	fmt.Println("KEK_3")
	err := smtp.SendMail(servername, nil, "bot@mailer.ru.com", to, msg)
	fmt.Println("KEK_4")
	if err != nil {
		fmt.Println("Error in sendAnswer2", err.Error())
		return err
	}
	fmt.Println("success sendAnswer2", servername)
	return nil
}

func SendLetter(letter *smtp2.Letter) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in getanswer2", r)
		}
	}()
	fmt.Println("LOL_1")
	servername := getHost(letter.Receiver) + ":25"
	to := []string{letter.Receiver}
	msg := strings.NewReader("To: " + letter.Receiver + "\r\n" +
		"From: " + letter.Sender + "\r\n" +
		letter.Theme + "\r\n" +
		"\r\n" +
		letter.Text + "\r\n")
	fmt.Println("LOL_2")
	err := smtp.SendMail(servername, nil, letter.Sender, to, msg)
	fmt.Println("LOL3")
	if err != nil {
		fmt.Println("Error in sendLETTER2", err.Error())
		return err
	}
	fmt.Println("success sendLETTER2", servername)
	return nil
}

func getMailDomain(email string) string {
	flag := false
	var domail string
	for _, char := range email {
		if char == '@' {
			flag = true
			continue
		}
		if flag {
			domail += string(char)
		}
	}
	return domail
}

func getHost(email string) string {
	mxs, err := net.LookupMX(getMailDomain(email))
	if err != nil {
		panic(err)
	}
	if mxs[0].Host[len(mxs[0].Host)-1] == '.' {
		mxs[0].Host = mxs[0].Host[:len(mxs[0].Host)-1]
	}
	return mxs[0].Host
}
