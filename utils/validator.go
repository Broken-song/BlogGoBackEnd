package utils

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// ValidateEmail 校验邮箱
func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	ok, _ := regexp.MatchString(`^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$`, email)
	if !ok {
		return false
	}
	return true
}
