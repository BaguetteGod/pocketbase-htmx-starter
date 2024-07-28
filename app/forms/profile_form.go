package forms

import (
	"date-rate/app/forms/validation_rules"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

type ProfileFormValue struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (rfv ProfileFormValue) Validate(e *core.ServeEvent) error {
	return validation.ValidateStruct(&rfv,
		validation.Field(&rfv.Email, validation.Required, is.Email, validation.By(validation_rules.ValidateEmailUnique(e))),
		validation.Field(&rfv.Username, validation.Required, validation.Length(5, 30), validation.By(validation_rules.ValidateUsernameUnique(e))),
	)
}

func GetProfileFormValue(c echo.Context) ProfileFormValue {
	return ProfileFormValue{
		Email:    c.FormValue("email"),
		Username: c.FormValue("username"),
	}
}
