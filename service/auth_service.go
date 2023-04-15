package service

import (
	"tugas-sesi12/entity"
	"tugas-sesi12/pkg/errrs"
	"tugas-sesi12/pkg/helpers"
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

func NewAuthService(userRepo user_repository.UserRepository, productRepo product_repository.ProductRepository) AuthService {
	return &authService{
		userRepo:    userRepo,
		productRepo: productRepo,
	}
}

func (a *authService) Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("userData").(entity.User)

		productId, err := helpers.GetParamsId(c, "productId")

		if err != nil {
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}

		product, err := a.productRepo.GetProductById(productId)

		if err != nil {
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if user.Level == entity.Admin {
			c.Next()
			return
		}

		if product.UserId != user.Id {
			unauthorizedErr := errrs.NewUnauthorizedError("you are not authorized to modify the product data")
			c.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}

		c.Next()
	}
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
