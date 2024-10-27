package validation_rules

import (
	"errors"
	"pb-starter/app/lib"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateUsernamePattern() validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		r := regexp.MustCompile(lib.REGEX_USERNAME)
		if valid := r.MatchString(s); !valid {
			return errors.New("username can only contain lowercase letters and numbers")
		}
		return nil
	}
}
