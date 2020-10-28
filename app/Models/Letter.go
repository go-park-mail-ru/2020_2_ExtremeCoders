package Models

type Letter struct {
	sender   User
	receiver User
	theme    string
	text     string
}