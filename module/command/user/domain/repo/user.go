package repo

import "gin-ddd-demo/module/command/user/domain/entity"

type UserRepo interface {
	SaveOneUser(user *entity.UserEntity) error
	FindUserByUserID(userID string) (*entity.UserEntity, error)
	RemoveOneUser(user *entity.UserEntity) error
}
