package Model

type Letter struct {
	Id        uint64
	Sender    string
	Receiver  string
	Theme     string
	Text      string
	DateTime  int64
	IsWatched bool
	Directory string
}
