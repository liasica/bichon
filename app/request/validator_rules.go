package request

import (
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/go-playground/validator/v10"
)

func isValidAddress(fl validator.FieldLevel) bool {
    return model.IsValidAddress(fl.Field().String())
}
