package inner

import (
	"gin-ddd-demo/infra/inner/interfaces"
	"gin-ddd-demo/module/command/user/application"
	"gin-ddd-demo/module/command/user/domain/valobj"
	"gin-ddd-demo/types"
)

type WxworkClient struct {
	wxworkService interfaces.WxworkInternalService
}

func NewWxworkClient(wxworkService interfaces.WxworkInternalService) *WxworkClient {
	return &WxworkClient{
		wxworkService: wxworkService,
	}
}

var _ application.Wxwork = &WxworkClient{}

func (client *WxworkClient) WxworkUserList() []*valobj.UserProfile {
	users, err := client.wxworkService.GetWxUserList()
	if err != nil {
		return nil
	}
	userList := make([]*valobj.UserProfile, len(users))
	for k, v := range users {
		userList[k] = client.userProfileTransfer(v)
	}
	return userList
}

func (WxworkClient) userProfileTransfer(detail *types.UserDetail) *valobj.UserProfile {
	profile := &valobj.UserProfile{}
	profile.SetUserID(detail.UserID)
	profile.SetUserName(detail.Name)
	profile.SetEmail(detail.Email)
	profile.SetAvatar(detail.Avatar)
	profile.SetPhone(detail.Phone)
	return profile
}
