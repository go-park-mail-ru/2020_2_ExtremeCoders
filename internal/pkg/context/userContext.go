package context

import (
	"CleanArch/config"
	"CleanArch/internal/User/UserModel"
	"context"
	crypto "crypto/rand"
	"errors"
	"math/big"
)

var UserFromContextError = errors.New("Could not get user from context!")

type userKey struct {
}

func GetUserFromCtx(ctx context.Context) (error, UserModel.User) {
	ctxUser := ctx.Value(userKey{})
	user, ok := ctxUser.(UserModel.User)
	if !ok {
		return UserFromContextError, UserModel.User{}
	}
	return nil, user
}

func SaveUserToContext(ctx context.Context, user UserModel.User) context.Context {
	return context.WithValue(ctx, userKey{}, user)
}

func GenerateCSRF() string{
	var token string
	for i:=0;i<config.CsrfSize;i++ {
		pos, _ := crypto.Int(crypto.Reader, big.NewInt(int64(len(config.SidRunes))))
		token+=string(config.SidRunes[pos.Int64()])
	}
	return token
}