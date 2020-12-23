package SendLetters

import (
	pb "Mailer/SmtpService/proto/smtp"
	smtp2 "Mailer/SmtpService/proto/smtp"
	"context"
	"fmt"
	"github.com/emersion/go-smtp"
	"net"
	"strings"
	"time"
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

	servername := getHost(letter.Receiver) + ":25"
	to := []string{letter.Receiver}
	msg := strings.NewReader("To: " + letter.Receiver + "\r\n" +
		"From: " + letter.Sender + "\r\n" +
		letter.Theme + "\r\n" +
		"\r\n" +
		letter.Text + "\r\n")
	flag:=false
	for i:=0;i<100;i++{
		err := smtp.SendMail(servername, nil, letter.Sender, to, msg)
		if err != nil {
			fmt.Println("Repeat: ", err.Error())
			time.Sleep(1*time.Second)
		}else{
			flag=true
		}
	}
	if flag{
		fmt.Println("success sendLETTER2", servername)
	}else{
		fmt.Println("Could not send letter(")
	}
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
