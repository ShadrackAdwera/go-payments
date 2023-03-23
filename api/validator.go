package api

import (
	"github.com/ShadrackAdwera/go-payments/utils"
	"github.com/go-playground/validator/v10"
)

var validRole validator.Func = func(fl validator.FieldLevel) bool {
	if role, ok := fl.Field().Interface().(string); ok {
		return utils.IsSupportedRole(role)
	}

	return false

}
