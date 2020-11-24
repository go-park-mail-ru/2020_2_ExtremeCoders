package convert

import (
	"MailService/internal/Model"
	pb "MailService/proto"
)

func ModelToProto(letter Model.Letter) pb.Letter{
	pbLetter:=pb.Letter{
		Sender: letter.Sender,
		Receiver: letter.Receiver,
		Lid: letter.Id,
		DateTime: uint64(letter.DateTime),
		Theme: letter.Theme,
		Text: letter.Text,
		IsWatched: letter.IsWatched,
	}
	return pbLetter
}

func ProtoToModel(letter *pb.Letter) Model.Letter{
	Letter:=Model.Letter{
		Sender: letter.Sender,
		Receiver: letter.Receiver,
		Id: letter.Lid,
		DateTime: int64(letter.DateTime),
		Theme: letter.Theme,
		Text: letter.Text,
		IsWatched: letter.IsWatched,
	}
	return Letter
}