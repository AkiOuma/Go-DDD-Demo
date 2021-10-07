package inner

import (
	"gin-ddd-demo/module/command/wxwork/application"
	"gin-ddd-demo/module/command/wxwork/domain/valobj"
	"gin-ddd-demo/types"
)

type WxworkInternalService struct {
	app *application.WxworkApp
}

func NewWxworkInternalService(app *application.WxworkApp) *WxworkInternalService {
	return &WxworkInternalService{
		app: app,
	}
}

func (inner *WxworkInternalService) GetWxUserList() ([]*types.UserDetail, error) {
	userList, err := inner.app.GetWxUserList()
	if err != nil {
		return nil, err
	}
	users := make([]*types.UserDetail, len(userList))
	for k, v := range userList {
		users[k] = inner.userDetailTransfer(v)
	}
	return users, nil
}

func (WxworkInternalService) userDetailTransfer(origin *valobj.UserDetail) *types.UserDetail {
	return &types.UserDetail{
		UserID: origin.UserID(),
		Name:   origin.Name(),
		Email:  origin.Email(),
		Phone:  origin.Phone(),
		Avatar: origin.Avatar(),
	}
}
