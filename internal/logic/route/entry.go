package route

import "goframe/internal/service"

func init() {
	service.RegisterRoute(NewRoute()) // 系统配置
}
