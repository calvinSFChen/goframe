package frontend

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe/route/frontend/setting"
)

func InitFrontend(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/frontend", func(group *ghttp.RouterGroup) {
		setting.InitRoute(ctx, group)
	})
}
