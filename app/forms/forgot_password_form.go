package forms

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/labstack/echo/v5"
)

type ForgotPasswordFormValue struct {
	Email string `json:"email"`
}

func (rfv ForgotPasswordFormValue) Validate() error {
	return validation.ValidateStruct(&rfv,
		validation.Field(&rfv.Email, validation.Required, is.Email),
	)
}

func GetForgotPasswordFormValue(c echo.Context) ForgotPasswordFormValue {
	return ForgotPasswordFormValue{
		Email: c.FormValue("email"),
	}
}
