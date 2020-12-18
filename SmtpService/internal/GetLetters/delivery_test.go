package GetLetters

import (
	"github.com/emersion/go-smtp"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestBackend_AnonymousLogin(t *testing.T) {
	state := &smtp.ConnectionState{}

	session := &Session{}
	back := Backend{}

	output, err := back.AnonymousLogin(state)

	assert.Nil(t, err)
	assert.Equal(t, session, output)
}

func TestBackend_Login(t *testing.T) {
	state := &smtp.ConnectionState{}
	username := "username"
	password := "password"

	session := &Session{}
	back := Backend{}

	output, err := back.Login(state, username, password)

	assert.Nil(t, err)
	assert.Equal(t, session, output)
}

func TestSession_Data(t *testing.T) {
	session := &Session{}
	r := strings.NewReader("Hello, Reader!")
	err := session.Data(r)
	assert.Nil(t, err)
}

func TestSession_Logout(t *testing.T) {
	session := &Session{}
	err := session.Logout()
	assert.Nil(t, err)
}

func TestSession_Mail(t *testing.T) {
	session := &Session{}

	opts := smtp.MailOptions{}
	from := "roofinda@gmail.com"

	err := session.Mail(from, opts)
	assert.Nil(t, err)
}

func TestSession_Rcpt(t *testing.T) {
	session := &Session{}
	to := "roofinda@gmail.com"

	err := session.Rcpt(to)
	assert.Nil(t, err)
}

func TestSession_Reset(t *testing.T) {
	session := &Session{}
	err := session.Logout()
	assert.Nil(t, err)
}
