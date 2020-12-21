package FolderDelivery

import (
	"Mailer/MainApplication/internal/pkg/context"
	letterService "Mailer/MainApplication/proto/MailService"
	mailProto "Mailer/MainApplication/proto/MailService"
	userProto "Mailer/MainApplication/proto/UserServise"
	userService "Mailer/MainApplication/proto/UserServise"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"github.com/gorilla/mux"
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

func New(usClient userService.UserServiceClient, lsClient letterService.LetterServiceClient) Interface {
	return Delivery{usClient: usClient, lsClient: lsClient}
}

type Delivery struct {
	usClient userService.UserServiceClient
	lsClient letterService.LetterServiceClient
}

//delete /user/folders/{recived/sended}/folderName  - удалить папку

//get /user/folders/{recived/sended} - список папок
func (d Delivery) GetFolderList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("url", r.URL, strings.Contains(r.URL.Path, "recived"), strings.Contains(r.URL.Path, "sended"))
	kind := "sended"
	if strings.Contains(r.URL.Path, "recived") {
		kind = "recived"
	}
	fmt.Println("KIND", kind)
	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}
	folders, er := d.usClient.GetFoldersList(r.Context(), &userProto.FolderUidType{Uid: user.Id, Type: kind})
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}
	fmt.Println(len(folders.Res))
	w.Write(ProtoFolderListResponse(folders.Res))
}

//get /user/foders/{recived/sended}/folderName - письма
func (d Delivery) GetLettersByFolder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	folderName := vars["folderName"]
	fmt.Println("url", r.URL, strings.Contains(r.URL.Path, "recived"), strings.Contains(r.URL.Path, "sended"))
	kind := "sended"
	if strings.Contains(r.URL.Path, "recived") {
		kind = "recived"
	}
	fmt.Println("KIND", kind)
	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		fmt.Println("Er", er)
		w.Write(GetFoldersError(er))
		return
	}
	folderId, er := d.usClient.GetFolderId(r.Context(), &userProto.Folder{Uid: user.Id, Name: folderName, Type: kind})
	if er != nil {
		fmt.Println("Er", er)
		w.Write(GetFoldersError(er))
		return
	}
	var letterList *mailProto.LetterListResponse
	if kind == "recived" {
		letterList, er = d.lsClient.GetLettersByDirRecv(r.Context(), &mailProto.DirName{DirName: folderId.Id})
	} else {
		letterList, er = d.lsClient.GetLettersByDirSend(r.Context(), &mailProto.DirName{DirName: folderId.Id})
	}
	if er != nil {
		fmt.Println("Er", er)
		w.Write(GetFoldersError(er))
		return
	}
	fmt.Println(len(letterList.Letter))
	w.Write(ProtoLetterListAnswer(letterList))
}

//post /user/folders/{recived/sended}/folderName - добавить папку
func (d Delivery) AddFolder(w http.ResponseWriter, r *http.Request) {
	if r.Method==http.MethodPut{
		d.RenameFolder(w, r)
	}
	folderName := r.FormValue("folderName")
	fmt.Println("url", r.URL, strings.Contains(r.URL.Path, "recived"), strings.Contains(r.URL.Path, "sended"),
		folderName)
	kind := "sended"
	if strings.Contains(r.URL.Path, "recived") {
		kind = "recived"
	}
	fmt.Println("KIND", kind)
	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}
	userProtoStruct:=&userProto.Folder{Uid: user.Id, Name: folderName, Type: kind}
	_, er = d.usClient.CreateFolder(r.Context(), userProtoStruct)
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}
	fmt.Println("hello")
	w.Write(SuccessRespAns())
}

