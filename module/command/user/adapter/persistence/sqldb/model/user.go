package model

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID               int            `gorm:"column:id"`
	UserID           string         `gorm:"column:userid"`
	UserName         string         `gorm:"column:username"`
	Password         string         `gorm:"column:password"`
	Email            string         `gorm:"column:email"`
	Phone            string         `gorm:"column:phone"`
	OrganizationCode string         `gorm:"column:organization_code"`
	Avatar           string         `gorm:"column:avatar"`
	Status           bool           `gorm:"column:status"`
	CreatedAt        time.Time      `gorm:"column:created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (UserModel) TableName() string {
	return "user"
}
