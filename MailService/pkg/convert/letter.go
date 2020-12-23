package convert

import (
	"Mailer/MailService/internal/Model"
	pb "Mailer/MailService/proto"
)

func ModelToProto(letter Model.Letter) pb.Letter {
	return pb.Letter{
		Sender:    letter.Sender,
		Receiver:  letter.Receiver,
		Lid:       letter.Id,
		DateTime:  uint64(letter.DateTime),
		Theme:     letter.Theme,
		Text:      letter.Text,
		IsWatched: letter.IsWatched,
		Spam: letter.Spam,
		Box: letter.Box,
		DirectoryRecv: int64(letter.DirectoryRecv),
	}
}

func ModelToProtoList(letters *[]Model.Letter) []*pb.Letter {
	var list []*pb.Letter
	for _, letter := range *letters {
		pbLetter := pb.Letter{
			Sender:    letter.Sender,
			Receiver:  letter.Receiver,
			Lid:       letter.Id,
			DateTime:  uint64(letter.DateTime),
			Theme:     letter.Theme,
			Text:      letter.Text,
			IsWatched: letter.IsWatched,
			DirectoryRecv: int64(letter.DirectoryRecv),
			Spam: letter.Spam,
			Box: letter.Box,
		}
		list = append(list, &pbLetter)
	}
	return list
}

func ProtoToModel(letter *pb.Letter) Model.Letter {
	Letter := Model.Letter{
		Sender:    letter.Sender,
		Receiver:  letter.Receiver,
		Id:        letter.Lid,
		DateTime:  int64(letter.DateTime),
		Theme:     letter.Theme,
		Text:      letter.Text,
		IsWatched: letter.IsWatched,
		Spam: letter.Spam,
		Box: letter.Box,
		DirectoryRecv: uint64(letter.DirectoryRecv),
	}
	return Letter
}
