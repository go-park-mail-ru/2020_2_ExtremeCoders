package test

import (
	"CleanArch/internal/Letter/LetterModel"
	"CleanArch/internal/Letter/LetterRepository"
	"CleanArch/internal/Letter/LetterRepository/LetterPostgres"
	"CleanArch/internal/User/UserModel"
	"CleanArch/internal/User/UserRepository"
	"CleanArch/internal/User/UserRepository/UserPostgres"
	"github.com/stretchr/testify/assert"
	"gitlab.com/slax0rr/go-pg-wrapper/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
	ormmocks "gitlab.com/slax0rr/go-pg-wrapper/mocks/orm"
)

var let=LetterModel.Letter{
	Id: 123,
	Sender: "dellvin.black@gmail.com",
	Receiver: "dellvin.black@gmail.com",
	Theme: "Meeting",
	Text: "Tomorrow at 6 am",
	DateTime: 78654678,
}

var user=UserModel.User{
	Id: 321,
	Name: "dellvin",
	Surname: "black",
	Email: "dellvin.black@gmail.com",
	Password: "1538",
}

var usertest=UserModel.User{
	Id: 0x141,
	Name: "",
	Surname: "",
	Email: "",
	Password: "",
}

var session=UserModel.Session{
	Id: "alsuiehfoqwuefhoq8723yroq7eroq73r",
	UserId: 321,
	User: nil,
}

func mockLetterDB() (*mocks.DB, LetterRepository.LetterDB) {
	db := new(mocks.DB)
	r := LetterPostgres.New(db)
	return db, r
}

func mockUserDB() (*mocks.DB, UserRepository.UserDB) {
	db := new(mocks.DB)
	r := UserPostgres.New(db)
	return db, r
}

func mockLetter(db *mocks.DB) *ormmocks.Query  {
	query := new(ormmocks.Query)
	mockCall := db.On("Model", mock.AnythingOfType("*models.letters")).Return(query)
	mockCall.RunFn = func(args mock.Arguments) {
		letter := args[0].(*LetterModel.Letter)
		*letter = let
	}
	return query
}

func mockUser(db *mocks.DB) *ormmocks.Query  {
	query := new(ormmocks.Query)
	mockCall := db.On("Model", mock.AnythingOfType("*models.users")).Return(query)
	mockCall.RunFn = func(args mock.Arguments) {
		User := args[0].(*UserModel.User)
		*User = user
	}
	return query
}

func mockSession(db *mocks.DB) *ormmocks.Query  {
	query := new(ormmocks.Query)
	mockCall := db.On("Model", mock.AnythingOfType("*models.sessions")).Return(query)
	mockCall.RunFn = func(args mock.Arguments) {
		letter := args[0].(*UserModel.Session)
		*letter = session
	}
	return query
}


func TestGetUserByID(t *testing.T) {
	db, r := mockUserDB()
	query := mockUser(db)

	query.On("Where", "id = ?", user.Id).Return(query)
	query.On("Select").Return(nil)

	answerCorrect, err := r.GetUserByUID(user.Id)
	assert.Nil(t, err)
	assert.Equal(t, user, *answerCorrect)
}

//func TestSaveLetter(t *testing.T){
//	db, r := mockDB()
//	query := mockGetRecvLetter(db)
//
//	query.On("Where", "receiver = ?", let.Receiver).Return(query)
//	query.On("Select").Return(nil)
//	foo, err := r.GetReceivedLetters(let.Receiver)
//	assert.Nil(t, err)
//	assert.Equal(t, let, foo)
//}

func TestGetRecvLetter(t *testing.T) {
	//db, r := mockDB()
	//query := mockGetRecvLetter(db)
	//
	//query.On("Where", "receiver = ?", let.Receiver).Return(query)
	//query.On("Select").Return(nil)
	//foo, err := r.GetReceivedLetters(let.Receiver)
	//assert.Nil(t, err)
	//assert.Equal(t, let, foo)
}
