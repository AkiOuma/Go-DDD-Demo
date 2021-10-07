package main

import (
	"gin-ddd-demo/infra/inner"
	"gin-ddd-demo/infra/persistence"
	user_inner "gin-ddd-demo/module/command/user/adapter/inner"
	"gin-ddd-demo/module/command/user/adapter/persistence/sqldb/dao"
	user_app "gin-ddd-demo/module/command/user/application"
	wxwork_inner "gin-ddd-demo/module/command/wxwork/adapter/inner"
	"gin-ddd-demo/module/command/wxwork/adapter/persistence/kv"
	wxwork_app "gin-ddd-demo/module/command/wxwork/application"
)

func main() {
	// init infra component
	db := persistence.NewSqlDB()
	container := inner.NewContainer()

	// init repositories
	userDAO := dao.NewUserDAO(db)
	contactDAO := kv.NewContactDAO()

	// init modules
	wxworkApp := wxwork_app.NewWxworkApp(contactDAO)
	wxworkInnerService := wxwork_inner.NewWxworkInternalService(wxworkApp)
	container.SetWxworkInnerService(wxworkInnerService)

	userApp := user_app.NewUserApp(
		userDAO,
		user_inner.NewWxworkClient(container.GetWxworkInnerService()),
	)
	userApp.SyncWxworkUser()
}
