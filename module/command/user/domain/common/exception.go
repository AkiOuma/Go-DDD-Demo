package common

import "errors"

var ErrUserNotExist = errors.New("error: user does not exist")
var ErrEmailFormat = errors.New("error: invalid email format")
var ErrPasswordFormat = errors.New("error: invalid password format")
var ErrPassword = errors.New("error: invalid password")
var ErrPhoneFormat = errors.New("error: invalid phone format")
var ErrAvatarUrlFormat = errors.New("error: invalid avatar url format")
