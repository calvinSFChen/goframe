package setting

import (
	"goframe/internal/controller/backend/setting/rabc"
	"goframe/internal/model/setting"
	"goframe/internal/service"
)

func init() {
	service.Route().AddRoute(
		// admin 管理员
		setting.RouteItem{Module: Module, IsAuth: true, Func: rabc.NewAdmin.Add},
		setting.RouteItem{Module: Module, IsAuth: true, Func: rabc.NewAdmin.Edit},
		setting.RouteItem{Module: Module, IsAuth: true, Func: rabc.NewAdmin.List},
		setting.RouteItem{Module: Module, IsAuth: false, Func: rabc.NewAdmin.Login},
		setting.RouteItem{Module: Module, IsAuth: false, Func: rabc.NewAdmin.Logout},

		// menu 菜单管理
		setting.RouteItem{Module: Module, IsAuth: true, Func: rabc.NewMenu.Add},
		setting.RouteItem{Module: Module, IsAuth: true, Func: rabc.NewMenu.Edit},
		setting.RouteItem{Module: Module, IsAuth: true, Func: rabc.NewMenu.List},

		// route 路由管理
		setting.RouteItem{Module: Module, IsAuth: true, Func: rabc.NewRoute.Add},
		setting.RouteItem{Module: Module, IsAuth: true, Func: rabc.NewRoute.Edit},
		setting.RouteItem{Module: Module, IsAuth: true, Func: rabc.NewRoute.List},
	)
}
