package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gotu/bookstore/internal/types"
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-faker/faker/v4"
)

func TestMain(m *testing.M) {
	os.Setenv("APP_ENV", "test")
	os.Exit(m.Run())
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("error opening a stub database connection: '%s'", err)
	}

	return db, mock
}

func TestList(t *testing.T) {
	ctx := context.Background()

	db, mock := NewMock()
	defer db.Close()
	repo := NewBookRepository(db)

	// mock book 1
	var b1 types.Book
	if err := faker.FakeData(&b1); err != nil {
		t.Errorf("err: %v", err)
	}

	// mock book 2
	var b2 types.Book
	if err := faker.FakeData(&b2); err != nil {
		t.Errorf("err: %v", err)
	}

	// mock list of books
	mockList := []*types.Book{
		&b1,
		&b2,
	}

	t.Run("success", func(t *testing.T) {
		sql := "SELECT id, title, author, description, created_at, updated_at FROM books"

		mockRows := sqlmock.NewRows([]string{"id", "title", "author", "description", "created_at", "updated_at"}).
			AddRow(b1.ID, b1.Title, b1.Author, b1.Description, b1.CreatedAt, b1.UpdatedAt).
			AddRow(b1.ID, b1.Title, b1.Author, b1.Description, b1.CreatedAt, b1.UpdatedAt)

		mock.ExpectQuery(sql).WillReturnRows(mockRows)

		// list
		res, _ := repo.List(ctx)

		// assert
		if len(res) != len(mockList) {
			t.Errorf("got: %d; want: %d", len(res), len(mockList))
		}
	})

	t.Run("error - query", func(t *testing.T) {
		sql := "SELECT id, title, author, description, created_at, updated_at FROM books"

		mock.ExpectQuery(sql).WillReturnError(fmt.Errorf("some error"))

		// list and error
		if _, err := repo.List(ctx); err == nil {
			t.Errorf("expecting an error, but there was none")
		}

		// make sure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}
	})

	t.Run("error - scan", func(t *testing.T) {
		sql := "SELECT id, title, author, description, created_at, updated_at FROM books"

		// "abc" is not a valid id
		mockRows := sqlmock.NewRows([]string{"id", "title", "author", "description", "created_at", "updated_at"}).
			AddRow("abc", b1.Title, b1.Author, b1.Description, b1.CreatedAt, b1.UpdatedAt)

		mock.ExpectQuery(sql).WillReturnRows(mockRows)

		// list and error
		if _, err := repo.List(ctx); err == nil {
			t.Errorf("expecting an error, but there was none")
		}

		// make sure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}
	})
}
