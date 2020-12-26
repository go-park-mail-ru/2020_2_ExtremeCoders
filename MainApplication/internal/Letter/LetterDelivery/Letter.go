package LetterDelivery

import (
	"Mailer/MainApplication/internal/Letter/LetterModel"
	"Mailer/MainApplication/internal/Letter/LetterUseCase"
	"Mailer/MainApplication/internal/errors"
	"Mailer/MainApplication/internal/pkg/context"
	protoUs "Mailer/UserService/proto"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"time"
)
//go:generate mockgen -source=./Letter.go -destination=../../../test/mock_LetterDelivery/LetterDeliveryMock.go

type Interface interface {
	SendLetter(w http.ResponseWriter, r *http.Request)
	GetRecvLetters(w http.ResponseWriter, r *http.Request)
	GetSendLetters(w http.ResponseWriter, r *http.Request)
	WatchLetter(w http.ResponseWriter, r *http.Request)
	DeleteLetter(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
	GetLetterBy(w http.ResponseWriter, r *http.Request)
	SetLetterInSpam(w http.ResponseWriter, r *http.Request)
	SetLetterInBox(w http.ResponseWriter, r *http.Request)
}

type delivery struct {
	Uc LetterUseCase.LetterUseCase
	UserManager protoUs.UserServiceClient

}

func New(usecase LetterUseCase.LetterUseCase, userManager protoUs.UserServiceClient) Interface {
	return delivery{Uc: usecase, UserManager: userManager}
}

// Letter delete letter
// @Summary delete letter
// @Description delete letter {id:10}
// @ID delete-letter
// @Accept  json
// @Produce  json
// @Param id path int true "Letter ID"
// @Success 200
// @Router /letter [delete]
func (de delivery)DeleteLetter(w http.ResponseWriter, r *http.Request){
	id := context.GetStrFormValueSafety(r, "id")
	intID,err:=strconv.Atoi(id)
	if err!=nil{
		w.Write(GetDeleteLetterError(err))
		return
	}
	err=de.Uc.DeleteLetter(uint64(intID))
	w.Write(GetDeleteLetterError(err))
}

// Letter send letter
// @Summary send letter
// @Description send {to:'receiver', theme:'theme', text:'letter content'}
// @ID send-letter
// @Accept  json
// @Produce  json
// @Param letter body LetterModel.Letter true "Letter ID"
// @Success 200
// @Router /letter [post]
func (de delivery) SendLetter(w http.ResponseWriter, r *http.Request) {
	if r.Method==http.MethodDelete{
		de.DeleteLetter(w, r)
	}
	if r.Method != http.MethodPost {
		w.Write(errors.GetErrorNotPostAns())
		return
	}
	var letter LetterModel.Letter
	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetLettersError(er, nil))
		return
	}
	letter.Sender = user.Email
	letter.Receiver = context.GetStrFormValueSafety(r, "to")
	letter.Theme = context.GetStrFormValueSafety(r, "theme")
	letter.Text = context.GetStrFormValueSafety(r, "text")
	letter.DateTime = time.Now().Unix()
	err := de.Uc.SaveLetter(&letter)
	w.Write(SendLetterError(err, letter))


	receiver, err := de.UserManager.GetUserByEmail(r.Context(), &protoUs.Email{Email: letter.Receiver})
	if err != nil {
		fmt.Printf("\n\nNOTIFY TO %s FAILED %s\n", letter.Receiver, err.Error())
		return
	}

	err, letters := de.Uc.GetReceivedLetters(receiver.Email, uint64(5), uint64(0))
	if err != nil {
		fmt.Printf("\n\nNOTIFY TO %s FAILED %s\n", letter.Receiver, err.Error())
		return
	}

	if val, ok := context.WebSockets[receiver.Uid]; ok {
		err = val.WriteMessage(websocket.TextMessage,GetLettersError(err, letters))
		if err!=nil {
			fmt.Printf("\n\nNOTIFY TO %s FAILED %s\n", letter.Receiver, err.Error())
		}
	} else {
		fmt.Printf("\n\nNOTIFY TO %s FAILED\n", letter.Receiver)
	}
}

