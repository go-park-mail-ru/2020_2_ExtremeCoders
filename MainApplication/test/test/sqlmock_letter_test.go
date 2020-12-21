package test
//
//import (
//	"Mailer/MainApplication/internal/Letter/LetterModel"
//	"Mailer/MainApplication/internal/Letter/LetterRepository"
//	"Mailer/MainApplication/internal/Letter/LetterRepository/LetterPostgres"
//	"Mailer/MainApplication/internal/User/UserModel"
//	"Mailer/MainApplication/internal/User/UserRepository"
//	"Mailer/MainApplication/internal/User/UserRepository/UserPostgres"
//	"github.com/stretchr/testify/assert"
//	"gitlab.com/slax0rr/go-pg-wrapper/mocks"
//	"testing"
//
//	model "github.com/go-pg/pg/v9/orm"
//	"github.com/stretchr/testify/mock"
//	ormmocks "gitlab.com/slax0rr/go-pg-wrapper/mocks/orm"
//)
//
//var let=LetterModel.Letter{
//	Id: 123,
//	Sender: "dellvin.black@gmail.com",
//	Receiver: "dellvin.black@gmail.com",
//	Theme: "Meeting",
//	Text: "Tomorrow at 6 am",
//	DateTime: 78654678,
//}
//
//var lets []LetterModel.Letter
//
//
//var user=UserModel.User{
//	Id: 321,
//	Name: "dellvin",
//	Surname: "black",
//	Email: "dellvin.black@gmail.com",
//	Password: "1538",
//}
//
//var usertest=UserModel.User{
//	Id: 321,
//	Name: "",
//	Surname: "",
//	Email: "",
//	Password: "",
//}
//
//var session=UserModel.Session{
//	Id: "alsuiehfoqwuefhoq8723yroq7eroq73r",
//	UserId: 321,
//	User: nil,
//}
//
//func mockLetterDB() (*mocks.DB, LetterRepository.LetterDB) {
//	db := new(mocks.DB)
//	r := LetterPostgres.New(db)
//	return db, r
//}
//
//func mockUserDB() (*mocks.DB, UserRepository.UserDB) {
//	db := new(mocks.DB)
//	r := UserPostgres.New(db)
//	return db, r
//}
//
//func mockLetter(db *mocks.DB) *ormmocks.Query  {
//	query := new(ormmocks.Query)
//	mockCall := db.On("Model", mock.AnythingOfType("*LetterModel.Letter")).Return(query)
//	mockCall.RunFn = func(args mock.Arguments) {
//		letter := args[0].(*LetterModel.Letter)
//		*letter = let
//	}
//	return query
//}
//
//func mockLetters(db *mocks.DB) *ormmocks.Query  {
//	lets =append(lets, let)
//	query := new(ormmocks.Query)
//	mockCall := db.On("Model", mock.AnythingOfType("*[]LetterModel.Letter")).Return(query)
//	mockCall.RunFn = func(args mock.Arguments) {
//		letters := args[0].(*[]LetterModel.Letter)
//		*letters = lets
//	}
//	return query
//}
//
//
//func mockSession(db *mocks.DB) *ormmocks.Query  {
//	query := new(ormmocks.Query)
//	mockCall := db.On("Model", mock.AnythingOfType("*UserModel.Session")).Return(query)
//	mockCall.RunFn = func(args mock.Arguments) {
//		Session := args[0].(*UserModel.Session)
//		*Session = session
//	}
//	return query
//}
//
//func mockUser(db *mocks.DB) *ormmocks.Query  {
//	query := new(ormmocks.Query)
//	mockCall := db.On("Model", mock.AnythingOfType("*UserModel.User")).Return(query)
//	mockCall.RunFn = func(args mock.Arguments) {
//		User := args[0].(*UserModel.User)
//		*User = usertest
//	}
//	return query
//}
//
//func TestGetUserByID(t *testing.T) {
//	db, r := mockUserDB()
//	query := mockUser(db)
//
//	query.On("Where", "id=?", usertest.Id).Return(query)
//	query.On("Select").Return(nil)
//
//	answerCorrect, err := r.GetUserByUID(usertest.Id)
//	assert.Nil(t, err)
//	assert.Equal(t, usertest, *answerCorrect)
//}
//
//func TestIsOkSession(t *testing.T) {
//	db, r := mockUserDB()
//	query := mockSession(db)
//
//	query.On("Where", "id=?", session.Id).Return(query)
//	query.On("Select").Return(nil)
//
//	answerCorrect, err := r.IsOkSession(session.Id)
//	assert.Nil(t, err)
//	assert.Equal(t, session.UserId, int64(answerCorrect))
//}
//
//func TestGetSessionByUID(t *testing.T) {
//	db, r := mockUserDB()
//	query := mockSession(db)
//	query.On("Where", "user_id=?", uint64(session.UserId)).Return(query)
//	query.On("Select").Return(nil)
//	answerCorrect, err := r.GetSessionByUID(uint64(session.UserId))
//	assert.Nil(t, err)
//	assert.Equal(t, session.Id, answerCorrect)
//}
//
//func TestIsEmailExists(t *testing.T) {
//	db, r := mockUserDB()
//	query := mockUser(db)
//	query.On("Where", "email=?", user.Email).Return(query)
//	query.On("Select").Return(nil)
//	err := r.IsEmailExists(user.Email)
//	assert.Equal(t, err, UserRepository.EmailAlreadyExists)
//}
//
//func TestAddUser(t *testing.T) {
//	db, r := mockUserDB()
//	query := mockUser(db)
//	mockRes:= MockResult{}
//	query.On("Insert").Return(mockRes, nil)
//	err := r.AddUser(&usertest)
//	assert.Nil(t, err)
//}
//
//func TestAddSession(t *testing.T) {
//	db, r := mockUserDB()
//	query := mockSession(db)
//	mockRes:= MockResult{}
//	query.On("Insert").Return(mockRes, nil)
//	err := r.AddSession(session.Id, usertest.Id, &usertest)
//	assert.Nil(t, err)
//}
//
//func TestGetUserByEmail(t *testing.T) {
//	db, r := mockUserDB()
//	query := mockUser(db)
//	query.On("Where", "email=?", usertest.Email).Return(query)
//	query.On("Select").Return(nil)
//	answerCorrect, err := r.GetUserByEmail(usertest.Email)
//	assert.Nil(t, err)
//	assert.Equal(t, usertest, *answerCorrect)
//}
//
//func TestUpdateProfile(t *testing.T) {
//	db, r := mockUserDB()
//	query := mockUser(db)
//	mockRes:= MockResult{}
//
//	query.On("Where", "email=?", usertest.Email).Return(query)
//	query.On("Select").Return(nil)
//
//	query.On("Column","name", "surname", "img").Return(query)
//	query.On("Where", "email=?", usertest.Email)
//	query.On("Update").Return(mockRes, nil)
//	err := r.UpdateProfile(usertest, usertest.Email)
//	assert.Nil(t, err)
//}
//
//func TestRemoveSession(t *testing.T) {
//	db, r := mockUserDB()
//	query := mockSession(db)
//	mockRes:= MockResult{}
//	query.On("Where", "id=?", session.Id).Return(query)
//	query.On("Select").Return(nil)
//	query.On("Where", "id=?", session.Id)
//	query.On("Delete").Return(mockRes, nil)
//	err,res := r.RemoveSession(session.Id)
//	assert.Nil(t, err)
//	assert.Equal(t, session.UserId, int64(res))
//}
//
////TODO ==================================================
//
//func TestSaveMailRep(t *testing.T) {
//	db, r := mockLetterDB()
//	query := mockLetter(db)
//	mockRes:= MockResult{}
//
//	query.On("Insert").Return(mockRes, nil)
//	err := r.SaveMail(let)
//	assert.Nil(t, err)
//}
//
//func TestIsUserExistRep(t *testing.T) {
//	db, r := mockLetterDB()
//	query := mockUser(db)
//
//	query.On("Where", "email=?", usertest.Email).Return(query)
//	query.On("Select").Return(nil)
//	err := r.IsUserExist(usertest.Email)
//	assert.Nil(t, err)
//}
//
//func TestGetReceivedLettersRep(t *testing.T) {
//	db, r := mockLetterDB()
//	query := mockLetters(db)
//
//	lets =append(lets, let)
//	query.On("Where", "receiver=?", usertest.Email).Return(query)
//	query.On("Select").Return(nil)
//	err, letters := r.GetReceivedLetters(usertest.Email)
//	assert.Nil(t, err)
//	assert.Equal(t, lets, letters)
//}
//
//func TestGetSendedLettersRep(t *testing.T) {
//	db, r := mockLetterDB()
//	query := mockLetters(db)
//
//	lets =append(lets, let)
//	query.On("Where", "sender=?", usertest.Email).Return(query)
//	query.On("Select").Return(nil)
//	err, letters := r.GetSendedLetters(usertest.Email)
//	assert.Nil(t, err)
//	assert.Equal(t, lets, letters)
//}
//type MockResult struct {
//}
//
//func (r MockResult) Model() model.Model {
//	panic("implement me!")
//}
//func (r MockResult) RowsAffected() int {
//	panic("implement me!")
//}
//func (r MockResult) RowsReturned() int {
//	panic("implement me!")
//}