package UserPostgres

import (
	"Mailer/UserService/internal/UserModel"
	"Mailer/UserService/internal/UserRepository"
	model "github.com/go-pg/pg/v9/orm"
	"github.com/stretchr/testify/assert"
	"gitlab.com/slax0rr/go-pg-wrapper/mocks"
	ormmocks "gitlab.com/slax0rr/go-pg-wrapper/mocks/orm"
	"testing"
)

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

var (
	testUser = UserModel.User{
		Id:      	1,
		Name: 		"UserName",
		Surname:  	"UserSurname",
		Email:      "UserEmail",
		Password:   "UserPassword",
		Img:     	"UserImg",
	}

	testSession = UserModel.Session{
		Id:        	"SessionId",
		UserId: 	1,
		User:		&testUser,
	}

	testFolder = UserModel.Folder{
		Uid:		1,
		Type: 		"FolderType",
		Name:		"FolderName",
	}
)

var repository *dataBase

var mockRes MockResult

// go test -coverprofile=coverage.out -coverpkg=./... -cover ./... && go tool cover -html=coverage.out
// go test `go list all | grep "UserService" | grep -v "mocks"`  -coverprofile=coverage.out.tmp -cover ./...

func NewRepo(t *testing.T) *dataBase {
	db := new(mocks.DB)
	repository = &dataBase{DB: db}
	return repository
}

func TestDataBase_AddSession(t *testing.T) {
	mockRes:= MockResult{}

	db := new(mocks.DB)
	query := new(ormmocks.Query)

	db.On("Model", &testSession).Return(query)
	query.On("Insert").Return(mockRes, nil)

	repository = &dataBase{DB: db}

	err := repository.AddSession(testSession.Id, uint64(testSession.UserId), testSession.User)
	assert.Nil(t, err)
}

func TestDataBase_AddUser(t *testing.T) {
	mockRes:= MockResult{}

	db := new(mocks.DB)
	query := new(ormmocks.Query)

	db.On("Model", &testUser).Return(query)
	query.On("Insert").Return(mockRes, nil)

	repository = &dataBase{DB: db}

	err := repository.AddUser(&testUser)
	assert.Nil(t, err)
}

func TestDataBase_CreateFolder(t *testing.T) {
	db := new(mocks.DB)
	whereQuery := new(ormmocks.Query)
	selectQuery := new(ormmocks.Query)

	db.On("Model", &testFolder).Return(whereQuery)
	whereQuery.On("Where", "type=? and name=? and uid=?",
		testFolder.Type, testFolder.Name, testFolder.Uid).Return(selectQuery)
	selectQuery.On("Select").Return(nil)

	repository = &dataBase{DB: db}

	err := repository.CreateFolder(testFolder.Name, testFolder.Type, testFolder.Uid)
	assert.NotNil(t, err)
}

func TestDataBase_GenerateSID(t *testing.T) {
	//db := new(mocks.DB)
	//whereQuery := new(ormmocks.Query)
	//selectQuery := new(ormmocks.Query)
	//
	//db.On("Model", &testSession).Return(whereQuery)
	//whereQuery.On("Where", "id=?", sid).Return(selectQuery)
	//
	//selectQuery.On("Select").Return(mockRes, nil)
	//
	//repository = &dataBase{DB: db}
	//
	//_, error := repository.GenerateSID()
	//assert.Nil(t, error)
}

func TestDataBase_GenerateUID(t *testing.T) {
	//db := new(mocks.DB)
	//whereQuery := new(ormmocks.Query)
	//selectQuery := new(ormmocks.Query)
	//
	//uid, _ := crypto.Int(crypto.Reader, big.NewInt(4294967295))
	//testUser.Id = uid.Uint64()
	//
	//db.On("Model", &testUser).Return(whereQuery)
	//whereQuery.On("Where", "id=?", uid.Int64()).Return(selectQuery)
	//
	//selectQuery.On("Select").Return(mockRes, nil)
	//
	//repository = &dataBase{DB: db}
	//
	//_, error := repository.GenerateUID()
	//assert.Nil(t, error)
}

