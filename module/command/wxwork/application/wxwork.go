package application

import (
	"gin-ddd-demo/module/command/wxwork/domain/repo"
	"gin-ddd-demo/module/command/wxwork/domain/valobj"
)

type WxworkApp struct {
	repo repo.ContactRepo
}

func NewWxworkApp(repo repo.ContactRepo) *WxworkApp {
	return &WxworkApp{
		repo: repo,
	}
}

func (app *WxworkApp) GetWxUserList() ([]*valobj.UserDetail, error) {
	contact := app.repo.NewContact()
	return contact.GetWxUserList()
}
