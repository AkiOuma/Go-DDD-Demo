package kv

import (
	"gin-ddd-demo/module/command/wxwork/domain/entity"
	"gin-ddd-demo/module/command/wxwork/domain/repo"
)

type ContactDAO struct{}

func NewContactDAO() *ContactDAO {
	return &ContactDAO{}
}

var _ repo.ContactRepo = &ContactDAO{}

func (dao *ContactDAO) NewContact() *entity.ContactEntity {
	return entity.NewContactEntity("token")
}