func TestDataBase_GetFolderId(t *testing.T) {
	db := new(mocks.DB)
	whereQuery := new(ormmocks.Query)
	selectQuery := new(ormmocks.Query)

	db.On("Model", &testFolder).Return(whereQuery)
	whereQuery.On("Where", "type=? and uid=? and name=?",
		testFolder.Type, testFolder.Uid, testFolder.Name).Return(selectQuery)

	selectQuery.On("Select").Return(nil)

	repository = &dataBase{DB: db}

	_, error := repository.GetFolderId(testFolder.Uid, testFolder.Type, testFolder.Name)
	assert.Nil(t, error)
}

func TestDataBase_GetFoldersList(t *testing.T) {
	var folder_list []*UserModel.Folder

	db := new(mocks.DB)
	whereQuery := new(ormmocks.Query)
	selectQuery := new(ormmocks.Query)

	db.On("Model", &folder_list).Return(whereQuery)
	whereQuery.On("Where", "uid=? and type=?", testFolder.Uid, testFolder.Type).Return(selectQuery)

	selectQuery.On("Select").Return(nil)

	repository = &dataBase{DB: db}

	_, error := repository.GetFoldersList(testFolder.Uid, testFolder.Type)
	assert.Nil(t, error)
}

func TestDataBase_GetSessionByUID(t *testing.T) {
	session := UserModel.Session{UserId: int64(testUser.Id)}

	db := new(mocks.DB)
	whereQuery := new(ormmocks.Query)
	selectQuery := new(ormmocks.Query)

	db.On("Model", &session).Return(whereQuery)
	whereQuery.On("Where", "user_id=?", testUser.Id).Return(selectQuery)

	selectQuery.On("Select").Return(nil)

	repository = &dataBase{DB: db}

	_, error := repository.GetSessionByUID(testUser.Id)
	assert.Nil(t, error)
}

func TestDataBase_GetUserByEmail(t *testing.T) {
	user := UserModel.User{Email: testUser.Email}

	db := new(mocks.DB)
	whereQuery := new(ormmocks.Query)
	selectQuery := new(ormmocks.Query)

	db.On("Model", &user).Return(whereQuery)
	whereQuery.On("Where", "email=?", testUser.Email).Return(selectQuery)

	selectQuery.On("Select").Return(nil)

	repository = &dataBase{DB: db}

	_, error := repository.GetUserByEmail(testUser.Email)
	assert.Nil(t, error)
}

func TestDataBase_GetUserByUID(t *testing.T) {
	user := UserModel.User{}

	db := new(mocks.DB)
	whereQuery := new(ormmocks.Query)
	selectQuery := new(ormmocks.Query)

	db.On("Model", &user).Return(whereQuery)
	whereQuery.On("Where", "id=?", testUser.Id).Return(selectQuery)

	selectQuery.On("Select").Return(nil)

	repository = &dataBase{DB: db}

	_, error := repository.GetUserByUID(testUser.Id)
	assert.Nil(t, error)
}

func TestDataBase_IsEmailExists(t *testing.T) {
	user := UserModel.User{Email: testUser.Email}

	db := new(mocks.DB)
	whereQuery := new(ormmocks.Query)
	selectQuery := new(ormmocks.Query)

	db.On("Model", &user).Return(whereQuery)
	whereQuery.On("Where", "email=?", testUser.Email).Return(selectQuery)

	selectQuery.On("Select").Return(nil)

	repository = &dataBase{DB: db}

	error := repository.IsEmailExists(testUser.Email)
	assert.Exactly(t, UserRepository.EmailAlreadyExists, error)
}

func TestDataBase_IsEmailNotExists(t *testing.T) {
	user := UserModel.User{Email: testUser.Email}

	db := new(mocks.DB)
	whereQuery := new(ormmocks.Query)
	selectQuery := new(ormmocks.Query)

	db.On("Model", &user).Return(whereQuery)
	whereQuery.On("Where", "email=?", testUser.Email).Return(selectQuery)

	selectQuery.On("Select").Return(nil)

	repository = &dataBase{DB: db}

	error := repository.IsEmailExists(testUser.Email)
	assert.Exactly(t, UserRepository.EmailAlreadyExists, error)
}

