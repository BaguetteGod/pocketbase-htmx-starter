package forms

import (
	"pb-starter/app/forms/validation_rules"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v5"
)

type ConfirmPasswordResetFormValue struct {
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
	Token           string `json:"token"`
}

func (cprf ConfirmPasswordResetFormValue) Validate() error {
	return validation.ValidateStruct(&cprf,
		validation.Field(&cprf.Token, validation.Required),
		validation.Field(&cprf.Password, validation.Required),
		validation.Field(&cprf.PasswordConfirm, validation.Required, validation.By(validation_rules.ValidatePasswordsEqual(cprf.Password))),
	)
}

func GetConfirmPasswordResetForm(c echo.Context) ConfirmPasswordResetFormValue {
	return ConfirmPasswordResetFormValue{
		Password:        c.FormValue("password"),
		PasswordConfirm: c.FormValue("passwordConfirm"),
		Token:           c.FormValue("token"),
	}
}
