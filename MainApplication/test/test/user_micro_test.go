package test

import (
	"Mailer/MainApplication/internal/User/UserModel"
	"Mailer/MainApplication/internal/User/UserRepository/UserMicroservice"
	mock "Mailer/MainApplication/test/mock_UserProto"
	userService "Mailer/UserService/proto"
	"context"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestIsEmailExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id:       123,
		Name:     "Dellvin",
		Surname:  "Black",
		Email:    "dellvin.black@gmail.com",
		Password: "1538",
	}
	mockLetter := mock.NewMockUserServiceClient(ctrl)
	ctx := context.Background()
	mockLetter.EXPECT().IsEmailExists(ctx, &userService.Email{Email: user.Email}).Times(1)
	uc := UserMicroservice.New(mockLetter)

	_ = uc.IsEmailExists(user.Email)
}

func TestAddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id:       123,
		Name:     "Dellvin",
		Surname:  "Black",
		Email:    "dellvin.black@gmail.com",
		Password: "1538",
	}
	mockLetter := mock.NewMockUserServiceClient(ctrl)
	ctx := context.Background()
	mockLetter.EXPECT().AddUser(ctx, &userService.User{Email: user.Email,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
		Uid:      user.Id}).Times(1)
	uc := UserMicroservice.New(mockLetter)

	_ = uc.AddUser(&user)
}

func TestAddSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id:       123,
		Name:     "Dellvin",
		Surname:  "Black",
		Email:    "dellvin.black@gmail.com",
		Password: "1538",
	}
	mockLetter := mock.NewMockUserServiceClient(ctrl)
	ctx := context.Background()
	u := userService.User{
		Email:    user.Email,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
		Uid:      user.Id,
	}
	sid := "asjhdflashdbfp"
	msg := userService.AddSessionMsg{
		Sid:  sid,
		User: &u,
	}
	mockLetter.EXPECT().AddSession(ctx, &msg).Times(1)
	uc := UserMicroservice.New(mockLetter)

	_ = uc.AddSession(sid, user.Id, &user)
}

func TestGetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id:       123,
		Name:     "Dellvin",
		Surname:  "Black",
		Email:    "dellvin.black@gmail.com",
		Password: "1538",
	}
	userSer := userService.User{
		Uid:      123,
		Name:     "Dellvin",
		Surname:  "Black",
		Email:    "dellvin.black@gmail.com",
		Password: "1538",
	}
	mockLetter := mock.NewMockUserServiceClient(ctrl)
	ctx := context.Background()
	mockLetter.EXPECT().GetUserByEmail(ctx, &userService.Email{Email: user.Email}).Return(&userSer, nil)
	uc := UserMicroservice.New(mockLetter)

	_, _ = uc.GetUserByEmail(user.Email)
}

func TestGetUserByUID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id:       123,
		Name:     "Dellvin",
		Surname:  "Black",
		Email:    "dellvin.black@gmail.com",
		Password: "1538",
	}
	userSer := userService.User{
		Uid:      123,
		Name:     "Dellvin",
		Surname:  "Black",
		Email:    "dellvin.black@gmail.com",
		Password: "1538",
	}
	mockLetter := mock.NewMockUserServiceClient(ctrl)
	ctx := context.Background()
	mockLetter.EXPECT().GetUserByUID(ctx, &userService.Uid{Uid: user.Id}).Return(&userSer, nil)
	uc := UserMicroservice.New(mockLetter)

	_, _ = uc.GetUserByUID(user.Id)
}

func TestIsOkSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var uid userService.Uid
	uid.Uid = 1876543

	mockLetter := mock.NewMockUserServiceClient(ctrl)
	ctx := context.Background()
	sid := "asdfaweyurgoaeyf"
	mockLetter.EXPECT().IsOkSession(ctx, &userService.Sid{Sid: string(sid)}).Return(&uid, nil)
	uc := UserMicroservice.New(mockLetter)
	_, _ = uc.IsOkSession(string(sid))
}

func TestUpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := UserModel.User{
		Id:       123,
		Name:     "Dellvin",
		Surname:  "Black",
		Email:    "dellvin.black@gmail.com",
		Password: "1538",
	}
	u := userService.User{
		Email:    user.Email,
		Name:     user.Name,
		Surname:  user.Surname,
		Password: user.Password,
		Uid:      user.Id,
	}

	var nothing userService.Nothing

	mockLetter := mock.NewMockUserServiceClient(ctrl)
	ctx := context.Background()

	msg := userService.UpdateProfileMsg{
		Email:   user.Email,
		NewUser: &u,
	}
	mockLetter.EXPECT().UpdateProfile(ctx, &msg).Return(&nothing, nil)
	uc := UserMicroservice.New(mockLetter)

	_ = uc.UpdateProfile(user, user.Email)
}

func TestRemoveSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var uid userService.Uid
	uid.Uid = 1876543
	sid := "alshdifbasdf"
	mockLetter := mock.NewMockUserServiceClient(ctrl)
	ctx := context.Background()

	mockLetter.EXPECT().RemoveSession(ctx, &userService.Sid{Sid: sid}).Return(&uid, nil)
	uc := UserMicroservice.New(mockLetter)

	_, _ = uc.RemoveSession(sid)
}

func TestGetSessionByUID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var uid userService.Uid
	uid.Uid = 1876543
	var sid = userService.Sid{
		Sid: "alshdifbasdf",
	}

	mockLetter := mock.NewMockUserServiceClient(ctrl)
	ctx := context.Background()

	mockLetter.EXPECT().GetSessionByUID(ctx, &userService.Uid{Uid: uid.Uid}).Return(&sid, nil)
	uc := UserMicroservice.New(mockLetter)

	_, _ = uc.GetSessionByUID(uid.Uid)
}
