package backend

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe/route/backend/setting"
)

func InitBackend(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/backend", func(group *ghttp.RouterGroup) {
		setting.InitRoute(ctx, group)
	})
}
