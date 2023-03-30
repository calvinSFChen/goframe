package rabc

import "github.com/gogf/gf/v2/frame/g"

type RoleIndexReq struct {
	g.Meta `path:"rabc/role/index"  tags:"管理员管理" method:"get" description:""`
	Name   string `json:"name" summary:"测试"`
	Age    int    `json:"age" summary:"cwes"`
}

type RoleIndexRes struct {
	g.Meta `mime:"application/json" tags:"管理员管理" example:"string"`
	Test   string `json:"test"`
}
