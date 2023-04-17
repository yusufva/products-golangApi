package user_pg

import (
	"strings"
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
		if strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint") {
			return errrs.NewConflictError("this email has been used")
		}
		return errrs.NewInternalServerError("something went wrong")
	}

	return nil
}
func (u *userPG) GetUserById(userId int) (*entity.User, errrs.MessageErr) {
	return nil, nil
}
func (u *userPG) GetUserByEmail(userEmail string) (*entity.User, errrs.MessageErr) {
	user := entity.User{Email: userEmail}
	err := u.db.First(&user, "email = ?", user.Email).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errrs.NewNotFoundError("user not found")
		}
		return nil, errrs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}
