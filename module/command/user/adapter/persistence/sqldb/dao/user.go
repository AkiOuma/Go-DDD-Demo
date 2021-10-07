package dao

import (
	"gin-ddd-demo/module/command/user/adapter/persistence/sqldb/model"
	"gin-ddd-demo/module/command/user/domain/entity"
	"gin-ddd-demo/module/command/user/domain/repo"
	"gin-ddd-demo/module/command/user/domain/valobj"

	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

var _ repo.UserRepo = &UserDAO{}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) SaveOneUser(user *entity.UserEntity) error {
	row := dao.entity2Model(user)
	if _, err := dao.FindUserByUserID(user.UserID()); err != nil {
		// user does not exist, insert it
		return dao.db.Create(row).Error
	}
	// user exists, update it
	return dao.db.Where(&model.UserModel{
		UserID: user.UserID(),
	}).Updates(row).Error
}

func (dao *UserDAO) FindUserByUserID(userID string) (*entity.UserEntity, error) {
	var row model.UserModel
	if err := dao.db.First(
		&row,
		&model.UserModel{UserID: userID},
	).Error; err != nil {
		return nil, err
	}
	return dao.model2Entity(&row)
}

func (dao *UserDAO) RemoveOneUser(user *entity.UserEntity) error {
	return dao.db.Delete(
		&model.UserModel{},
		&model.UserModel{UserID: user.UserID()},
	).Error
}

func (UserDAO) entity2Model(entity *entity.UserEntity) *model.UserModel {
	return &model.UserModel{
		UserID:           entity.UserID(),
		UserName:         entity.UserName(),
		Password:         entity.Password(),
		Email:            entity.Email(),
		Phone:            entity.Phone(),
		OrganizationCode: entity.OrganizationCode(),
		Avatar:           entity.Avatar(),
	}
}

func (UserDAO) model2Entity(row *model.UserModel) (*entity.UserEntity, error) {
	password, err := valobj.NewPassword(row.Password)
	if err != nil {
		return nil, err
	}
	email, err := valobj.NewEmail(row.Email)
	if err != nil {
		return nil, err
	}
	phone, err := valobj.NewPhone(row.Phone)
	if err != nil {
		return nil, err
	}
	avatar, err := valobj.NewAvatar(row.Avatar)
	if err != nil {
		return nil, err
	}

	return entity.NewUserEntity(
		row.UserID,
		row.UserName,
		password,
		email,
		phone,
		row.OrganizationCode,
		avatar,
		row.Status,
	), nil
}
