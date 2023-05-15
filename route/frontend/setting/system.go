package setting

import (
	"goframe/internal/controller/frontend/setting/system"
	"goframe/internal/model/setting"
	"goframe/internal/service"
)

func init() {

	service.Route().AddRoute(
		setting.RouteItem{Module: Module, IsAuth: false, Func: system.NewConfig.List},
	)
}
