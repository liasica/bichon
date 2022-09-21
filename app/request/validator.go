package request

import (
    "errors"
    "github.com/chatpuppy/puppychat/app/model"
    enLocales "github.com/go-playground/locales/en"
    ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
    enTranslations "github.com/go-playground/validator/v10/translations/en"
    "strings"
)

type Validator struct {
    validator *validator.Validate
    trans     ut.Translator
}

func (cv *Validator) Validate(i any) error {
    err := cv.validator.Struct(i)
    if err != nil {
        errs := err.(validator.ValidationErrors)
        var msg []string
        for _, e := range errs {
            msg = append(msg, e.Translate(cv.trans))
        }
        err = errors.New(strings.Join(msg, "，"))
    }
    return err
}

// NewValidator Create Validator
func NewValidator() *Validator {
    validate := validator.New()
    zh := enLocales.New()
    uni := ut.New(zh, zh)
    trans, _ := uni.GetTranslator("en")
    _ = enTranslations.RegisterDefaultTranslations(validate, trans)

    cv := &Validator{validator: validate, trans: trans}

    // register rules
    cv.customValidation("address", validateAddress, model.ErrInvalidAddress.Error())
    cv.customValidation("enum", validateEnum, model.ErrInvaildArgument.Error())
    return cv
}

func (cv *Validator) customValidation(tag string, fn validator.Func, message ...string) {
    _ = cv.validator.RegisterValidation(tag, fn)
    _ = cv.validator.RegisterTranslation(
        tag,
        cv.trans,
        func(ut ut.Translator) error {
            text := "{0} validate failed"
            if len(message) > 0 {
                text = message[0]
            }
            return ut.Add(tag, text, true)
        }, func(ut ut.Translator, fe validator.FieldError) string {
            t, _ := ut.T(tag, fe.Field())
            return t
        },
    )
}
