package SmtpManager

import (
	"github.com/stretchr/testify/assert"
	"smtpTest/proto/smtp"
	"testing"
)

func TestManager_GetLettersByDir(t *testing.T) {
	pbLetter := &smtp.Letter{
		Lid: 1,
		Sender: "Sender",
		Receiver: "Reciever",
		Theme: "Theme",
		Text: "Text",
		DateTime: 1,
	}

	manager := Manager{}

	assertPanic(t, manager, pbLetter)
}

// https://stackoverflow.com/questions/31595791/how-to-test-panics
func assertPanic(t *testing.T, manager Manager, letter *smtp.Letter) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	pbResponse := &smtp.Response{
		Ok: true,
		Description: "Description",
	}

	output := manager.GetLettersByDir(letter)

	assert.Equal(t, pbResponse, output)

}