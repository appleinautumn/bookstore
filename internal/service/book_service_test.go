package service

import (
	"context"
	"errors"
	"os"
	"reflect"
	"testing"

	"gotu/bookstore/internal/repository/mocks"
	"gotu/bookstore/internal/types"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	os.Setenv("APP_ENV", "test")
	os.Exit(m.Run())
}

func TestList(t *testing.T) {
	repository := mocks.NewBookRepository(t)

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

	t.Run("success", func(t *testing.T) {
		mockList := []*types.Book{
			&b1,
			&b2,
		}

		repository.On("List", mock.Anything).Return(mockList, nil).Once()
		service := NewBookService(repository)

		// get list
		res, err := service.List(context.TODO())

		if !reflect.DeepEqual(mockList, res) {
			t.Errorf("res and mockList are not deep equal")
		}

		// assert
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("error", func(t *testing.T) {
		repository.On("List", mock.Anything).Return(nil, errors.New("error")).Once()
		service := NewBookService(repository)

		// get list
		res, err := service.List(context.TODO())

		// assert
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
