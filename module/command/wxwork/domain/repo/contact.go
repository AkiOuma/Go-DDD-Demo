package repo

import "gin-ddd-demo/module/command/wxwork/domain/entity"

type ContactRepo interface {
	NewContact() *entity.ContactEntity
}
