package system

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe/internal/service"
)

const Module = "frontend/system"

func InitRoute(ctx context.Context, group *ghttp.RouterGroup) {
	service.Route().InitModule(Module, true)
	// setting路由
	group.Group("/system", func(group *ghttp.RouterGroup) {
		//gfToken := service.Middleware().InitGfToken(true)
		//gfToken.Middleware(ctx, group)
		//加载公共路由
		service.Route().Middleware(group, false, false)
		group.Middleware(service.Middleware().Auth)
		// 加载路由
		service.Route().Middleware(group, true, true)
		// 保存缓存
		//route.SaveCache(ctx, "backend")
	})
}
