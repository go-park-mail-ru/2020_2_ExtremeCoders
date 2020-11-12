package test

import (
	//"CleanArch/internal/Letter/LetterDelivery"
	//"CleanArch/internal/Letter/LetterModel"
	//"github.com/golang/mock/gomock"
	//"net/http"
	//"strings"

	"fmt"
	"github.com/golang/mock/gomock"
	//"github.com/jarcoal/httpmock"
	//"github.com/tv42/mockhttp"
	"net/http"
	"testing"
)

type MyWriter struct {
	Str []byte
}

func (writer MyWriter) Header() http.Header {
	fmt.Println("implement me")
	return nil
}

func (writer *MyWriter) Write(bytes []byte) (int, error) {
	writer.Str = append(writer.Str, bytes...)
	return len(writer.Str), nil
}

func (writer MyWriter) WriteHeader(statusCode int) {
	fmt.Println("implement me")
}



func TestSendLetter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	//Letter := &LetterModel.Letter{Receiver: "dellvin.black@gmail.com"}
	////mockUseCase := mock.NewMockLetterUseCase(ctrl)
	////mockUseCase.EXPECT().SaveLetter(&Letter).Return(nil)
	//uc := LetterDelivery.New(mockUseCase)
	//
	//writer := MyWriter{}
	//r:= http.Request{}
	//uc.SendLetter(&writer, &r)
	//fmt.Println(writer.Str)
}
