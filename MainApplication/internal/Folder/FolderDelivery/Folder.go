package FolderDelivery

import (
	letterService "Mailer/MailService/proto"
	"Mailer/MainApplication/internal/pkg/context"
	userService "Mailer/UserService/proto"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)


//go:generate mockgen -source=./Folder.go -destination=../../../test/mock_FolderDelivery/FolderDeliveyeMock.go

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

// Folder GetFolderList godoc
// @Summary getFolderList
// @Description user/folders/{recived/sended} - список папок в отправленных (полученных) письмах
// @ID GetFolderList
// @Accept  json
// @Produce  json
// @Success 200 {object} FolderList fl
// @Router /user/folders/{recived/sended} [get]
func (d Delivery) GetFolderList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("url", r.URL, strings.Contains(r.URL.Path, "received"), strings.Contains(r.URL.Path, "sended"))
	kind := "sended"
	if strings.Contains(r.URL.Path, "received") {
		kind = "received"
	}
	fmt.Println("KIND", kind)
	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}
	folders, er := d.usClient.GetFoldersList(r.Context(), &userService.FolderUidType{Uid: user.Id, Type: kind})
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}
	fmt.Println(len(folders.Res))
	w.Write(ProtoFolderListResponse(folders.Res))
}

// Folder GetLettersByFolder godoc
// @Summary Get letters by folder
// @Description письма из папки в полученых (отправленных) user/foders/{recived/sended}/{folderName}/{limit}/{offset}
// @ID GetLettersByFolder
// @Accept  json
// @Produce  json
// @Param limit path int true "limit"
// @Param offset path int true "offset"
// @Success 200 {object} LetterList ll
// @Router /user/folders/{recived/sended}/{folderName}/{limit}/{offset} [get]
func (d Delivery) GetLettersByFolder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	folderName := vars["folderName"]
	fmt.Println("url", r.URL, strings.Contains(r.URL.Path, "received"), strings.Contains(r.URL.Path, "sended"))
	kind := "sended"
	if strings.Contains(r.URL.Path, "received") {
		kind = "received"
	}
	fmt.Println("KIND", kind)
	folderId, err:=strconv.Atoi(folderName)
	if err!=nil{
		w.Write(GetFoldersError(err))
		return
	}
	var letterList *letterService.LetterListResponse
	if kind == "received" {
		letterList, err = d.lsClient.GetLettersByDirRecv(r.Context(), &letterService.DirName{DirName: uint64(folderId)})
	} else {
		letterList, err = d.lsClient.GetLettersByDirSend(r.Context(), &letterService.DirName{DirName: uint64(folderId)})
	}

	if err != nil {
		fmt.Println("Er", err)
		w.Write(GetFoldersError(err))
		return
	}
	fmt.Println(len(letterList.Letter))
	w.Write(ProtoLetterListAnswer(letterList))
}

// Folder AddFolder godoc
// @Summary Add folder
// @Description добавить папку в полученные (отправленные) post user/folders/{recived/sended}/folderName {folderName:"folderName"}
// @ID AddFolder
// @Accept  json
// @Produce  json
// @Param folderName body Folder true "folder name"
// @Success 200
// @Router /user/folders/recived/folderName [post]
func (d Delivery) AddFolder(w http.ResponseWriter, r *http.Request) {
	if r.Method==http.MethodPut{
		d.RenameFolder(w, r)
		return
	}
	if r.Method==http.MethodDelete{
		d.RemoveFolder(w, r)
		return
	}
	folderName := r.FormValue("folderName")
	fmt.Println("url", r.URL, strings.Contains(r.URL.Path, "received"), strings.Contains(r.URL.Path, "sended"),
		folderName)
	kind := "sended"
	if strings.Contains(r.URL.Path, "received") {
		kind = "received"
	}
	fmt.Println("KIND", kind)
	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}
	{}
	userServiceStruct:=&userService.Folder{Uid: user.Id, Name: folderName, Type: kind}
	_, er = d.usClient.CreateFolder(r.Context(), userServiceStruct)
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}
	fmt.Println("hello")
	w.Write(SuccessRespAns())
}

