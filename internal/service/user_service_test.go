package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"gotu/bookstore/internal/repository/mocks"
	"gotu/bookstore/internal/request"
	"gotu/bookstore/internal/types"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	repository := mocks.NewUserRepository(t)

	t.Run("success", func(t *testing.T) {
		// mock signUp request
		var signupReq request.SignUpRequest
		if err := faker.FakeData(&signupReq); err != nil {
			t.Errorf("err: %v", err)
		}

		// mock user
		var user1 types.User
		if err := faker.FakeData(&user1); err != nil {
			t.Errorf("err: %v", err)
		}

		repository.On("CreateUser", mock.Anything, mock.Anything).Return(&user1, nil).Once()
		service := NewUserService(repository)

		// create
		res, err := service.CreateUser(context.TODO(), &signupReq)

		if !reflect.DeepEqual(&user1, res) {
			t.Errorf("res and mockEntity are not deep equal")
		}

		// assert
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("error - create user", func(t *testing.T) {
		// mock signUp request
		var signupReq request.SignUpRequest
		if err := faker.FakeData(&signupReq); err != nil {
			t.Errorf("err: %v", err)
		}

		repository.On("CreateUser", mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
		service := NewUserService(repository)

		// create
		res, err := service.CreateUser(context.TODO(), &signupReq)

		// assert
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
