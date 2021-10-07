package interfaces

import "gin-ddd-demo/types"

type WxworkInternalService interface {
	GetWxUserList() ([]*types.UserDetail, error)
}
