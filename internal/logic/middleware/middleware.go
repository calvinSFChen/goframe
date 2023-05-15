package middleware

import (
	"goframe/internal/service"
)

type sMiddleware struct {
}

var Middleware = sMiddleware{}

func init() {
	service.RegisterMiddleware(New())
}
