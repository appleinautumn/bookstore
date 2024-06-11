package repository

import (
	"context"
	"fmt"
	"math/rand/v2"
	"testing"

	"gotu/bookstore/internal/types"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-faker/faker/v4"
)

func TestListOrdersByUserId(t *testing.T) {
	ctx := context.Background()

	// mock db
	db, mock := NewMock()
	defer db.Close()

	// create order repo
	repo := NewOrderRepository(db)

	// the userID who makes the request
	userID := int64(rand.IntN(100))

	// mock order 1
	var o1 types.Order
	if err := faker.FakeData(&o1); err != nil {
		t.Errorf("err: %v", err)
	}
	o1.UserID = userID

	// mock order 2
	var o2 types.Order
	if err := faker.FakeData(&o2); err != nil {
		t.Errorf("err: %v", err)
	}
	o2.UserID = userID

	// mock list of books
	mockList := []*types.Order{
		&o1,
		&o2,
	}

	t.Run("success", func(t *testing.T) {
		sql := `SELECT id, user_id, created_at, updated_at FROM orders WHERE user_id = \$1`

		mockRows := sqlmock.NewRows([]string{"id", "user_id", "created_at", "updated_at"}).
			AddRow(o1.ID, o1.UserID, o1.CreatedAt, o1.UpdatedAt).
			AddRow(o2.ID, o2.UserID, o2.CreatedAt, o2.UpdatedAt)

		mock.ExpectQuery(sql).WithArgs().WillReturnRows(mockRows)

		// call ListOrdersByUserId
		res, _ := repo.ListOrdersByUserId(ctx, userID)

		// assert
		if len(res) != len(mockList) {
			t.Errorf("got: %d; want: %d", len(res), len(mockList))
		}
	})

	t.Run("error - query", func(t *testing.T) {
		sql := `SELECT id, user_id, created_at, updated_at FROM orders WHERE user_id = \$1`

		mock.ExpectQuery(sql).WillReturnError(fmt.Errorf("some error"))

		// call ListOrdersByUserId and throw error
		if _, err := repo.ListOrdersByUserId(ctx, userID); err == nil {
			t.Errorf("expecting an error, but there was none")
		}

		// make sure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}
	})

	t.Run("error - scan", func(t *testing.T) {
		sql := `SELECT id, user_id, created_at, updated_at FROM orders WHERE user_id = \$1`

		// "abc" is not a valid id
		mockRows := sqlmock.NewRows([]string{"id", "user_id", "created_at", "updated_at"}).
			AddRow("abc", o1.UserID, o1.CreatedAt, o1.UpdatedAt).
			AddRow(o2.ID, o2.UserID, o2.CreatedAt, o2.UpdatedAt)

		mock.ExpectQuery(sql).WillReturnRows(mockRows)

		// call ListOrdersByUserId and throw error
		if _, err := repo.ListOrdersByUserId(ctx, userID); err == nil {
			t.Errorf("expecting an error, but there was none")
		}

		// make sure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}
	})
}

func TestCreateOrder(t *testing.T) {
	ctx := context.Background()

	// mock db
	db, mock := NewMock()
	defer db.Close()

	// create order repo
	repo := NewOrderRepository(db)

	// mock order
	var o1 types.Order
	if err := faker.FakeData(&o1); err != nil {
		t.Errorf("err: %v", err)
	}

	t.Run("success", func(t *testing.T) {
		sql := `INSERT INTO orders \(user_id\)
				VALUES \(\$1\)
				RETURNING id, user_id, created_at, updated_at`

		mockRow := sqlmock.NewRows([]string{"id", "user_id", "created_at", "updated_at"}).
			AddRow(o1.ID, o1.UserID, o1.CreatedAt, o1.UpdatedAt)

		mock.ExpectQuery(sql).WithArgs(o1.UserID).WillReturnRows(mockRow)

		// create
		res, _ := repo.CreateOrder(ctx, &o1)

		// make sure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}

		if res.ID != o1.ID {
			t.Errorf("got: %d; want: %d", res.ID, o1.ID)
		}
		if res.UserID != o1.UserID {
			t.Errorf("got: %d; want: %d", res.UserID, o1.UserID)
		}
		if res.CreatedAt != o1.CreatedAt {
			t.Errorf("got: %v; want: %v", res.CreatedAt, o1.CreatedAt)
		}
		if res.UpdatedAt != o1.UpdatedAt {
			t.Errorf("got: %v; want: %v", res.UpdatedAt, o1.UpdatedAt)
		}
	})

	t.Run("error - scan", func(t *testing.T) {
		sql := `INSERT INTO orders \(user_id\)
				VALUES \(\$1\)
				RETURNING id, user_id, created_at, updated_at`

		// "abc" is not a valid id
		mockRow := sqlmock.NewRows([]string{"id", "user_id", "created_at", "updated_at"}).
			AddRow("abc", o1.UserID, o1.CreatedAt, o1.UpdatedAt)

		mock.ExpectQuery(sql).WithArgs(o1.UserID).WillReturnRows(mockRow)

		// create and error
		if _, err := repo.CreateOrder(ctx, &o1); err == nil {
			t.Errorf("expecting an error, but there was none")
		}

		// make sure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}
	})
}

func TestCreateOrderItem(t *testing.T) {
	ctx := context.Background()

	// mock db
	db, mock := NewMock()
	defer db.Close()

	// create order repo
	repo := NewOrderRepository(db)

	// mock order item
	var oi types.OrderItem
	if err := faker.FakeData(&oi); err != nil {
		t.Errorf("err: %v", err)
	}

	t.Run("success", func(t *testing.T) {
		sql := `INSERT INTO order_books \(order_id, book_id, quantity\)
				VALUES \(\$1, \$2, \$3\)`

		// not expecting to return anything
		mock.ExpectQuery(sql).WithArgs(oi.OrderID, oi.BookID, oi.Quantity).WillReturnRows(sqlmock.NewRows(nil))

		// create order item
		if err := repo.CreateOrderItem(ctx, &oi); err != nil {
			t.Errorf("not expecting an error: %s", err)
		}

		// make sure all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}
	})
}
