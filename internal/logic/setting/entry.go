package setting

import (
	"goframe/internal/service"
)

func init() {
	service.RegisterSystemConfig(NewSystemConfig()) // 系统配置
	service.RegisterRabcAdmin(NewRabcAdmin())       // 管理员
	service.RegisterRabcMenu(NewRabcMenu())         // 菜单
	service.RegisterRabcRoute(NewRabcRoute())       // 路由
}
