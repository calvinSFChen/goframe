package route

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe/internal/controller"
	"goframe/internal/service"
	"goframe/route/backend"
	"goframe/route/frontend"
)

func InitRoute(ctx context.Context, s *ghttp.Server) {
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Bind(controller.NewIndex())
		group.Middleware(
			service.Middleware().CORS,
			service.Middleware().Ctx,
			service.Middleware().HandlerResponse,
		)
		// 加载后端路由
		backend.InitBackend(ctx, group)

		// 加载前端路由
		frontend.InitFrontend(ctx, group)
	})
}
