package service

import (
	"net/http"
	"strings"
	"testing"
	"tugas-sesi12/dto"
	"tugas-sesi12/entity"
	"tugas-sesi12/pkg/errrs"
	"tugas-sesi12/repository/user_repository"

	"github.com/stretchr/testify/assert"
)

func TestUserService_CreateNewUser_Success(t *testing.T) {
	userRepo := user_repository.NewUserRepoMock()

	userService := NewUserService(userRepo)

	payload := dto.NewUserRequest{
		Email:    "test@mail.com",
		Password: "123456",
	}

	user_repository.CreateNewUser = func(user entity.User) errrs.MessageErr {
		return nil
	}

	response, err := userService.CreateNewUser(payload)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, "success", response.Result)
	assert.Equal(t, "user registered successfully", response.Message)
}

func TestUserService_CreateNewUser_InvalidRequestBodyError(t *testing.T) {
	userRepo := user_repository.NewUserRepoMock()

	tt := []struct {
		name        string
		payload     dto.NewUserRequest
		expectation errrs.MessageErr
	}{
		{
			name: "invalid email",
			payload: dto.NewUserRequest{
				Email:    "",
				Password: "123456",
			},
			expectation: errrs.NewBadRequest("email cannot be empty"),
		},
		{
			name: "invalid password",
			payload: dto.NewUserRequest{
				Email:    "test@mail.com",
				Password: "",
			},
			expectation: errrs.NewBadRequest("password cannot be empty"),
		},
	}

	for _, eachTest := range tt {
		t.Run(eachTest.name, func(t *testing.T) {
			userService := NewUserService(userRepo)
			response, err := userService.CreateNewUser(eachTest.payload)

			assert.NotNil(t, err)
			assert.Nil(t, response)
			assert.Equal(t, eachTest.expectation.Status(), err.Status())
			assert.Equal(t, eachTest.expectation.Message(), err.Message())
			assert.Equal(t, eachTest.expectation.Error(), err.Error())
		})
	}

}

func TestUserService_CreateNewUser_HashPasswordError(t *testing.T) {
	userService := NewUserService(nil)

	longChar := strings.Repeat("a", 73)

	payload := dto.NewUserRequest{
		Email:    "test@mail.com",
		Password: longChar,
	}

	response, err := userService.CreateNewUser(payload)

	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, "something went wrong", err.Message())
	assert.Equal(t, "INTERNAL_SERVER_ERROR", err.Error())
}

func TestUserService_CreateNewUser_UserRepoError(t *testing.T) {
	userRepo := user_repository.NewUserRepoMock()

	userService := NewUserService(userRepo)

	payload := dto.NewUserRequest{
		Email:    "test@mail.com",
		Password: "123456",
	}

	user_repository.CreateNewUser = func(user entity.User) errrs.MessageErr {
		return errrs.NewInternalServerError("something went wrong")
	}

	response, err := userService.CreateNewUser(payload)

	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, "something went wrong", err.Message())
	assert.Equal(t, "INTERNAL_SERVER_ERROR", err.Error())
}
