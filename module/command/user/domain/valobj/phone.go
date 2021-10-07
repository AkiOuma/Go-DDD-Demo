package valobj

import (
	"gin-ddd-demo/module/command/user/domain/common"
	"regexp"
)

type Phone struct {
	number string
}

func NewPhone(number string) (*Phone, error) {
	phone := &Phone{}
	if err := phone.setPhone(number); err != nil {
		return nil, err
	}
	return phone, nil
}

func (phone *Phone) validatePhone(phoneNumber string) error {
	pattern := regexp.MustCompile(`((\d{11})|^((\d{7,8})|(\d{4}|\d{3})-(\d{7,8})|(\d{4}|\d{3})-(\d{7,8})-(\d{4}|\d{3}|\d{2}|\d{1})|(\d{7,8})-(\d{4}|\d{3}|\d{2}|\d{1}))$)`)
	if !pattern.MatchString(phoneNumber) {
		return common.ErrPhoneFormat
	}
	return nil
}

func (phone *Phone) setPhone(phoneNumber string) error {
	if err := phone.validatePhone(phoneNumber); err != nil {
		return err
	}
	phone.number = phoneNumber
	return nil
}

func (phone *Phone) GetPhoneNumber() string {
	return phone.number
}
