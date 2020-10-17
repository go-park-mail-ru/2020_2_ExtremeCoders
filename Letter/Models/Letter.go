package Models

import (
	"CleanArch/User/Models"
)

type Letter struct {
	sender   Models.User
	receiver Models.User
	theme    string
	text     string
}