// Letter get received letter
// @Summary get received letter
// @Description get user/letter/sent/{limit}/{offset} - получить полученные письма
// @ID get-received-letter
// @Accept  json
// @Produce  json
// @Param limit path int true "limit"
// @Param offset path int true "offset"
// @Success 200 {array} LetterModel.Letter
// @Router /user/letter/received/{limit}/{offset} [get]
func (de delivery) GetRecvLetters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	limit := vars["limit"]
	offset:=vars["offset"]
	intLim, err:=strconv.Atoi(limit)
	if err != nil {
		intLim = 5
	}
	intOff, err:=strconv.Atoi(offset)
	if err != nil {
		intOff = 0
	}
	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetLettersError(er, nil))
		return
	}
	err, letters := de.Uc.GetReceivedLetters(user.Email, uint64(intLim), uint64(intOff))
	w.Write(GetLettersError(err, letters))
}

// Letter get sended letter
// @Summary get sended letter
// @Description get user/letter/sent/{limit}/{offset} - получить отправленные письма
// @ID get-send-letter
// @Accept  json
// @Produce  json
// @Param limit path int true "limit"
// @Param offset path int true "offset"
// @Success 200 {array} LetterModel.Letter
// @Router /user/letter/sent/{limit}/{offset} [get]
func (de delivery) GetSendLetters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	limit := vars["limit"]
	offset:=vars["offset"]
	intLim, err:=strconv.Atoi(limit)
	if err != nil {
		intLim = 5
	}
	intOff, err:=strconv.Atoi(offset)
	if err != nil {
		intOff = 0
	}

	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetLettersError(er, nil))
		return
	}
	err, letters := de.Uc.GetSendedLetters(user.Email, uint64(intLim), uint64(intOff))
	w.Write(GetLettersError(err, letters))
}

// Letter Watch
// @Summary set watch togle
// @Description отметить письмо как прочитанное/непрочитанное /watch/letter {id:'id'}
// @ID watch-letter
// @Accept  json
// @Produce  json
// @Param id path int true "letter id"
// @Success 200
// @Router /watch/letter [put]
func (de delivery) WatchLetter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Write(errors.GetErrorNotPostAns())
		return
	}
	id := context.GetStrFormValueSafety(r, "id")
	num, _:=strconv.Atoi(id)
	_, _ = de.Uc.WatchLetter(uint64(num))
}


// Search search
// @Summary Search in letter
// @Description get letter/{similar} - поиск по всем письмам
// @ID search-search
// @Accept  json
// @Produce  json
// @Param similar path string true "search template"
// @Success 200 string Res
// @Router /letter/{similar} [get]
func (de delivery) Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Write(errors.GetErrorUnexpectedAns())
		return
	}
	vars := mux.Vars(r)
	sim := vars["similar"]
	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetLettersError(er, nil))
		return
	}

	searchRes:=de.Uc.FindSim(sim, user.Email)
	w.Write([]byte(searchRes))
}

// Search Search
// @Summary Search in letter
// @Description поиск по всем файлам
// @Description get letter/by/{what}/{value} - what может быть равен
// @Description (id, sender, receiver, theme, text, date_time, directory_recv, directory_send)
// @ID all-search
// @Accept  json
// @Produce  json
// @Param what path string true "search type"
// @Param value path string true "search template"
// @Success 200 {array} LetterModel.Letter
// @Router /letter/by/{what}/{value} [get]
func (de delivery) GetLetterBy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Write(errors.GetErrorUnexpectedAns())
		return
	}
	vars := mux.Vars(r)
	what := vars["what"]
	val:=vars["value"]
	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetLettersError(er, nil))
		return
	}
	err, letters:=de.Uc.GetLetterBy(what, val, user.Email)
	w.Write(GetLettersError(err, letters))
}

func (de delivery) SetLetterInSpam(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Write(errors.GetLetterSpamError())
		return
	}
	lidstr:=context.GetStrFormValueSafety(r, "lid")
	lid,err:=strconv.Atoi(lidstr)
	if err!=nil{
		w.Write(errors.GetErrorUnexpectedAns())
		return
	}
	err=de.Uc.SetLetterInSpam(uint64(lid))
	w.Write(GetLetterSpamError(err))
}

func (de delivery) SetLetterInBox(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Write(errors.GetErrorUnexpectedAns())
		return
	}
	lidstr:=context.GetStrFormValueSafety(r, "lid")
	lid,err:=strconv.Atoi(lidstr)
	if err!=nil{
		w.Write(errors.GetErrorUnexpectedAns())
		return
	}
	err=de.Uc.SetLetterInBox(uint64(lid))
	w.Write(GetLetterBoxError(err))
}