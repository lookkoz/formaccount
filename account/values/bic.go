package values

import (
	"errors"
	"fmt"
)

type Bic string

// String returns the string representation of the country code
func (x Bic) String() string {
	return string(x)
}

// Validate checks the value of the bic
func (x Bic) Validate() error {
	if !(len([]rune(x)) == 11 || len([]rune(x)) == 8) {
		return errors.New(fmt.Sprintf("SWIFT BIC code must be either 8 or 11 characters (have `%s`)", x.String()))
	}
	return nil
}
