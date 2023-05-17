package system

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe/internal/service"
)

const Module = "backend/setting"

func InitRoute(ctx context.Context, group *ghttp.RouterGroup) {
	service.Route().InitModule(Module, true)
	// setting路由
	group.Group("/setting", func(group *ghttp.RouterGroup) {
		// 加载公共路由
		service.Route().Middleware(group, false, false)
		group.Middleware(service.Middleware().Auth)
		//加载路由
		service.Route().Middleware(group, true, true)
	})
}
