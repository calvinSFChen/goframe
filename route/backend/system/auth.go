package system

import (
	"goframe/internal/controller/backend/system/auths"
	"goframe/internal/model/system"
	"goframe/internal/service"
)

func init() {
	service.Route().AddRoute(
		// admin 管理员
		system.RouteItem{Module: Module, IsAuth: true, Func: auths.NewAdmin.Add},
		system.RouteItem{Module: Module, IsAuth: true, Func: auths.NewAdmin.Edit},
		system.RouteItem{Module: Module, IsAuth: true, Func: auths.NewAdmin.List},
		system.RouteItem{Module: Module, IsAuth: false, Func: auths.NewAdmin.Login},
		system.RouteItem{Module: Module, IsAuth: false, Func: auths.NewAdmin.Logout},

		// 	// menu 菜单管理
		system.RouteItem{Module: Module, IsAuth: true, Func: auths.NewMenu.Add},
		system.RouteItem{Module: Module, IsAuth: true, Func: auths.NewMenu.Edit},
		system.RouteItem{Module: Module, IsAuth: true, Func: auths.NewMenu.TreeList},
		system.RouteItem{Module: Module, IsAuth: true, Func: auths.NewMenu.UniqueAuthList},
		system.RouteItem{Module: Module, IsAuth: true, Func: auths.NewMenu.List},

		// route 路由管理
		system.RouteItem{Module: Module, IsAuth: true, Func: auths.NewRoute.Add},
		system.RouteItem{Module: Module, IsAuth: true, Func: auths.NewRoute.Edit},
		system.RouteItem{Module: Module, IsAuth: true, Func: auths.NewRoute.List},
	)
}
