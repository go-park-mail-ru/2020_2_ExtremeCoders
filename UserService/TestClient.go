package main

import (
	proto "UserService/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {

	grcpFileService, err := grpc.Dial(
		"127.0.0.1:8082",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc file service")
	}
	defer grcpFileService.Close()
	fileManager := proto.NewUserServiceClient(grcpFileService)

	ctx := context.Background()
	//fid, err := fileManager.RemoveFolder(ctx, &proto.Folder{
	//	Name: "SUKO",
	//	Type: "sanded",
	//	Uid:  1,
	//})
	//
	//if err != nil {
	//	fmt.Println("ERRRR", err)
	//	return
	//}
	folders, err := fileManager.GetFoldersList(ctx, &proto.Uid{Uid: 1})
	if err != nil {
		fmt.Println("ERRRR", err)
		return
	}
	for _,value:=range folders.Res{
		fmt.Println(value)
	}
	//_, err = fileManager.RenameFolder(ctx, &proto.RenameFolderMsg{
	//	OldName: "hui",
	//	NewName: "SUKO",
	//	Type: "sanded",
	//	Uid:  1,
	//})


	//fmt.Println("FID", fid)
}
