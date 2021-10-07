package valobj

import (
	"gin-ddd-demo/module/command/user/domain/common"
	"regexp"
)

type Email struct {
	address string
}

func NewEmail(address string) (*Email, error) {
	email := &Email{}
	if err := email.setEmail(address); err != nil {
		return nil, err
	}
	return email, nil
}

func (email *Email) validateEmail(address string) error {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)
	if !pattern.MatchString(address) {
		return common.ErrEmailFormat
	}
	return nil
}

func (email *Email) setEmail(address string) error {
	if err := email.validateEmail(address); err != nil {
		return err
	}
	email.address = address
	return nil
}

func (email *Email) UpdateEmail(newAddress string) error {
	return email.setEmail(newAddress)
}

func (email *Email) GetEmailAddress() string {
	return email.address
}
