package validation_rules

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pocketbase/pocketbase/core"
)

func ValidateUserPassword(e *core.ServeEvent, password string) validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		user, _ := e.App.Dao().FindAuthRecordByEmail("users", s)
		if !user.ValidatePassword(password) {
			return errors.New("incorrect email or password")
		}
		return nil
	}
}
