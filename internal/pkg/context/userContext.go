package context

import (
	"CleanArch/internal/User/UserModel"
	"context"
)

type userKey struct {}

func GetUserFromCtx(ctx context.Context) UserModel.User {
	ctxUser := ctx.Value(userKey{})
	user, ok := ctxUser.(UserModel.User)
	if !ok {
		//херово
	}
	return user
}

func SaveUserToContext(ctx context.Context, user UserModel.User) {
	context.WithValue(ctx, userKey{}, user)
}