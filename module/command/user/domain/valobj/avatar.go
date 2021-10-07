package valobj

import "gin-ddd-demo/module/command/user/domain/common"

type Avatar struct {
	url string
}

func NewAvatar(url string) (*Avatar, error) {
	avatar := &Avatar{}
	if err := avatar.setAvatarUrl(url); err != nil {
		return nil, err
	}
	return avatar, nil
}

func (avatar *Avatar) validateAvatarUrl(url string) error {
	return nil
}

func (avatar *Avatar) setAvatarUrl(url string) error {
	if err := avatar.validateAvatarUrl(url); err != nil {
		return common.ErrAvatarUrlFormat
	}
	avatar.url = url
	return nil
}

func (avatar *Avatar) GetAvatarUrl() string {
	return avatar.url
}
