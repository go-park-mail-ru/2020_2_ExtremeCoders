package test

import (
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestSendLetter(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "http://127.0.0.1/session",
		httpmock.NewStringResponder(200, `[{"code": 1, "description": "ok"}]`))
	httpmock.GetTotalCallCount()

	//ctrl := gomock.NewController(t)
	//defer ctrl.Finish()
	//Letter := &LetterModel.Letter{Receiver: "dellvin.black@gmail.com"}
	//mockLetter := mock.NewMockLetterUseCase(ctrl)
	//mockLetter.EXPECT().SaveLetter(&Letter).Return(nil)
	//uc := LetterDelivery.Delivery{Uc: mockLetter}
	//w :=http.ResponseWriter
	//r:= http.Request{}
	//uc.SendLetter(w, &r)
}
