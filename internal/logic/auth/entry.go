package auth

import "goframe/internal/service"

func init() {
	service.RegisterAuth(NewAuthJwt())
}