//put /user/folders/{recived/sended}/folderName/letter body{letterID: id} - добавить письмо в папку
func (d Delivery) AddLetterInFolder(w http.ResponseWriter, r *http.Request) {
	if r.Method==http.MethodDelete{
		d.RemoveLetterInFolder(w, r)
	}
	param := r.FormValue("letterId")
	lid, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.Write(GetFoldersError(err))
		return
	}
	fmt.Println("url", r.URL, strings.Contains(r.URL.Path, "recived"), strings.Contains(r.URL.Path, "sended"), lid, err)
	kind := "sended"
	if strings.Contains(r.URL.Path, "recived") {
		kind = "recived"
	}
	fmt.Println("KIND", kind)
	folderName:=context.GetStrFormValueSafety(r, "folderName")
	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}

	folderId, er := d.usClient.GetFolderId(r.Context(), &userProto.Folder{Uid: user.Id, Type: kind, Name: folderName})
	fmt.Println("FOLDER ID", folderId)
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}

	var resp *mailProto.Response
	if kind == "recived" {
		resp, er = d.lsClient.AddLetterToDir(r.Context(), &mailProto.DirLid{
			Did:  folderId.Id,
			Lid:  lid,
			Type: true,
		})
	} else {
		resp, er = d.lsClient.AddLetterToDir(r.Context(), &mailProto.DirLid{
			Did:  folderId.Id,
			Lid:  lid,
			Type: false,
		})
	}
	if er != nil {
		fmt.Println("Er", er)
		w.Write(GetFoldersError(er))
		return
	}
	fmt.Println(resp)
	w.Write(ProtoResponseAnswer(resp))
}

//put /user/folders/{recived/sended}/folderName body:{ name: newName} - переименовать папку
func (d Delivery) RenameFolder(w http.ResponseWriter, r *http.Request) {
	fmt.Print("\n\n\nHUIn\n\n\n")
	oldName := r.FormValue("oldName")
	newName := r.FormValue("newName")
	fmt.Println("url", r.URL, strings.Contains(r.URL.Path, "recived"), strings.Contains(r.URL.Path, "sended"), oldName, newName)
	kind := "sended"
	if strings.Contains(r.URL.Path, "recived") {
		kind = "recived"
	}
	fmt.Println("KIND", kind)
	fmt.Print("\n\n\nHUIn\n\n\n")
	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}
	_, er = d.usClient.RenameFolder(r.Context(), &userProto.RenameFolderMsg{Uid: user.Id, Type: kind, OldName: oldName, NewName: newName})
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}
	fmt.Println("OK")
	w.Write(SuccessRespAns())
}

//delete /user/folders/{recived/sended}/folderName/letter body{letterID:Id} - удалить письмо из папки
func (d Delivery) RemoveLetterInFolder(w http.ResponseWriter, r *http.Request) {
	param := r.FormValue("letterId")
	lid, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.Write(GetFoldersError(err))
		return
	}
	fmt.Println("url", r.URL, strings.Contains(r.URL.Path, "recived"), strings.Contains(r.URL.Path, "sended"), lid)
	kind := true
	textKind := "recieved"
	if strings.Contains(r.URL.Path, "recived") {
		kind = true
	} else {
		textKind = "sended"
		kind = false
	}

	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}

	folderId, er := d.usClient.GetFolderId(r.Context(), &userProto.Folder{Uid: user.Id, Type: textKind})
	fmt.Println("FOLDER ID", folderId)
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}

	resp, err := d.lsClient.RemoveLetterFromDir(r.Context(), &mailProto.DirLid{
		Did:  folderId.Id,
		Lid:  lid,
		Type: kind,
	})

	w.Write(ProtoResponseAnswer(resp))
}

func (d Delivery) RemoveFolder(w http.ResponseWriter, r *http.Request) {
	folderName := r.FormValue("folderName")
	fmt.Println("url", r.URL, strings.Contains(r.URL.Path, "recived"), strings.Contains(r.URL.Path, "sended"), folderName)
	kind := true
	textKind := "recieved"
	if strings.Contains(r.URL.Path, "recived") {
		kind = true
	} else {
		textKind = "sended"
		kind = false
	}

	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}

	folderId, er := d.usClient.RemoveFolder(r.Context(), &userProto.Folder{Uid: user.Id, Type: textKind})
	fmt.Println("FOLDER ID", folderId)
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}
	resp, _ := d.lsClient.RemoveDir(r.Context(), &mailProto.DirLid{
		Did:  folderId.Id,
		Type: kind,
	})
	w.Write(ProtoResponseAnswer(resp))
}
