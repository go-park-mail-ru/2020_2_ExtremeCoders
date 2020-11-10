package Postgres

import (
	"CleanArch/internal/Letter/LetterModel"
	"CleanArch/internal/User/UserModel"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type DataBase struct {
	DB           *pg.DB
	User         string
	Password     string
	DataBaseName string
}

func (dbInfo *DataBase) Init(user string, password string, name string) (*pg.DB, error) {
	dbInfo.User = user
	dbInfo.Password = password
	dbInfo.DataBaseName = name

	dbInfo.DB = pg.Connect(&pg.Options{
		User:     dbInfo.User,
		Password: dbInfo.Password,
		Database: dbInfo.DataBaseName,
	})
	fmt.Println(dbInfo.User, dbInfo.Password, dbInfo.DataBaseName)
	err:=createSchema(dbInfo.DB)
	dbInfo.DB = pg.Connect(&pg.Options{
		User:     dbInfo.User,
		Password: dbInfo.Password,
		Database: dbInfo.DataBaseName,
	})
	return dbInfo.DB, err
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*UserModel.User)(nil),
		(*UserModel.Session)(nil),
		(*LetterModel.Letter)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
