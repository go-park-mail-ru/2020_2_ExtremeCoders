package context

import (
	"Mailer/config"
	"Mailer/MainApplication/internal/User/UserModel"

	"context"
	crypto "crypto/rand"
	"errors"
	"github.com/microcosm-cc/bluemonday"
	"math/big"
	"net/http"
	"time"
)

const (
	CookieName     = "session_id"
	CsrfCookieName = "token"
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

func GenerateCSRF() string {
	var token string
	for i := 0; i < config.CsrfSize; i++ {
		pos, _ := crypto.Int(crypto.Reader, big.NewInt(int64(len(config.SidRunes))))
		token += string(config.SidRunes[pos.Int64()])
	}
	return token
}

func GetStrFormValueSafety(r *http.Request, field string) string {
	xss:=r.FormValue(field)
	p := bluemonday.UGCPolicy()
	ok:=p.Sanitize(xss)
	return ok
}

func CreateCsrfCookie() *http.Cookie {
	cookie := &http.Cookie{
		Name:    CsrfCookieName,
		Value:   GenerateCSRF(),
		Expires: time.Now().Add(15 * time.Minute),
	}
	cookie.Path = "/"
	return cookie
}
