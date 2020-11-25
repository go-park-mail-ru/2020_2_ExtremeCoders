package SmtpManager

import (
	pb "smtpTest/proto"
	"smtpTest/internal/SendLetters"
)

type Manager struct{

}

func (m Manager)GetLettersByDir(letter *pb.Letter) *pb.Response{
	err:=SendLetters.SendLetter(letter)
	resp:=pb.Response{Ok: true, Description: err.Error()}
	if err!=nil{
		resp.Ok=false
	}
	return &resp
}
