package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"goframe/route"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 加载路由
			route.InitRoute(ctx, s)
			//s.Plugin(&swagger.Swagger{})
			s.Run()
			return nil
		},
	}
)
