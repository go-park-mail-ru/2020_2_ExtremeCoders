package test

import (
	"CleanArch/internal/User/UserModel"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

// a successful case
func TestShouldUpdateStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	user:=UserModel.User{
		Id: 3006865986,
		Name: "Dellvin",
		Surname: "Black",
		Email: "dellvin.black@gmail.com",
		Password: "1538",
	}
	rows := sqlmock.
		NewRows([]string{"id", "name", "surname", "email", "password","img"})

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT * FROM users where email=").
		WithArgs(user.Name).WillReturnRows(rows)

	mock.ExpectCommit()

	//UserPostgres.New(sql(db))
	//if err = recordStats(db, 2, 3); err != nil {
	//	t.Errorf("error was not expected while updating stats: %s", err)
	//}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
