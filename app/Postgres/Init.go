package Postgres

import (
	"CleanArch/app/Letter/LetterModel"
	"CleanArch/app/User/UserModel"
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

func (dbInfo *DataBase) Init() {
	if dbInfo == nil {
		dbInfo.User = "postgres"
		dbInfo.Password = "123456yhn"
		dbInfo.DataBaseName = "maila"
	}
	dbInfo.DB = pg.Connect(&pg.Options{
		User:     dbInfo.User,
		Password: dbInfo.Password,
		Database: dbInfo.DataBaseName,
	})
	fmt.Println(dbInfo.User, dbInfo.Password, dbInfo.DataBaseName)

	err := createSchema(dbInfo.DB)
	dbInfo.DB.Close()
	dbInfo.DB = pg.Connect(&pg.Options{
		User:     dbInfo.User,
		Password: dbInfo.Password,
		Database: dbInfo.DataBaseName,
	})
	if err != nil {
		panic(err)
	}
}

// createSchema creates database schemas.
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
