package forms

import (
	"date-rate/app/forms/validation_rules"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

type LoginFormValue struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (lfv LoginFormValue) Validate(e *core.ServeEvent) error {
	return validation.ValidateStruct(&lfv,
		validation.Field(&lfv.Email, validation.Required, is.Email, validation.By(validation_rules.ValidateUserPassword(e, lfv.Password))),
		validation.Field(&lfv.Password, validation.Required),
	)
}

func GetLoginFormValue(c echo.Context) LoginFormValue {
	return LoginFormValue{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
}
