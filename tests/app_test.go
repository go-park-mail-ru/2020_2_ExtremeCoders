package tests

import (
	"fmt"
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &ItemRepository{
		DB: db,
	}

	title := "title"
	descr := "description"
	testItem := &Item{
		Title:       title,
		Description: descr,
	}

	//ok query
	mock.
		ExpectExec(`INSERT INTO items`).
		WithArgs(title, descr).
		WillReturnResult(sqlmock.NewResult(1, 1))

	id, err := repo.Add(testItem)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if id != 1 {
		t.Errorf("bad id: want %v, have %v", id, 1)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// query error
	mock.
		ExpectExec(`INSERT INTO items`).
		WithArgs(title, descr).
		WillReturnError(fmt.Errorf("bad query"))

	_, err = repo.Add(testItem)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// result error
	mock.
		ExpectExec(`INSERT INTO items`).
		WithArgs(title, descr).
		WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("bad_result")))

	_, err = repo.Add(testItem)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// last id error
	mock.
		ExpectExec(`INSERT INTO items`).
		WithArgs(title, descr).
		WillReturnResult(sqlmock.NewResult(0, 0))

	_, err = repo.Add(testItem)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}