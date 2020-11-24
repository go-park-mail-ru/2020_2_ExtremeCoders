package middleware

import (
	"CleanArch/MainApplication/internal/errors"
)

func authError(err error) []byte{
	switch err {
	case nil:
		return errors.GetErrorBadCsrfAns(err)
	}
	return nil
}
