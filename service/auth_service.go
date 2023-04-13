package service

import (
	"tugas-sesi12/entity"
	"tugas-sesi12/pkg/errrs"
	"tugas-sesi12/repository/product_repository"
	"tugas-sesi12/repository/user_repository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	Authorization() gin.HandlerFunc
}

type authService struct {
	userRepo    user_repository.UserRepository
	productRepo product_repository.ProductRepository
}

func (a *authService) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		var invalidTokenErr = errrs.NewUnauthenticatedError("invalid token")
		bearerToken := c.GetHeader("Authorization")

		var user entity.User

		err := user.ValidateToken(bearerToken)

		if err != nil {
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}

		_, err = a.userRepo.GetUserByEmail(user.Email)

		if err != nil {
			c.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		c.Set("userData", user)

		c.Next()
	}
}
