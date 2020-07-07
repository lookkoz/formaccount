package values

import (
	"errors"
	"fmt"
)

type Currency string

// String returns the string representation of the currency
func (x Currency) String() string {
	return string(x)
}

// Validate checks the value of the currency
func (x Currency) Validate() error {
	// todo : proper currency validation
	if len([]rune(x)) != 3 {
		return errors.New(fmt.Sprintf("current code must be 3 characters (have `%s`)", x.String()))
	}
	return nil
}
