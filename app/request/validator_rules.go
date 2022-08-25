package request

import (
    "github.com/go-playground/validator/v10"
    "regexp"
)

func isValidAddress(fl validator.FieldLevel) bool {
    return regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`).MatchString(fl.Field().String())
}
