package valueobject

import (
	"errors"
	"regexp"
)

type Email struct {
	value string
}

var emailRegex = regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`)

func NewEmail(val string) (Email, error) {
	if !emailRegex.MatchString(val) {
		return Email{}, errors.New("invalid email format")
	}
	return Email{value: val}, nil
}

func (e Email) String() string {
	return e.value
}
