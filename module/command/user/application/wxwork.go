package application

import "gin-ddd-demo/module/command/user/domain/valobj"

type Wxwork interface {
	WxworkUserList() []*valobj.UserProfile
}
