package entity

import (
	"gin-ddd-demo/module/command/user/domain/common"
	"gin-ddd-demo/module/command/user/domain/valobj"
)

type UserEntity struct {
	userID           string
	userName         string
	password         *valobj.Password
	email            *valobj.Email
	phone            *valobj.Phone
	organizationCode string
	avatar           *valobj.Avatar
	status           bool
}

func NewUserEntity(
	userID string,
	userName string,
	password *valobj.Password,
	email *valobj.Email,
	phone *valobj.Phone,
	organizationCode string,
	avatar *valobj.Avatar,
	status bool,
) *UserEntity {
	return &UserEntity{
		userID:           userID,
		userName:         userName,
		password:         password,
		email:            email,
		phone:            phone,
		organizationCode: organizationCode,
		avatar:           avatar,
		status:           status,
	}
}

func (entity *UserEntity) ChangeUsername(name string) {
	entity.userName = name
}

func (entity *UserEntity) Active() {
	entity.status = true
}

func (entity *UserEntity) Deactive() {
	entity.status = false
}

func (entity *UserEntity) ChangePassword(oldPassword, newPassword string) error {
	if err := entity.ValidatePassword(oldPassword); err != nil {
		return err
	}
	return entity.password.ChangePassword(newPassword)
}

func (entity *UserEntity) ChangeEmail(newAddress string) error {
	return entity.email.UpdateEmail(newAddress)
}

func (entity *UserEntity) ChangePhone(newNumber string) error {
	return entity.phone.UpdatePhone(newNumber)
}

func (entity *UserEntity) ChangeOrganization(code string) {
	entity.organizationCode = code
}

func (entity *UserEntity) ChangeAvatar(url string) error {
	return entity.avatar.ChangeAvatar(url)
}

func (entity *UserEntity) ValidatePassword(password string) error {
	if password != entity.password.GetPassword() {
		return common.ErrPassword
	}
	return nil
}

func (entity *UserEntity) ResetPassword() string {
	return entity.password.ResetPassword()
}

// Getter
func (entity *UserEntity) UserID() string {
	return entity.userID
}

func (entity *UserEntity) UserName() string {
	return entity.userName
}

func (entity *UserEntity) Password() string {
	return entity.password.GetPassword()
}

func (entity *UserEntity) Email() string {
	return entity.email.GetEmailAddress()
}

func (entity *UserEntity) Phone() string {
	return entity.phone.GetPhoneNumber()
}

func (entity *UserEntity) OrganizationCode() string {
	return entity.organizationCode
}

func (entity *UserEntity) Avatar() string {
	return entity.avatar.GetAvatarUrl()
}

func (entity *UserEntity) Status() bool {
	return entity.status
}
