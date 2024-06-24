package validation_rules

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pocketbase/pocketbase/core"
)

func ValidateUsernameUnique(e *core.ServeEvent) validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		user, _ := e.App.Dao().FindAuthRecordByUsername("users", s)
		if user != nil {
			return errors.New("username is already in use")
		}
		return nil
	}
}
