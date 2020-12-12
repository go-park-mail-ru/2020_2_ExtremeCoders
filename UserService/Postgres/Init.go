package Postgres

import (
	"Mailer/UserService/internal/UserModel"
	"fmt"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	pgwrapper "gitlab.com/slax0rr/go-pg-wrapper"
)

type DataBase struct {
	DB           pgwrapper.DB
	User         string
	Password     string
	DataBaseName string
}

func (dbInfo *DataBase) Init(user string, password string, name string) (pgwrapper.DB, error) {
	dbInfo.User = user
	dbInfo.Password = password
	dbInfo.DataBaseName = name

	dbInfo.DB = pgwrapper.NewDB(pg.Connect(&pg.Options{
		User:     dbInfo.User,
		Password: dbInfo.Password,
		Database: dbInfo.DataBaseName,
	}))
	err := createSchema(dbInfo.DB)
	dbInfo.DB = pgwrapper.NewDB(pg.Connect(&pg.Options{
		User:     dbInfo.User,
		Password: dbInfo.Password,
		Database: dbInfo.DataBaseName,
	}))
	return dbInfo.DB, err
}

func createSchema(db pgwrapper.DB) error {
	models := []interface{}{
		(*UserModel.User)(nil),
		(*UserModel.Session)(nil),
		(*UserModel.Folder)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			fmt.Println("ERR", err)
			return err
		}
	}
	return nil
}
