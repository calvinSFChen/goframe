package system

import (
	"context"
	"goframe/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

const Module = "frontend/system"

func InitRoute(ctx context.Context, group *ghttp.RouterGroup) {
	service.Route().InitModule(Module, true)
	// setting路由
	group.Group("/system", func(group *ghttp.RouterGroup) {
		//gfToken := service.Middleware().InitGfToken(true)
		//gfToken.Middleware(ctx, group)
		//加载公共路由
		service.Route().InitRoute(group, false)
		group.Middleware(service.Middleware().Auth)
		// 加载路由
		service.Route().InitRoute(group, true)
		// 保存缓存
		//route.SaveCache(ctx, "backend")
	})
}
