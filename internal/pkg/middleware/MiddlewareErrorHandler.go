package middleware

import (
	"CleanArch/internal/errors"
)

func authError(err error) []byte{
	switch err {
	case nil:
		return errors.GetErrorBadCsrfAns(err)
	}
	return nil
}
