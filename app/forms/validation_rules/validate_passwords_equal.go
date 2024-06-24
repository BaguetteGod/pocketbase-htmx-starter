package validation_rules

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidatePasswordsEqual(password string) validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		if password != s {
			return errors.New("passwords don't match")
		}
		return nil
	}
}
