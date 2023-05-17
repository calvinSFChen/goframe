package backend

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe/route/backend/system"
)

func InitBackend(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/backend", func(group *ghttp.RouterGroup) {
		system.InitRoute(ctx, group)
	})
}
