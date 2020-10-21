package Models

type User struct {
	Id       uint64
	Name     string
	Surname  string
	Email    string
	Password string
	Img      string
}

type Session struct {
	Sid string
	User   *User `pg:"rel:has-one"`
}