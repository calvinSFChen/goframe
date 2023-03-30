package rabc

import (
	"context"
	"fmt"
	"goframe/api/v1/backend/setting/rabc"
)

type cAdmin struct{}

func NewAdmin() *cAdmin {
	return &cAdmin{}
}

func (c *cAdmin) Index(ctx context.Context, req *rabc.AdminIndexReq) (res *rabc.AdminIndexRes, err error) {
	//g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	fmt.Println(2323)
	res = &rabc.AdminIndexRes{
		Name: "测试",
		Age:  10,
	}
	return
}
