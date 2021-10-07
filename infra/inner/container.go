package inner

import "gin-ddd-demo/infra/inner/interfaces"

type Container struct {
	wxwork interfaces.WxworkInternalService
}

func NewContainer() *Container {
	return &Container{}
}

func (container *Container) SetWxworkInnerService(wxwork interfaces.WxworkInternalService) {
	container.wxwork = wxwork
}

func (container *Container) GetWxworkInnerService() interfaces.WxworkInternalService {
	return container.wxwork
}
