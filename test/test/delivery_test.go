package test

import (
	"CleanArch/internal/Letter/LetterDelivery"
	"CleanArch/internal/errors"
	mock "CleanArch/test/mock_LetterUseCase"
	"fmt"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

type MyWriter struct {
	Str []byte
}

type userKey struct {
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
	mockUseCase := mock.NewMockLetterUseCase(ctrl)
	mockUseCase.EXPECT().SaveLetter(nil).MaxTimes(0)

	uc := LetterDelivery.New(mockUseCase)
	writer := MyWriter{}
	r := http.Request{}
	uc.SendLetter(&writer, &r)
	fmt.Println(string(writer.Str))
	if string(writer.Str) != string(errors.GetErrorNotPostAns()) {
		fmt.Println("FAIL")
		panic("Expected error is  " + string(errors.GetErrorNotPostAns()))
	}
	writer.Str = []byte{}

	r = http.Request{Method: "POST"}
	uc.SendLetter(&writer, &r)
	fmt.Println(string(writer.Str))
	if string(writer.Str) != string(errors.GetErrorUnexpectedAns()) {
		fmt.Println("FAIL")
		panic("Expected error is  " + string(errors.GetErrorUnexpectedAns()))
	}

	//writer.Str = []byte{}
	//Letter := &LetterModel.Letter{Receiver: "dellvin.black@gmail.com", Sender: "dellvin.black@gmail.com", Theme: "dellvin.black@gmail.com",
	//	Text: "dellvin.black@gmail.com", DateTime: int64(0)}
	//user := UserModel.User{12323,"ds","d","ss","aa","d"}
	//r1 := r.WithContext(context.WithValue(r.Context(),userKey{}, user))
	////r.Context().( )
	////r.Context.WithValue(ctx, userKey{}, user)
	////r.Context().WithValue(ctx, userKey{}, user)
	//fmt.Println(r.Context().Value(userKey{}))
	////
	//uc.SendLetter(&writer, r1)
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
