package user_repository

import (
	"tugas-sesi12/entity"
	"tugas-sesi12/pkg/errrs"
)

type UserRepository interface {
	CreateNewUser(user entity.User) errrs.MessageErr
	GetUserById(userId int) (*entity.User, errrs.MessageErr)
	GetUserByEmail(userEmail string) (*entity.User, errrs.MessageErr)
}
