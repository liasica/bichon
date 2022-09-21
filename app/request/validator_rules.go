package request

import (
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/go-playground/validator/v10"
)

func validateAddress(fl validator.FieldLevel) bool {
    return model.IsValidAddress(fl.Field().String())
}

func validateEnum(fl validator.FieldLevel) bool {
    value := fl.Field().Interface().(model.Enum)
    return value.IsValid()
}
