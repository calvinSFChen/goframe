package backend

import (
	"context"
	"goframe/route/backend/system"

	"github.com/gogf/gf/v2/net/ghttp"
)

func InitBackend(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/backend", func(group *ghttp.RouterGroup) {
		system.InitRoute(ctx, group)
	})
}
