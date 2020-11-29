package FolderDelivery

import (
	"MainApplication/internal/pkg/context"
	letterService "MainApplication/proto/MailService"
	userService "MainApplication/proto/UserServise"
	"fmt"
	"net/http"
	"strings"
	mailProto "MainApplication/proto/MailService"
	userProto "MainApplication/proto/UserServise"


)

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
	usClient userService.UserServiceClient
	lsClient letterService.LetterServiceClient
}

//get /user/foders/{recived/sended}/folderName - письма
//post /user/folders/{recived/sended}/folderName - добавить папку
//put /user/folders/{recived/sended}/folderName/letter body{letterID: id} - добавить письмо в папку
//put /user/folders/{recived/sended}/folderName body:{ name: newName} - переименовать папку
//delete /user/folders/{recived/sended}/folderName/letter body{letterID:Id} - удалить письмо из папки
//delete /user/folders/{recived/sended}/folderName  - удалить папку

//get /user/folders/{recived/sended} - список папок
func (d Delivery) GetFolderList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("url",r.URL, strings.Contains(r.URL.Path, "recived"), strings.Contains(r.URL.Path, "sended"))
	kind := "sended"
	if strings.Contains(r.URL.Path, "recived"){
		kind = "recived"
	}
	er, user := context.GetUserFromCtx(r.Context())

	d.usClient.GetFoldersList(r.Context(), userProto.Uid{Uid:user.Id})
	w.Write([]byte("HLLO"))
}

func (d Delivery) GetLettersByFolder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("url",r.URL, strings.Contains(r.URL.Path, "recived"), strings.Contains(r.URL.Path, "sended"))
	w.Write([]byte("HLLO"))
}

func (d Delivery) AddFolder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("url",r.URL, strings.Contains(r.URL.Path, "recived"), strings.Contains(r.URL.Path, "sended"))
	w.Write([]byte("HLLO"))
}

func (d Delivery) AddLetterInFolder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("url",r.URL, strings.Contains(r.URL.Path, "recived"), strings.Contains(r.URL.Path, "sended"))
	panic("implement me")
}

func (d Delivery) RenameFolder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("url",r.URL, strings.Contains(r.URL.Path, "recived"), strings.Contains(r.URL.Path, "sended"))
	w.Write([]byte("HLLO"))
}

func (d Delivery) RemoveLetterInFolder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("url",r.URL, strings.Contains(r.URL.Path, "recived"), strings.Contains(r.URL.Path, "sended"))
	w.Write([]byte("HLLO"))
}

func (d Delivery) RemoveFolder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("url",r.URL, strings.Contains(r.URL.Path, "recived"), strings.Contains(r.URL.Path, "sended"))
	w.Write([]byte("HLLO"))
}
