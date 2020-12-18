package test

//
//import (
//	"Mailer/MailService/internal/Model"
//	"Mailer/MailService/pkg/convert"
//	//letterService "Mailer/MailService/proto"
//	"log"
//	"testing"
//)
//
//var letter =Model.Letter{
//	Receiver: "kek",
//	Sender: "lol",
//	IsWatched: true,
//}
//
//
//
//func TestModelToProto(t *testing.T) {
//	pb:=convert.ModelToProto(letter)
//	if pb.Sender!=letter.Sender{
//		log.Fatalf("Error at ModelToProto")
//	}
//}
//
//
//func TestProtoToModelList(t *testing.T) {
//	var list []Model.Letter
//	list=append(list, letter)
//	letter:=convert.ModelToProtoList(&list)
//	if letter[0].Sender!=list[0].Sender{
//		log.Fatalf("Error at ProtoToModel")
//	}
//}
