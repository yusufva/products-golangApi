package user_pg

import (
	"tugas-sesi12/entity"
	"tugas-sesi12/pkg/errrs"
	"tugas-sesi12/repository/user_repository"

	"gorm.io/gorm"
)

type userPG struct {
	db *gorm.DB
}

func NewUserPg(db *gorm.DB) user_repository.UserRepository {
	return &userPG{
		db: db,
	}
}

func (u *userPG) CreateNewUser(user entity.User) errrs.MessageErr {
	result := u.db.Create(&user)

	if result.Error != nil {
		return errrs.NewInternalServerError("something went wrong")
	}

	return nil
}
func (u *userPG) GetUserById(userId int) (*entity.User, errrs.MessageErr) {
	return nil, nil
}
func (u *userPG) GetUserByEmail(userEmail string) (*entity.User, errrs.MessageErr) {
	user := entity.User{Email: userEmail}
	err := u.db.First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errrs.NewNotFoundError("user not found")
		}
		return nil, errrs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}
