package rabc

import (
	"context"
	"goframe/api/v1/backend/setting/rabc"
	"goframe/internal/model/setting"
	"goframe/internal/service"
)

type cMenu struct {
}

var NewMenu = &cMenu{}

// Add 添加菜单
// @summary 添加菜单
// @description 添加菜单
// @tags 设置管理
// @accept application/json
// @produce application/json
// @params body body rabc.MenuAddReq
// @success 200 {object} response.ResponseRes={data=rabc.AdminLoginRes} "code错误码 msg操作信息 data返回信息"
// @route /backend/setting/rabc_menu/add [post]
func (c *cMenu) Add(ctx context.Context, req *rabc.MenuAddReq) (res *rabc.MenuAddRes, err error) {
	err = service.RabcMenu().Add(ctx, req)
	return
}

// Edit 编辑菜单
// @summary 编辑菜单
// @description 编辑菜单
// @tags 设置管理
// @accept application/json
// @produce application/json
// @params body body rabc.MenuEditReq
// @success 200 {object} response.ResponseRes={data=rabc.MenuEditRes} "code错误码 msg操作信息 data返回信息"
// @route /backend/setting/rabc_menu/edit [post]
func (c *cMenu) Edit(ctx context.Context, req *rabc.MenuEditReq) (res *rabc.MenuEditRes, err error) {
	err = service.RabcMenu().Edit(ctx, req)
	return
}

// List 菜单列表
// @summary 菜单列表
// @description 菜单列表
// @tags 设置管理
// @accept application/json
// @produce application/json
// @params body body rabc.MenuListReq
// @success 200 {object} response.ResponseRes={data=rabc.MenuListRes} "code错误码 msg操作信息 data返回信息"
// @route /backend/setting/rabc_menu/list [get]
func (c *cMenu) List(ctx context.Context, req *rabc.MenuListReq) (res *rabc.MenuListRes, err error) {
	var (
		out []setting.MenuListItem
	)
	out, err = service.RabcMenu().List(ctx, req)
	if err != nil {
		return
	}
	res = &rabc.MenuListRes{
		List: out,
	}
	return
}