func TestDataBase_IsOkSession(t *testing.T) {
	session := UserModel.Session{Id: testSession.Id}

	db := new(mocks.DB)
	whereQuery := new(ormmocks.Query)
	selectQuery := new(ormmocks.Query)

	db.On("Model", &session).Return(whereQuery)
	whereQuery.On("Where", "id=?", testSession.Id).Return(selectQuery)

	selectQuery.On("Select").Return(nil)

	repository = &dataBase{DB: db}

	_, error := repository.IsOkSession(testSession.Id)
	assert.Nil(t, error)
}

func TestDataBase_RemoveFolder(t *testing.T) {
	folder := UserModel.Folder{Id: testFolder.Id}

	db := new(mocks.DB)
	whereQuery := new(ormmocks.Query)
	deleteQuery := new(ormmocks.Query)

	db.On("Model", &folder).Return(whereQuery)
	whereQuery.On("Where", "id=?", testFolder.Id).Return(deleteQuery)

	deleteQuery.On("Delete").Return(mockRes, nil)

	repository = &dataBase{DB: db}

	error := repository.RemoveFolder(testFolder.Id)
	assert.Nil(t, error)
}

func TestDataBase_RemoveSession(t *testing.T) {
	session := UserModel.Session{Id: testSession.Id}

	db := new(mocks.DB)
	whereQuery := new(ormmocks.Query)
	selectDeleteQuery := new(ormmocks.Query)

	db.On("Model", &session).Return(whereQuery)
	whereQuery.On("Where", "id=?", testSession.Id).Return(selectDeleteQuery)

	selectDeleteQuery.On("Select").Return(nil)
	selectDeleteQuery.On("Delete").Return(mockRes, nil)

	repository = &dataBase{DB: db}

	error, _ := repository.RemoveSession(testSession.Id)
	assert.Nil(t, error)
}

func TestDataBase_RenameFolder(t *testing.T) {
	oldFolder := UserModel.Folder{Uid: testFolder.Uid, Name: testFolder.Name}

	db := new(mocks.DB)
	whereColumnQuery := new(ormmocks.Query)
	whereQuery := new(ormmocks.Query)
	selectUpdateQuery := new(ormmocks.Query)

	db.On("Model", &oldFolder).Return(whereColumnQuery)
	whereColumnQuery.On("Where", "uid=? and name=? and type=?",
		testFolder.Uid, testFolder.Name, testFolder.Type).Return(selectUpdateQuery)
	whereColumnQuery.On("Column", "name").Return(whereQuery)
	whereQuery.On("Where", "id=?", testFolder.Id).Return(selectUpdateQuery)

	selectUpdateQuery.On("Select").Return(nil)
	selectUpdateQuery.On("Update").Return(mockRes, nil)

	repository = &dataBase{DB: db}

	error := repository.RenameFolder(testFolder.Uid, testFolder.Type, testFolder.Name, testFolder.Name)
	assert.Nil(t, error)
}

func TestDataBase_UpdateProfile(t *testing.T) {
	oldUser := UserModel.User{Email: testUser.Email}

	db := new(mocks.DB)
	whereColumnQuery := new(ormmocks.Query)
	whereQuery := new(ormmocks.Query)
	selectUpdateQuery := new(ormmocks.Query)

	db.On("Model", &oldUser).Return(whereColumnQuery)
	whereColumnQuery.On("Where", "email=?", testUser.Email).Return(selectUpdateQuery)
	whereColumnQuery.On("Column", "name", "surname", "img").Return(whereQuery)
	whereQuery.On("Where", "email=?", testUser.Email).Return(selectUpdateQuery)

	selectUpdateQuery.On("Select").Return(nil)
	selectUpdateQuery.On("Update").Return(mockRes, nil)

	repository = &dataBase{DB: db}

	error := repository.UpdateProfile(oldUser, testUser.Email)
	assert.Nil(t, error)
}