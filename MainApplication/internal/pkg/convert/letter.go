package convert

import (
	pb "Mailer/MailService/proto"
	Model "Mailer/MainApplication/internal/Letter/LetterModel"
)

func ModelToProto(letter Model.Letter) *pb.Letter {
	pbLetter := pb.Letter{
		Sender:    letter.Sender,
		Receiver:  letter.Receiver,
		Lid:       letter.Id,
		DateTime:  uint64(letter.DateTime),
		Theme:     letter.Theme,
		Text:      letter.Text,
		IsWatched: letter.IsWatched,
	}
	return &pbLetter
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
		}
		list = append(list, &pbLetter)
	}
	return list
}

func ProtoToModelList(letters []*pb.Letter) []Model.Letter {
	var list []Model.Letter
	for _, letter := range letters {
		pbLetter := Model.Letter{
			Sender:    letter.Sender,
			Receiver:  letter.Receiver,
			Id:        letter.Lid,
			DateTime:  int64(letter.DateTime),
			Theme:     letter.Theme,
			Text:      letter.Text,
			IsWatched: letter.IsWatched,
		}
		list = append(list, pbLetter)
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
	}
	return Letter
}
