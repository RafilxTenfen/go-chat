package cli

import (
	"errors"
	"strings"
)

// UserError extracts the error's message and turn it into user friendly
func UserError(err error) error {
	if err == nil {
		return nil
	}

	s := err.Error()
	s = strings.ToUpper(s[:1]) + s[1:]

	if s[len(s)-1] != '.' {
		s += "."
	}

	return errors.New(s)
}
