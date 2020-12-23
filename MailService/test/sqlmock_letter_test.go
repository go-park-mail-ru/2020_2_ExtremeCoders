package test

import (

	"MailService/internal/Model"
	"MailService/internal/Repository"
	"MailService/internal/Repository/LetterPostgres"
	"github.com/stretchr/testify/assert"
	"gitlab.com/slax0rr/go-pg-wrapper/mocks"
	"testing"

	model "github.com/go-pg/pg/v9/orm"
	"github.com/stretchr/testify/mock"
	ormmocks "gitlab.com/slax0rr/go-pg-wrapper/mocks/orm"
)


var let= Model.Letter{
	Id: 123,
	Sender: "dellvin.black@gmail.com",
	Receiver: "dellvin.black@gmail.com",
	Theme:    "Meeting",
	Text:     "Tomorrow at 6 am",
	DateTime: 78654678,
}

var email="akk@adf.ru"

var did uint64=101

var lets []Model.Letter


func mockLetterDB() (*mocks.DB, Repository.LetterDB) {
	db := new(mocks.DB)
	r := LetterPostgres.New(db)
	return db, r
}

func mockLetter(db *mocks.DB) *ormmocks.Query {
	query := new(ormmocks.Query)
	mockCall := db.On("Model", mock.AnythingOfType("*Model.Letter")).Return(query)
	mockCall.RunFn = func(args mock.Arguments) {
		letter := args[0].(*Model.Letter)
		*letter = let
	}
	return query
}

func mockLetters(db *mocks.DB) *ormmocks.Query  {
	lets =append(lets, let)
	query := new(ormmocks.Query)
	mockCall := db.On("Model", mock.AnythingOfType("*[]Model.Letter")).Return(query)
	mockCall.RunFn = func(args mock.Arguments) {
		letters := args[0].(*[]Model.Letter)
		*letters = lets
	}
	return query
}


func TestSaveMailRep(t *testing.T) {
	db, r := mockLetterDB()
	query := mockLetter(db)
	mockRes:= MockResult{}

	query.On("Insert").Return(mockRes, nil)
	err := r.SaveMail(let)
	assert.Nil(t, err)
}

func TestGetReceivedLettersRep(t *testing.T) {
	db, r := mockLetterDB()
	query := mockLetters(db)
	lets =append(lets, let)
	query.On("Where", "receiver=?", email).Return(query)
	query.On("Select").Return(nil)
	err, letters := r.GetLettersRecv(email)

	assert.Nil(t, err)
	assert.Equal(t, lets, letters)
}


func TestGetReceivedLettersDirRep(t *testing.T) {
	db, r := mockLetterDB()
	query := mockLetters(db)
	lets =append(lets, let)
	query.On("Where", "directory_recv=?", did).Return(query)
	query.On("Select").Return(nil)
	err, letters := r.GetLettersRecvDir(did)

	assert.Nil(t, err)
	assert.Equal(t, lets, letters)
}

func TestGetSentLettersDirRep(t *testing.T) {
	db, r := mockLetterDB()
	query := mockLetters(db)
	lets =append(lets, let)
	query.On("Where", "directory_send=?", did).Return(query)
	query.On("Select").Return(nil)
	err, letters := r.GetLettersSentDir(did)

	assert.Nil(t, err)
	assert.Equal(t, lets, letters)
}

func TestGetSendedLettersRep(t *testing.T) {
	db, r := mockLetterDB()
	query := mockLetters(db)
	lets =append(lets, let)
	query.On("Where", "sender=?", email).Return(query)
	query.On("Select").Return(nil)
	err, letters := r.GetLettersSent(email)
	assert.Nil(t, err)
	assert.Equal(t, lets, letters)
}

func TestGetLetterByLidRep(t *testing.T) {
	db, r := mockLetterDB()
	query := mockLetter(db)
	lets =append(lets, let)
	query.On("Where", "id=?", did).Return(query)
	query.On("Select").Return(nil)
	err, _ := r.GetLetterByLid(did)
	assert.Nil(t, err)
}

func TestGetLetterByFolderRep(t *testing.T) {
	db, r := mockLetterDB()
	query := mockLetters(db)
	lets =append(lets, let)
	query.On("Where", "directory_recv=? or directory_send=?", did, did).Return(query)
	query.On("Select").Return(nil)
	err, _ := r.GetLettersByFolder(did)
	assert.Nil(t, err)
}

func TestSetLetterWatchedRep(t *testing.T) {
	db, r := mockLetterDB()
	query := mockLetter(db)

	mockRes:= MockResult{}

	query.On("Where", "id=?", did).Return(query)
	query.On("Select").Return(nil)

	query.On("Column", "is_watched").Return(query)
	query.On("Where", "id=?", did)
	query.On("Update").Return(mockRes, nil)
	err, _ := r.SetLetterWatched(did)
	assert.Nil(t, err)
}

func TestAddLetterToDirRep(t *testing.T) {
	db, r := mockLetterDB()
	query := mockLetter(db)

	mockRes:= MockResult{}

	query.On("Where", "id=?", did).Return(query)
	query.On("Select").Return(nil)

	query.On("Column", "directory_recv").Return(query)
	query.On("Where", "id=?", did)
	query.On("Update").Return(mockRes, nil)
	err := r.AddLetterToDir(did, did, true)
	assert.Nil(t, err)
}

func TestAddLetterToDirFalseRep(t *testing.T) {
	db, r := mockLetterDB()
	query := mockLetter(db)

	mockRes:= MockResult{}

	query.On("Where", "id=?", did).Return(query)
	query.On("Select").Return(nil)

	query.On("Column", "directory_send").Return(query)
	query.On("Where", "id=?", did)
	query.On("Update").Return(mockRes, nil)
	err := r.AddLetterToDir(did, did, false)
	assert.Nil(t, err)
}

func TestRemoveLetterToDirFalseRep(t *testing.T) {
	db, r := mockLetterDB()
	query := mockLetter(db)

	mockRes:= MockResult{}

	query.On("Where", "id=?", did).Return(query)
	query.On("Select").Return(nil)

	query.On("Column", "directory_send").Return(query)
	query.On("Where", "id=?", did)
	query.On("Update").Return(mockRes, nil)
	err := r.RemoveLetterFromDir(did, did, false)
	assert.Nil(t, err)
}

func TestRemoveLetterToDirTrueRep(t *testing.T) {
	db, r := mockLetterDB()
	query := mockLetter(db)

	mockRes:= MockResult{}

	query.On("Where", "id=?", did).Return(query)
	query.On("Select").Return(nil)

	query.On("Column", "directory_recv").Return(query)
	query.On("Where", "id=?", did)
	query.On("Update").Return(mockRes, nil)
	err := r.RemoveLetterFromDir(did, did, true)
	assert.Nil(t, err)
}

type MockResult struct {

}

func (r MockResult) Model() model.Model {
	panic("implement me!")
}
func (r MockResult) RowsAffected() int {
	panic("implement me!")
}
func (r MockResult) RowsReturned() int {
	panic("implement me!")
}
