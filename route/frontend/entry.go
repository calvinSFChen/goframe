package frontend

import (
	"context"
	"fmt"
	"goframe/route/frontend/system"

	"github.com/gogf/gf/v2/net/ghttp"
)

func InitFrontend(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/frontend", func(group *ghttp.RouterGroup) {
		system.InitRoute(ctx, group)
	})
}
