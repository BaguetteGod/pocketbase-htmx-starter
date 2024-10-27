package forms

import (
	"pb-starter/app/forms/validation_rules"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type ProfileFormValue struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	DisplayName string `json:"displayname"`
}

func (pfv ProfileFormValue) Validate(e *core.ServeEvent, c echo.Context) error {
	user := apis.RequestInfo(c).AuthRecord
	emailChanged := user.Email() != pfv.Email
	usernameChanged := user.Username() != pfv.Username
	isOauth2User := user.GetBool("oauth")

	return validation.ValidateStruct(&pfv,
		validation.Field(&pfv.Email,
			validation.When(!isOauth2User && emailChanged,
				validation.Required,
				is.Email,
				validation.By(validation_rules.ValidateEmailUnique(e)))),
		validation.Field(&pfv.Username,
			validation.When(usernameChanged,
				validation.Length(3, 20),
				validation.Required,
				validation.By(validation_rules.ValidateUsernamePattern()),
				validation.By(validation_rules.ValidateUsernameUnique(e)))),
		validation.Field(&pfv.DisplayName,
			validation.Required,
			validation.Length(5, 20)),
	)
}

func GetProfileFormValue(c echo.Context) ProfileFormValue {
	return ProfileFormValue{
		Email:       c.FormValue("email"),
		Username:    c.FormValue("username"),
		DisplayName: c.FormValue("displayname"),
	}
}
