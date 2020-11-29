package FolderDelivery

import "net/http"

//get /user/folders/{recived/sended} - список папок
//get /user/foders/{recived/sended}/folderName - письма
//post /user/folders/{recived/sended}/folderName - добавить папку
//put /user/folders/{recived/sended}/folderName/letter body{letterID: id} - добавить письмо в папку
//put /user/folders/{recived/sended}/folderName body:{ name: newName} - переименовать папку
//delete /user/folders/{recived/sended}/folderName/letter body{letterID:Id} - удалить письмо из папки
//delete /user/folders/{recived/sended}/folderName  - удалить папку

type Interface interface {
	GetFolderList(w http.ResponseWriter, r *http.Request)
	GetLettersByFolder(w http.ResponseWriter, r *http.Request)
	AddFolder(w http.ResponseWriter, r *http.Request)
	AddLetterInFolder(w http.ResponseWriter, r *http.Request)
	RenameFolder(w http.ResponseWriter, r *http.Request)
	RemoveLetterInFolder(w http.ResponseWriter, r *http.Request)
	RemoveFolder(w http.ResponseWriter, r *http.Request)
}

func New() Interface {
	return Delivery{}
}
type Delivery struct {
}

func (d Delivery) GetFolderList(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (d Delivery) GetLettersByFolder(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (d Delivery) AddFolder(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (d Delivery) AddLetterInFolder(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (d Delivery) RenameFolder(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (d Delivery) RemoveLetterInFolder(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (d Delivery) RemoveFolder(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