// Folder AddLetterInFolder godoc
// @Summary Add letter in folder
// @Description добавить писмо в папку post user/folders/{recived/sended}/folderName {folderName:"folderName", letterID: id}
// @ID AddLetterInFolder
// @Accept  json
// @Produce  json
// @Param folderName body Folder true "folder name"
// @Param letterId body int true "Letter id"
// @Success 200
// @Router /user/folders/recived/folderName/letter [put]
func (d Delivery) AddLetterInFolder(w http.ResponseWriter, r *http.Request) {
	if r.Method==http.MethodDelete{
		d.RemoveLetterInFolder(w, r)
		return
	}
	param := r.FormValue("letterId")
	lid, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.Write(GetFoldersError(err))
		return
	}
	fmt.Println("url", r.URL, strings.Contains(r.URL.Path, "received"), strings.Contains(r.URL.Path, "sended"), lid, err)
	kind := "sended"
	if strings.Contains(r.URL.Path, "received") {
		kind = "received"
	}

	fmt.Println("KIND", kind)
	folderName:=context.GetStrFormValueSafety(r, "folderName")
	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}

	folderId, er := d.usClient.GetFolderId(r.Context(), &userService.Folder{Uid: user.Id, Type: kind, Name: folderName})
	fmt.Println("FOLDER ID", folderId)
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}


	var resp *letterService.Response
	if kind == "received" {
		resp, er = d.lsClient.AddLetterToDir(r.Context(), &letterService.DirLid{
			Did:  folderId.Id,
			Lid:  lid,
			Type: true,
		})
	} else {
		resp, er = d.lsClient.AddLetterToDir(r.Context(), &letterService.DirLid{
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

// Folder RenameFolder godoc
// @Summary Rename Folder
// @Description Переименовать папку  user/folders/{recived/sended}/folderName {oldName:"oldName", newName:"newName"}
// @ID RenameFolder
// @Accept  json
// @Produce  json
// @Param folderName body Folder true "folder name"
// @Param letterId body int true "Letter id"
// @Success 200
// @Router /user/folders/recived/folderName [put]
func (d Delivery) RenameFolder(w http.ResponseWriter, r *http.Request) {
	fmt.Print("\n\n\nHUIn\n\n\n")
	oldName := r.FormValue("oldName")
	newName := r.FormValue("newName")
	fmt.Println("url", r.URL, strings.Contains(r.URL.Path, "received"), strings.Contains(r.URL.Path, "sended"), oldName, newName)
	kind := "sended"
	if strings.Contains(r.URL.Path, "received") {
		kind = "received"
	}
	fmt.Println("KIND", kind)
	fmt.Print("\n\n\nHUIn\n\n\n")
	er, user := context.GetUserFromCtx(r.Context())
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}
	_, er = d.usClient.RenameFolder(r.Context(), &userService.RenameFolderMsg{Uid: user.Id, Type: kind, OldName: oldName, NewName: newName})
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}
	fmt.Println("OK")
	w.Write(SuccessRespAns())
}

// Folder RemoveLetterInFolder godoc
// @Summary Remove Letter from Folder
// @Description Удалить письмо из папки user/folders/{recived/sended}/folderName
// @Description /user/folders/sended/folderName/letter body{letterID:Id}
// @ID RemoveLetterInFolder
// @Accept  json
// @Produce  json
// @Param folderName body Folder true "folder name"
// @Param letterId body int true "Letter id"
// @Success 200
// @Router /user/folders/sended/folderName/letter  [delete]
func (d Delivery) RemoveLetterInFolder(w http.ResponseWriter, r *http.Request) {
	param := r.FormValue("letterId")
	lid, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.Write(GetFoldersError(err))
		return
	}
	fmt.Println("url", r.URL, strings.Contains(r.URL.Path, "received"), strings.Contains(r.URL.Path, "sended"), lid)
	vars := mux.Vars(r)
	folderName := vars["folderName"]
	id, err:=strconv.Atoi(folderName)
	if err != nil {
		w.Write(GetFoldersError(err))
		return
	}

	resp, _ := d.lsClient.RemoveLetterFromDir(r.Context(), &letterService.DirLid{
		Did:  uint64(id),
		Lid:  lid,
		Type: true,
	})

	w.Write(ProtoResponseAnswer(resp))
}

// Folder RemoveFolder godoc
// @Summary Remove Folder
// @Description удалить папку delete user/folders/{recived/sended}/folderName {folderName:"folderName"}
// @ID RemoveFolder
// @Accept  json
// @Produce  json
// @Param folderName body Folder true "folder name"
// @Success 200
// @Router /user/folders/recived/folderName [delete]
func (d Delivery) RemoveFolder(w http.ResponseWriter, r *http.Request) {
	folderName := r.FormValue("folderName")
	fmt.Println("url", r.URL, strings.Contains(r.URL.Path, "received"), strings.Contains(r.URL.Path, "sended"), folderName)
	var kind bool
	textKind := "received"
	if strings.Contains(r.URL.Path, "received") {
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
	folderId, er := d.usClient.RemoveFolder(r.Context(), &userService.Folder{Uid: user.Id, Type: textKind, Name: folderName})
	fmt.Println("FOLDER ID", folderId)
	if er != nil {
		w.Write(GetFoldersError(er))
		return
	}
	resp, _ := d.lsClient.RemoveDir(r.Context(), &letterService.DirLid{
		Did:  folderId.Id,
		Type: kind,
	})
	w.Write(ProtoResponseAnswer(resp))
}
