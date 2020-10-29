package Postgres

import (
	"CleanArch/app/Models"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type DataBase struct {
	db *pg.DB
	User string
	Password string
	DataBaseName string
}

func (dbInfo *DataBase)Init() {
	if dbInfo ==nil{
		dbInfo.User="postgres"
		dbInfo.Password="1538"
		dbInfo.DataBaseName="maila"
	}
	dbInfo.db = pg.Connect(&pg.Options{
		User:     dbInfo.User,
		Password: dbInfo.Password,
		Database: dbInfo.DataBaseName,
	})


	err := createSchema(dbInfo.db)
	dbInfo.db.Close()
	dbInfo.db = pg.Connect(&pg.Options{
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
		(*Models.User)(nil),
		(*Models.Session)(nil),
		(*Models.Letter)(nil),
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
