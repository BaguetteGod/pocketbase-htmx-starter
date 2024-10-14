package validation_rules

import (
	"errors"
	middleware "pb-starter/app/middelware"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func ValidateUserPasswordByToken(e *core.ServeEvent, c echo.Context) validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)

		tokenCookie, _ := c.Cookie(middleware.AuthCookieName)
		user, _ := e.App.Dao().FindAuthRecordByToken(tokenCookie.Value, e.App.Settings().RecordAuthToken.Secret)

		if !user.ValidatePassword(s) {
			return errors.New("incorrect password")
		}
		return nil
	}
}
