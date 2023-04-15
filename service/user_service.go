package service

import (
	"net/http"
	"tugas-sesi12/dto"
	"tugas-sesi12/entity"
	"tugas-sesi12/pkg/errrs"
	"tugas-sesi12/pkg/helpers"
	"tugas-sesi12/repository/user_repository"
)

type UserService interface {
	CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errrs.MessageErr)
	Login(payload dto.NewUserRequest) (*dto.LoginResponse, errrs.MessageErr)
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errrs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	userEntity := entity.User{
		Email:    payload.Email,
		Password: payload.Password,
	}

	err = userEntity.HashPassword()

	if err != nil {
		return nil, err
	}

	err = u.userRepo.CreateNewUser(userEntity)

	if err != nil {
		return nil, err
	}

	response := dto.NewUserResponse{
		Result:     "success",
		Message:    "user registered successfully",
		StatusCode: http.StatusCreated,
	}

	return &response, nil
}

func (u *userService) Login(payload dto.NewUserRequest) (*dto.LoginResponse, errrs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByEmail(payload.Email)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errrs.NewUnauthenticatedError("invalid email/password")
		}
		return nil, err
	}

	isValidPassword := user.ComparePassword(payload.Password)

	if !isValidPassword {
		return nil, errrs.NewUnauthenticatedError("invalid email/password")
	}

	response := dto.LoginResponse{
		Result:     "success",
		Message:    "logged in successfully",
		StatusCode: http.StatusOK,
		Data: dto.TokenResponse{
			Token: user.GenerateToken(),
		},
	}

	return &response, nil
}
