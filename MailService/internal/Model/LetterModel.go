package Model

type Letter struct {
	Id            uint64
	Sender        string
	Receiver      string
	Theme         string
	Text          string
	DateTime      int64
	IsWatched     bool
	DirectoryRecv uint64
	DirectorySend uint64
	Spam          bool
	Box           bool
}
