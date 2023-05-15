package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe/internal/logic/auth"
)

func (s *sMiddleware) Auth(r *ghttp.Request) {
	// GfJWTMiddleware gf jwt集成的中间件
	// Auth是权限service中配置的gf jwt
	auth.AuthsService.AuthJwt.MiddlewareFunc()(r)
	r.Middleware.Next()
}
