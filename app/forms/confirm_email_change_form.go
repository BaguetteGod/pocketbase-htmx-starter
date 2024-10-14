package forms

import (
	"pb-starter/app/forms/validation_rules"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

type ConfirmEmailChangeForm struct {
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
	Token           string `json:"token"`
}

func (cecf ConfirmEmailChangeForm) Validate(e *core.ServeEvent, c echo.Context) error {
	return validation.ValidateStruct(&cecf,
		validation.Field(&cecf.Token, validation.Required),
		validation.Field(&cecf.Password, validation.Required, validation.By(validation_rules.ValidateUserPasswordByToken(e, c))),
	)
}

func GetConfirmEmailChangeForm(c echo.Context) ConfirmEmailChangeForm {
	return ConfirmEmailChangeForm{
		Password: c.FormValue("password"),
		Token:    c.FormValue("token"),
	}
}
