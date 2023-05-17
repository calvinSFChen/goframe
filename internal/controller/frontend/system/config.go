package system

import (
	"context"
	"fmt"
	"goframe/api/v1/frontend/system"
)

type cConfig struct {
}

var NewConfig = &cConfig{}

// List 系统配置列表
// @summary 系统配置列表
// @description 系统配置列表
// @tags 设置管理
// @accept application/json
// @produce application/json
// @param data query system.BasicConfigListReq
// @success 200 {object} response.ResponseRes={data=system.BasicConfigListRes} "code操作码 msg操作信息 data返回信息"
// @route /frontend/setting/system_config/list [get]
func (c *cConfig) List(ctx context.Context, req *system.BasicConfigListReq) (res *system.BasicConfigListRes, err error) {
	fmt.Println("this is frontend-list")
	return
}
