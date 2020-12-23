package SendLetters

import (
	"github.com/stretchr/testify/assert"
	"smtpTest/proto/smtp"
	"testing"
)

func TestSendAnswer(t *testing.T) {
	email := "roofinda@gmail.com"
	sendAnswer(email)
}

func TestSendAnswer2(t *testing.T) {
	email := "roofinda@gmail.com"
	err := SendAnswerCouldNotFindUser(email)
	assert.NotNil(t, err)
	//assertPanic1(t, SendAnswerCouldNotFindUser, email)
}

func TestSendLetter(t *testing.T) {
	pbLetter := &smtp.Letter{
		Lid: 1,
		Sender: "Sender",
		Receiver: "Reciever",
		Theme: "Theme",
		Text: "Text",
		DateTime: 1,
	}

	err := SendLetter(pbLetter)
	assert.Nil(t, err)
	//assertPanic2(t, SendLetter, pbLetter)
}

// https://stackoverflow.com/questions/31595791/how-to-test-panics
//func assertPanic1(t *testing.T, f func(em string) error, email string) {
//	defer func() {
//		if r := recover(); r == nil {
//			t.Errorf("The code did not panic")
//		}
//	}()
//
//	_ = f(email)
//}
//
//func assertPanic2(t *testing.T, f func(letter *pb.Letter) error, letter *pb.Letter) {
//	defer func() {
//		if r := recover(); r == nil {
//			t.Errorf("The code did not panic")
//		}
//	}()
//
//	err := f(letter)
//	assert.Nil(t, err)
//}