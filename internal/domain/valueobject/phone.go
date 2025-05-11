package valueobject

import (
	"errors"
	"regexp"
)

type Phone struct {
	value string
}

var phoneRegex = regexp.MustCompile(`^[0-9\-\+\(\) ]{7,20}$`)

func NewPhone(val string) (Phone, error) {
	if !phoneRegex.MatchString(val) {
		return Phone{}, errors.New("invalid phone format")
	}
	return Phone{value: val}, nil
}

func (p Phone) String() string {
	return p.value
}
