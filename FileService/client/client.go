package main

import (
	fileProto "Mailer/FileService/proto"
	"bytes"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func main() {

	grcpConn, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConn.Close()

	fileManager := fileProto.NewFileServiceClient(grcpConn)

	ctx := context.Background()


	file, err := os.Open("default.jpeg")
	if err != nil {
		fmt.Printf("ERR1", err.Error())
	}

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("ERR2", err.Error())
	}
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		fmt.Printf("err", err.Error())
	}
	//
	//var a []byte
	//file.Read(a)
	//avatar, _ := fileManager.GetAvatar(ctx, &fileProto.User{Email: "suko"})
	_,err = fileManager.SetAvatar(ctx, &fileProto.Avatar{
		Email:    "s@mail.ru",
		FileName: "text.jpeg",
		Content:  buffer.Bytes(),
	})
	fmt.Println(err)
}