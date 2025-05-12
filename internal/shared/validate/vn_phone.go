package validate

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateVietnamesePhoneNumber() validator.Func {
	return func(fl validator.FieldLevel) bool {
		phoneNumber := fl.Field().String()
		pattern := `^(0|\+?84)(3[2-9]|5[6|8|9]|7[0|6-9]|8[1-6|8]|9[0-4|6-9])\d{7}$`
		matched, _ := regexp.MatchString(pattern, phoneNumber)
		return matched
	}
}
