package system

import (
	"goframe/internal/controller/backend/system/auths"
	"goframe/internal/model/system"
	"goframe/internal/service"
)

func init() {
	service.Route().AddRoute(
		// admin 管理员
	// 	setting.RouteItem{Module: Module, IsAuth: true, Func: auth.NewAdmin.Add},
	// 	setting.RouteItem{Module: Module, IsAuth: true, Func: auth.NewAdmin.Edit},
	// 	setting.RouteItem{Module: Module, IsAuth: true, Func: auth.NewAdmin.List},
	system.RouteItem{Module: Module, IsAuth: false, Func: auths.NewAdmin.Login},
	system.RouteItem{Module: Module, IsAuth: false, Func: auths.NewAdmin.Logout},

	// 	// menu 菜单管理
	system.RouteItem{Module: Module, IsAuth: true, Func: auths.NewMenu.Add},
		// setting.RouteItem{Module: Module, IsAuth: true, Func: auth.NewMenu.Edit},
		// setting.RouteItem{Module: Module, IsAuth: true, Func: auth.NewMenu.List},

	// 	// route 路由管理
	// 	setting.RouteItem{Module: Module, IsAuth: true, Func: auth.NewRoute.Add},
	// 	setting.RouteItem{Module: Module, IsAuth: true, Func: auth.NewRoute.Edit},
	// 	setting.RouteItem{Module: Module, IsAuth: true, Func: auth.NewRoute.List},
	)
}
