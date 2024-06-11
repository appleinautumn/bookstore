package repository

import (
	"context"
	"testing"

	"gotu/bookstore/internal/types"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-faker/faker/v4"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	// mock db
	db, mock := NewMock()
	defer db.Close()

	// create user repo
	repo := NewUserRepository(db)

	// mock user 1
	var u1 types.User
	if err := faker.FakeData(&u1); err != nil {
		t.Errorf("err: %v", err)
	}

	t.Run("success", func(t *testing.T) {
		sql := `INSERT INTO users \(email, name, password\)
				VALUES \(\$1, \$2, \$3\)
				RETURNING id, email, name, created_at, updated_at`

		mockRow := sqlmock.NewRows([]string{"id", "email", "name", "created_at", "updated_at"}).
			AddRow(u1.ID, u1.Email, u1.Name, u1.CreatedAt, u1.UpdatedAt)

		mock.ExpectQuery(sql).WithArgs(u1.Email, u1.Name, u1.Password).WillReturnRows(mockRow)

		// create
		res, _ := repo.CreateUser(ctx, &u1)

		// make sure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}

		if res.ID != u1.ID {
			t.Errorf("got: %d; want: %d", res.ID, u1.ID)
		}
		if res.Email != u1.Email {
			t.Errorf("got: %s; want: %s", res.Email, u1.Email)
		}
		if res.Name != u1.Name {
			t.Errorf("got: %s; want: %s", res.Name, u1.Name)
		}
		if res.CreatedAt != u1.CreatedAt {
			t.Errorf("got: %v; want: %v", res.CreatedAt, u1.CreatedAt)
		}
		if res.UpdatedAt != u1.UpdatedAt {
			t.Errorf("got: %v; want: %v", res.UpdatedAt, u1.UpdatedAt)
		}
	})

	t.Run("error - scan", func(t *testing.T) {
		sql := `INSERT INTO users \(email, name, password\)
				VALUES \(\$1, \$2, \$3\)
				RETURNING id, email, name, created_at, updated_at`

		// "abc" is not a valid id
		mockRow := sqlmock.NewRows([]string{"id", "email", "name", "created_at", "updated_at"}).
			AddRow("abc", u1.Email, u1.Name, u1.CreatedAt, u1.UpdatedAt)

		mock.ExpectQuery(sql).WithArgs(u1.Email, u1.Name, u1.Password).WillReturnRows(mockRow)

		// create and error
		if _, err := repo.CreateUser(ctx, &u1); err == nil {
			t.Errorf("expecting an error, but there was none")
		}

		// make sure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}
	})
}
