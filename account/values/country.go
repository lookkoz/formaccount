package values

import (
	"errors"
	"fmt"
)

type CountryCode string

// String returns the string representation of the country code
func (x CountryCode) String() string {
	return string(x)
}

// Validate checks the value of the country code
func (x CountryCode) Validate() error {
	// todo : proper country validation
	if len([]rune(x)) != 2 {
		return errors.New(fmt.Sprintf("current code must be 2 characters (have `%s`)", x.String()))
	}
	return nil
}
