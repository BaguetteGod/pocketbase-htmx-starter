package validation_rules

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pocketbase/pocketbase/core"
)

func ValidateEmailUnique(e *core.ServeEvent) validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		user, _ := e.App.Dao().FindAuthRecordByEmail("users", s)
		if user != nil {
			return errors.New("email is already in use")
		}
		return nil
	}
}
