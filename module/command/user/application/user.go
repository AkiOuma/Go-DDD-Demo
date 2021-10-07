package application

import (
	"gin-ddd-demo/module/command/user/domain/common"
	"gin-ddd-demo/module/command/user/domain/entity"
	"gin-ddd-demo/module/command/user/domain/repo"
	"gin-ddd-demo/module/command/user/domain/valobj"
)

type UserApp struct {
	repo          repo.UserRepo
	wxworkService Wxwork
	// mailService   Mail
}

func NewUserApp(
	repo repo.UserRepo,
	wxworkService Wxwork,
) *UserApp {
	return &UserApp{
		repo:          repo,
		wxworkService: wxworkService,
	}
}

func (app *UserApp) ActiveUser(userID string) error {
	user, err := app.repo.FindUserByUserID(userID)
	if err != nil {
		return common.ErrUserNotExist
	}
	user.Active()
	return app.repo.SaveOneUser(user)
}

func (app *UserApp) DeactiveUser(userID string) error {
	user, err := app.repo.FindUserByUserID(userID)
	if err != nil {
		return common.ErrUserNotExist
	}
	user.Deactive()
	return app.repo.SaveOneUser(user)
}

func (app *UserApp) RemoveUser(userID string) error {
	user, err := app.repo.FindUserByUserID(userID)
	if err != nil {
		return common.ErrUserNotExist
	}
	return app.repo.RemoveOneUser(user)
}

func (app *UserApp) UpdateUserProfile(
	userID string,
	profile *valobj.UserProfile,
) error {
	user, err := app.repo.FindUserByUserID(userID)
	if err != nil {
		return common.ErrUserNotExist
	}
	user.ChangeUsername(profile.UserName())
	user.ChangeOrganization(profile.OrganizationCode())
	if err := user.ChangeEmail(profile.Email()); err != nil {
		return err
	}
	if err := user.ChangePhone(profile.Phone()); err != nil {
		return err
	}
	return app.repo.SaveOneUser(user)
}

func (app *UserApp) ResetPassword(userID string) error {
	user, err := app.repo.FindUserByUserID(userID)
	if err != nil {
		return common.ErrUserNotExist
	}
	// newPassword := user.ResetPassword()
	if err := app.repo.SaveOneUser(user); err != nil {
		return err
	}
	// app.mailService.SendPasswordResetMail(userID, newPassword)
	return nil
}

func (app *UserApp) CreateOneUser(profile *valobj.UserProfile) error {
	password, err := valobj.NewPassword("initpassword")
	if err != nil {
		return err
	}
	// newPassword := password.ResetPassword()
	email, err := valobj.NewEmail(profile.Email())
	if err != nil {
		return err
	}
	phone, err := valobj.NewPhone(profile.Phone())
	if err != nil {
		return err
	}
	avatar, err := valobj.NewAvatar(profile.Avatar())
	if err != nil {
		return err
	}
	status := true
	user := entity.NewUserEntity(
		profile.UserID(),
		profile.UserName(),
		password,
		email,
		phone,
		profile.OrganizationCode(),
		avatar,
		status,
	)
	if err := app.repo.SaveOneUser(user); err != nil {
		return err
	}
	// app.mailService.SendPasswordResetMail(profile.UserID(), newPassword)
	return nil
}

func (app *UserApp) SyncWxworkUser() error {
	userList := app.wxworkService.WxworkUserList()
	for _, v := range userList {
		if err := app.CreateOneUser(v); err != nil {
			return err
		}
	}
	return nil
}
