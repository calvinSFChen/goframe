package rabc

import (
	"context"
	"fmt"
	"goframe/api/v1/backend/setting/rabc"
)

type cRole struct{}

func NewRole() *cRole {
	return &cRole{}
}

func (c *cRole) Index(ctx context.Context, req *rabc.RoleIndexReq) (res *rabc.RoleIndexRes, err error) {
	fmt.Println("this is Role Index")
	res = &rabc.RoleIndexRes{
		Test: "测试",
	}
	return
}
