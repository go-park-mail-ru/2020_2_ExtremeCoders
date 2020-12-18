package test

//type MyWriter struct {
//	Str []byte
//}
//
//type userKey struct {
//}
//
//func (writer MyWriter) Header() http.Header {
//	fmt.Println("implement me")
//	return nil
//}
//
//func (writer *MyWriter) Write(bytes []byte) (int, error) {
//	writer.Str = append(writer.Str, bytes...)
//	return len(writer.Str), nil
//}
//
//func (writer MyWriter) WriteHeader(statusCode int) {
//	fmt.Println("implement me")
//}

//func TestSendLetter(t *testing.T) {
//
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//	mockUseCase := mock.NewMockLetterUseCase(ctrl)
//	mockUseCase.EXPECT().SaveLetter(nil).MaxTimes(0)
//
//	uc := LetterDelivery.New(mockUseCase)
//	writer := MyWriter{}
//	r := http.Request{}
//	uc.SendLetter(&writer, &r)
//	if string(writer.Str) != string(errors.GetErrorNotPostAns()) {
//		t.Errorf("Expected error is  " + string(errors.GetErrorNotPostAns()))
//	}
//	writer.Str = []byte{}
//
//	r = http.Request{Method: "POST"}
//	uc.SendLetter(&writer, &r)
//	if string(writer.Str) != string(errors.GetErrorUnexpectedAns()) {
//		t.Errorf("Expected error is  " + string(errors.GetErrorUnexpectedAns()))
//	}
//
//	writer.Str = []byte{}
//	mockUseCase.EXPECT().GetReceivedLetters(nil).MaxTimes(0)
//	r = http.Request{Method: "POST"}
//	uc.GetRecvLetters(&writer, &r)
//	if string(writer.Str)!=string(errors.GetErrorUnexpectedAns()){
//		t.Errorf("Expected error is  " + string(errors.GetErrorUnexpectedAns()))
//	}
//
//	writer.Str = []byte{}
//	mockUseCase.EXPECT().GetSendedLetters(nil).MaxTimes(0)
//	r = http.Request{Method: "POST"}
//	uc.GetSendLetters(&writer, &r)
//	if string(writer.Str)!=string(errors.GetErrorUnexpectedAns()){
//		t.Errorf("Expected error is  " + string(errors.GetErrorUnexpectedAns()))
//	}
//}
