package system

import (
	"goframe/internal/controller/backend/system/config"
	"goframe/internal/model/system"
	"goframe/internal/service"
)

func init() {
	service.Route().AddRoute(
		// config 设置管理
		system.RouteItem{Module: Module, IsAuth: true, Func: config.NewConfig.Add},
		system.RouteItem{Module: Module, IsAuth: true, Func: config.NewConfig.Edit},
		// setting.RouteItem{Module: Module, IsAuth: true, Func: system.NewConfig.List},
		system.RouteItem{Module: Module, IsAuth: true, Func: config.NewConfig.GetOne},
	)
}
