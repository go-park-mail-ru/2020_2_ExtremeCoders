package FolderDelivery

import (
	"MainApplication/internal/pkg/context"
	letterService "MainApplication/proto/MailService"
	mailProto "MainApplication/proto/MailService"
	userProto "MainApplication/proto/UserServise"
	userService "MainApplication/proto/UserServise"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	folders, er := d.usClient.GetFoldersList(r.Context(), &userProto.Uid{Uid: user.Id})
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}
	fmt.Println(len(folders.Res))
	w.Write([]byte("HLLO"))
}

//get /user/foders/{recived/sended}/folderName - письма
func (d Delivery) GetLettersByFolder(w http.ResponseWriter, r *http.Request) {
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
	folderId, er := d.usClient.GetFolderId(r.Context(), &userProto.Folder{Uid: user.Id, Type: kind})
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
	w.Write([]byte("HLLO"))
}

//post /user/folders/{recived/sended}/folderName - добавить папку
func (d Delivery) AddFolder(w http.ResponseWriter, r *http.Request) {
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
	_, er = d.usClient.CreateFolder(r.Context(), &userProto.Folder{Uid: user.Id, Name: folderName})
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}
	fmt.Println("hello")
	w.Write([]byte("HLLO"))
}

//put /user/folders/{recived/sended}/folderName/letter body{letterID: id} - добавить письмо в папку
func (d Delivery) AddLetterInFolder(w http.ResponseWriter, r *http.Request) {
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
	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}
	folderId, er := d.usClient.GetFolderId(r.Context(), &userProto.Folder{Uid: user.Id, Type: kind})
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
			Type: false,
		})
	} else {
		resp, er = d.lsClient.AddLetterToDir(r.Context(), &mailProto.DirLid{
			Did:  folderId.Id,
			Lid:  lid,
			Type: true,
		})
	}
	if er != nil {
		fmt.Println("Er", er)
		w.Write(GetFoldersError(er))
		return
	}
	fmt.Println(resp)
	w.Write([]byte("HLLO"))
}


//put /user/folders/{recived/sended}/folderName body:{ name: newName} - переименовать папку
func (d Delivery) RenameFolder(w http.ResponseWriter, r *http.Request) {
	oldName := r.FormValue("oldName")
	newName := r.FormValue("newName")
	fmt.Println("url", r.URL, strings.Contains(r.URL.Path, "recived"), strings.Contains(r.URL.Path, "sended"), oldName, newName)
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
	_, er = d.usClient.RenameFolder(r.Context(), &userProto.RenameFolderMsg{Uid: user.Id, Type: kind, OldName: oldName, NewName: newName})
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}
	fmt.Println("OK")
	w.Write([]byte("HLLO"))
}

//delete /user/folders/{recived/sended}/folderName/letter body{letterID:Id} - удалить письмо из папки
func (d Delivery) RemoveLetterInFolder(w http.ResponseWriter, r *http.Request) {
	param := r.FormValue("letterId")
	lid, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.Write(GetFoldersError(err))
		return
	}
	fmt.Println("url", r.URL, strings.Contains(r.URL.Path, "recived"), strings.Contains(r.URL.Path, "sended"),lid)
	kind := true
	if strings.Contains(r.URL.Path, "recived") {
		kind = true
	} else {
		kind = false
	}

	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}

	folderId, er := d.usClient.GetFolderId(r.Context(), &userProto.Folder{Uid: user.Id, Type: kind})
	fmt.Println("FOLDER ID", folderId)
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}

	d.lsClient.RemoveLetterFromDir(r.Context(), &mailProto.DirLid{
		Did:  folderId.Id,
		Lid:  lid,
		Type: kind,
	})
	w.Write([]byte("HLLO"))
}

func (d Delivery) RemoveFolder(w http.ResponseWriter, r *http.Request) {
	folderName := r.FormValue("folderName")
	fmt.Println("url", r.URL, strings.Contains(r.URL.Path, "recived"), strings.Contains(r.URL.Path, "sended"),folderName)
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
	d.lsClient.RemoveDir(r.Context(), &mailProto.DirLid{
		Did:  folderId.Id,
		Type: kind,
	})
	w.Write([]byte("HLLO"))
}
