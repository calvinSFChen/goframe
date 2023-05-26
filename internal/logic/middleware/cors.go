package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func (s *sMiddleware) CORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	//corsOptions.AllowDomain = []string{"goframe.org", "127.0.0.1"}
	corsOptions.AllowOrigin = "*"
	corsOptions.AllowMethods = "POST, GET, OPTIONS, PUT, DELETE,UPDATE"
	corsOptions.AllowCredentials = "false"
	corsOptions.MaxAge = 1728000
	corsOptions.AllowHeaders = "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma"
	corsOptions.ExposeHeaders = "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar"
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}
