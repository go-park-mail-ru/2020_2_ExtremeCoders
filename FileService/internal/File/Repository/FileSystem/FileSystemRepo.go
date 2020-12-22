package FileSystem

import (
	repo "Mailer/FileService/internal/File/Repository"
	fileProto "Mailer/FileService/proto"
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"strings"
)

type Repo struct {
}

func New() repo.Interface {
	curPath, _ := os.Getwd()
	fmt.Println("CUR PATh", curPath)
	err := os.Chdir(curPath + "/web/static")
	if err != nil {
		fmt.Println("ERROR", err)
	}
	curPath, _ = os.Getwd()
	fmt.Println("CUR PATh", curPath)
	return Repo{}
}

func (fsr Repo) SaveFiles(*fileProto.Files) error {
	fmt.Println("CALL SaveFiles")
	return nil
}

func (fsr Repo) SaveAvatar(avatar *fileProto.Avatar) error {
	fmt.Println("CALL SaveAvatar")
	split := strings.Split(avatar.FileName, ".")
	fmt.Println(split, len(split))
	temp := strings.Split(avatar.FileName, ".")
	ext := temp[len(temp)-1]
	_ = os.Mkdir(avatar.Email, 0777)
	path := avatar.Email + "/" + "avatar." + ext
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("ERROR", err)
		return repo.SaveAvatarError
	}
	defer f.Close()
	_, err = f.Write(avatar.Content)
	if err != nil {
		fmt.Println("ERROR", err)
		return repo.SaveAvatarError
	}
	return nil
}

func (fsr Repo) GetAvatar(user *fileProto.User) (*fileProto.Avatar, error) {

	files, err := ioutil.ReadDir(user.Email)
	if err != nil || len(files) == 0 {
		return &fileProto.Avatar{}, repo.GetAvatarError
	}

	fileName := files[0].Name()

	file, err := os.Open(user.Email + "/" + fileName)
	if err != nil {
		return &fileProto.Avatar{}, repo.GetAvatarError
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return &fileProto.Avatar{}, repo.GetAvatarError
	}

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		return &fileProto.Avatar{}, repo.GetAvatarError
	}

	return &fileProto.Avatar{
		Email:    user.Email,
		FileName: fileName,
		Content:  buffer.Bytes(),
	}, nil
}

func (fsr Repo) GetDefaultAvatar() (*fileProto.Avatar, error) {
	file, err := os.Open("default.jpeg")
	if err != nil {
		return &fileProto.Avatar{}, repo.GetAvatarError
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return &fileProto.Avatar{}, repo.GetAvatarError
	}

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		return &fileProto.Avatar{}, repo.GetAvatarError
	}

	return &fileProto.Avatar{
		FileName: "default.jpeg",
		Content:  buffer.Bytes(),
	}, nil
}

func (fsr Repo) GetFiles(id *fileProto.LetterId) (*fileProto.Files, error) {
	fmt.Println("CALL GetFiles")
	return &fileProto.Files{}, nil
}
