package helpers

import (
	"tugas-sesi12/pkg/errrs"

	"github.com/asaskevich/govalidator"
)

func ValidateStruct(payload interface{}) errrs.MessageErr {
	_, err := govalidator.ValidateStruct(payload)

	if err != nil {
		return errrs.NewBadRequest(err.Error())
	}

	return nil
}
