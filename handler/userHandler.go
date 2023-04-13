package handler

import (
	"tugas-sesi12/dto"
	"tugas-sesi12/pkg/errrs"
	"tugas-sesi12/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func (uh *userHandler) Register(c *gin.Context) {
	var newUserRequest dto.NewUserRequest

	if err := c.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJson := errrs.NewUnprocessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := uh.userService.CreateNewUser(newUserRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(result.StatusCode, result)
}

func (uh *userHandler) Login(c *gin.Context) {
	var newUserRequest dto.NewUserRequest

	if err := c.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJson := errrs.NewUnprocessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := uh.userService.Login(newUserRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(result.StatusCode, result)
}
