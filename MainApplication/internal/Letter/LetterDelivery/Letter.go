package LetterDelivery

import (
	"Mailer/MainApplication/internal/Letter/LetterModel"
	"Mailer/MainApplication/internal/Letter/LetterUseCase"
	"Mailer/MainApplication/internal/errors"
	"Mailer/MainApplication/internal/pkg/context"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type Interface interface {
	SendLetter(w http.ResponseWriter, r *http.Request)
	GetRecvLetters(w http.ResponseWriter, r *http.Request)
	GetSendLetters(w http.ResponseWriter, r *http.Request)
	WatchLetter(w http.ResponseWriter, r *http.Request)
	DeleteLetter(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
	GetLetterBy(w http.ResponseWriter, r *http.Request)
}

type delivery struct {
	Uc LetterUseCase.LetterUseCase
}

func New(usecase LetterUseCase.LetterUseCase) Interface {
	return delivery{Uc: usecase}
}

func (de delivery) DeleteLetter(w http.ResponseWriter, r *http.Request) {
	id := context.GetStrFormValueSafety(r, "id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		_, _ = w.Write(GetDeleteLetterError(err))
		return
	}
	err = de.Uc.DeleteLetter(uint64(intID))
	_, _ = w.Write(GetDeleteLetterError(err))
}

func (de delivery) SendLetter(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		de.DeleteLetter(w, r)
	}
	if r.Method != http.MethodPost {
		_, _ = w.Write(errors.GetErrorNotPostAns())
		return
	}
	var letter LetterModel.Letter
	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		_, _ = w.Write(GetLettersError(er, nil))
		return
	}
	letter.Sender = user.Email
	letter.Receiver = context.GetStrFormValueSafety(r, "to")
	letter.Theme = context.GetStrFormValueSafety(r, "theme")
	letter.Text = context.GetStrFormValueSafety(r, "text")
	letter.DateTime = time.Now().Unix()
	err := de.Uc.SaveLetter(&letter)
	_, _ = w.Write(SendLetterError(err, letter))
}

func (de delivery) GetRecvLetters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	limit := vars["limit"]
	offset := vars["offset"]
	intLim, _ := strconv.Atoi(limit)
	intOff, err := strconv.Atoi(offset)

	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		_, _ = w.Write(GetLettersError(er, nil))
		return
	}
	err, letters := de.Uc.GetReceivedLetters(user.Email, uint64(intLim), uint64(intOff))
	_, _ = w.Write(GetLettersError(err, letters))
}

func (de delivery) GetSendLetters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	limit := vars["limit"]
	offset:=vars["offset"]
	intLim, err:=strconv.Atoi(limit)
	intOff, err:=strconv.Atoi(offset)

	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		_, _ = w.Write(GetLettersError(er, nil))
		return
	}

	err, letters := de.Uc.GetSendedLetters(user.Email, uint64(intLim), uint64(intOff))
	w.Write(GetLettersError(err, letters))
}

func (de delivery) WatchLetter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		_, _ = w.Write(errors.GetErrorNotPostAns())
		return
	}
	id := context.GetStrFormValueSafety(r, "id")
	num, _ := strconv.Atoi(id)
	_, _ = de.Uc.WatchLetter(uint64(num))
}

func (de delivery) Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		_, _ = w.Write(errors.GetErrorUnexpectedAns())
		return
	}
	vars := mux.Vars(r)
	sim := vars["similar"]
	searchRes := de.Uc.FindSim(sim)
	_, _ = w.Write([]byte(searchRes))
}

func (de delivery) GetLetterBy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		_, _ = w.Write(errors.GetErrorUnexpectedAns())
		return
	}
	vars := mux.Vars(r)
	what := vars["what"]
	val := vars["value"]
	err, letters := de.Uc.GetLetterBy(what, val)
	_, _ = w.Write(GetLettersError(err, letters))
}
