package Postgres

import (
	"CleanArch/User/Models"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type DataBase struct {
	db *pg.DB
	User string
	Password string
}

func (dbInfo *DataBase)Init() {
	if dbInfo ==nil{
		dbInfo.User="postgres"
		dbInfo.Password="1538"
	}
	dbInfo.db = pg.Connect(&pg.Options{
		User:     dbInfo.User,
		Password: dbInfo.Password,
	})
	defer dbInfo.db.Close()

	err := createSchema(dbInfo.db)
	if err != nil {
		panic(err)
	}
}


// createSchema creates database schemas.
func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*Models.User)(nil),
		(*Models.Session)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
