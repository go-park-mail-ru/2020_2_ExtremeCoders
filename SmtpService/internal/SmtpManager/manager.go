package SmtpManager

import (
	"Mailer/SmtpService/internal/SendLetters"
	"Mailer/SmtpService/proto/smtp"
)

type Manager struct {
}

func (m Manager) GetLettersByDir(letter *smtp.Letter) *smtp.Response {
	err := SendLetters.SendLetter(letter)
	resp := smtp.Response{Ok: true, Description: err.Error()}
	if err != nil {
		resp.Ok = false
	}
	return &resp
}
