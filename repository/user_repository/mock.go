package user_repository

import (
	"tugas-sesi12/entity"
	"tugas-sesi12/pkg/errrs"
)

var (
	CreateNewUser  func(user entity.User) errrs.MessageErr
	GetUserById    func(userId int) (*entity.User, errrs.MessageErr)
	GetUserByEmail func(userEmail string) (*entity.User, errrs.MessageErr)
)

type userRepoMock struct{}

func NewUserRepoMock() UserRepository {
	return &userRepoMock{}
}

func (u *userRepoMock) CreateNewUser(user entity.User) errrs.MessageErr {
	return CreateNewUser(user)
}
func (u *userRepoMock) GetUserById(userId int) (*entity.User, errrs.MessageErr) {
	return GetUserById(userId)
}
func (u *userRepoMock) GetUserByEmail(userEmail string) (*entity.User, errrs.MessageErr) {
	return GetUserByEmail(userEmail)
}
