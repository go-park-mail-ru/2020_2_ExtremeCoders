package context

import (
	"CleanArch/internal/User/UserModel"
	"context"
	"errors"
)

var UserFromContextError=errors.New("Could not get user from context!")

type userKey struct{
}

func GetUserFromCtx(ctx context.Context) (error,UserModel.User) {
	ctxUser := ctx.Value(userKey{})
	user, ok := ctxUser.(UserModel.User)
	if !ok {
		return UserFromContextError, UserModel.User{}
	}
	return nil, user
}

func SaveUserToContext(ctx context.Context, user UserModel.User) context.Context{
	return context.WithValue(ctx, userKey{}, user)
}
