package SendLetters
import (
	"crypto/tls"
	"fmt"
	"github.com/emersion/go-smtp"
	"log"
	"net"
	"net/mail"
	baseSMTP "net/smtp"
	pb "smtpTest/proto"
	"strings"
)

func sendAnswer(email string){
	if email=="bot@mailer.ru.com"{
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in getanswer", r)
		}
	}()
	fmt.Println("HUI_1")
	from := mail.Address{"", "bot@mailer.ru.com"}
	to   := mail.Address{"", email}
	subj := "Hello"
	body := "We are happy to see you in our alfa smtp-test!"

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj
	fmt.Println("HUI_2")
	// Setup message
	message := ""
	for k,v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := getHost(email)+":25"

	host, _, _ := net.SplitHostPort(servername)
	fmt.Println("HUI_3")
	auth := baseSMTP.PlainAuth("",from.String(), "keklol123", host)
	fmt.Println("HUI_4")
	// TLS config
	tlsconfig := &tls.Config {
		InsecureSkipVerify: true,
		ServerName: host,
	}

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("HUI_5")
	c, err := baseSMTP.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("HUI_6")
	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}
	fmt.Println("HUI_7")
	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}
	fmt.Println("HUI_8")
	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}
	fmt.Println("HUI_9")
	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("HUI_10")
	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("HUI_11")
	err = w.Close()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("HUI_12")
	c.Quit()
	fmt.Println("Sent answer to: ", email)
}


func SendAnswer2(email string)error{
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in getanswer2", r)
		}
	}()
	fmt.Println("KEK_1")
	if email=="bot@mailer.ru.com"{
		return nil
	}
	//auth := sasl.NewPlainClient("", "bot@mailer.ru.com", "password")
	fmt.Println("KEK_2")
	servername := getHost(email)+":25"
	to := []string{email}
	msg := strings.NewReader("To: "+email+"\r\n" +
		"From: "+"bot@mailer.ru.com\r\n" +
		"Subject: Hello SMTP!!!\r\n" +
		"\r\n" +
		"We are happy to see you in our alfa smtp-test!\r\n")
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

func SendLetter(letter *pb.Letter)error{
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in getanswer2", r)
		}
	}()
	fmt.Println("LOL_1")
	servername := getHost(letter.Receiver)+":25"
	to := []string{letter.Receiver}
	msg := strings.NewReader("To: "+letter.Receiver+"\r\n" +
		"From: "+letter.Sender+"\r\n" +
		letter.Theme+"\r\n" +
		"\r\n" +
		letter.Text+"\r\n")
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

func getMailDomain(email string) string{
	flag:=false
	var domail string
	for _, char:= range email{
		if char=='@'{
			flag=true
			continue
		}
		if flag{
			domail+=string(char)
		}
	}
	return domail
}

func getHost (email string) string {
	mxs, err := net.LookupMX(getMailDomain(email))
	if err != nil {
		panic(err)
	}
	if mxs[0].Host[len(mxs[0].Host)-1]=='.'{
		mxs[0].Host=mxs[0].Host[:len(mxs[0].Host)-1]
	}
	return mxs[0].Host
}
