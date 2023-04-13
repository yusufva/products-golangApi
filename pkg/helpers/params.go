package helpers

import (
	"strconv"
	"tugas-sesi12/pkg/errrs"

	"github.com/gin-gonic/gin"
)

func GetParamsId(c *gin.Context, key string) (int, errrs.MessageErr) {
	value := c.Param(key)

	id, err := strconv.Atoi(value)

	if err != nil {
		return 0, errrs.NewBadRequest("invalid parameter id")
	}

	return id, nil
}
