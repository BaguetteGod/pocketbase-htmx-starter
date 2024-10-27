package forms

import (
	"pb-starter/app/forms/validation_rules"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

type RegisterFormValue struct {
	Email           string `json:"email"`
	Username        string `json:"username"`
	DisplayName     string `json:"displayname"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

func (rfv RegisterFormValue) Validate(e *core.ServeEvent) error {
	return validation.ValidateStruct(&rfv,
		validation.Field(&rfv.Email,
			validation.Required,
			is.Email,
			validation.By(validation_rules.ValidateEmailUnique(e))),
		validation.Field(&rfv.Username,
			validation.Required,
			validation.Length(3, 20),
			validation.By(validation_rules.ValidateUsernamePattern()),
			validation.By(validation_rules.ValidateUsernameUnique(e))),
		validation.Field(&rfv.DisplayName,
			validation.Required,
			validation.Length(5, 20)),
		validation.Field(&rfv.Password,
			validation.Required,
			validation.Length(8, 30)),
		validation.Field(&rfv.PasswordConfirm,
			validation.Required,
			validation.By(validation_rules.ValidatePasswordsEqual(rfv.Password))),
	)
}

func GetRegisterFormValue(c echo.Context) RegisterFormValue {
	return RegisterFormValue{
		Email:           c.FormValue("email"),
		Username:        c.FormValue("username"),
		DisplayName:     c.FormValue("displayname"),
		Password:        c.FormValue("password"),
		PasswordConfirm: c.FormValue("passwordConfirm"),
	}
}
