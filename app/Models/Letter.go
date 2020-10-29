package Models

type Letter struct {
	Id       uint64
	Sender   string
	Receiver string
	Theme    string
	Text     string
	DateTime int64
}
