package system

import (
	"goframe/internal/service"
)

func init() {
	// service.RegisterSystemConfig(NewSystemConfig()) // 系统配置
	service.RegisterAuthsAdmin(NewAuthsAdmin()) // 管理员
	service.RegisterAuthsMenu(NewAuthsMenu())   // 菜单
	service.RegisterAuthsRoute(NewAuthsRoute()) // 路由
}
