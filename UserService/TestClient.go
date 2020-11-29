package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	proto "UserService/proto"
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
	fid, err := fileManager.GetFolderId(ctx, &proto.Folder{
		Name: "hui",
		Type: "sanded",
		Uid:  1,
	})
	if err != nil {
		fmt.Println("ERRRR", err)
		return
	}

	_, err = fileManager.RenameFolder(ctx, &proto.RenameFolderMsg{
		OldName: "hui",
		NewName: "SUKO",
		Type: "sanded",
		Uid:  1,
	})


	fmt.Println("FID", fid)
}
