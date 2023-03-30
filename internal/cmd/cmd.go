package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"goframe/internal/controller/backend/setting/rabc"
	"goframe/internal/controller/backend/setting/system"
	"goframe/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(groupApi *ghttp.RouterGroup) {
				groupApi.Middleware(service.Middleware().HandlerResponse)
				groupApi.Group("/backend/setting", func(group *ghttp.RouterGroup) {
					group.Bind(
						rabc.NewAdmin(),        // 管理员
						rabc.NewRole(),         // 角色管理
						system.NewBasicSetup(), // 基础设置
						system.NewShopSetup(),  // 商城设置
						system.NewUserSetup(),  // 用户设置
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
