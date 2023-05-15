package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe/internal/model"
	"goframe/internal/service"
)

// Ctx 自定义上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: r.Session,
		Data:    make(g.Map),
	}
	// 初始化自定义上下文
	service.BizCtx().Init(r, customCtx)
	// 执行下一步请求逻辑
	r.Middleware.Next()
}
