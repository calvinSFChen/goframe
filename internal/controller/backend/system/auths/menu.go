package auths

import (
	"context"
	"goframe/api/v1/backend/system/auths"
	"goframe/internal/model/system"
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
// @params body body auth.MenuAddReq
// @success 200 {object} response.ResponseRes={data=auth.AdminLoginRes} "code错误码 msg操作信息 data返回信息"
// @route /backend/setting/auth_menu/add [post]
func (c *cMenu) Add(ctx context.Context, req *auths.MenuAddReq) (res *auths.MenuAddRes, err error) {
	err = service.AuthsMenu().Add(ctx, req)
	return
}

// Edit 编辑菜单
// @summary 编辑菜单
// @description 编辑菜单
// @tags 设置管理
// @accept application/json
// @produce application/json
// @params body body auth.MenuEditReq
// @success 200 {object} response.ResponseRes={data=auth.MenuEditRes} "code错误码 msg操作信息 data返回信息"
// @route /backend/setting/auth_menu/edit [post]
func (c *cMenu) Edit(ctx context.Context, req *auths.MenuEditReq) (res *auths.MenuEditRes, err error) {
	err = service.AuthsMenu().Edit(ctx, req)
	return
}

func (c *cMenu) TreeList(ctx context.Context, req *auths.MenuTreeListReq) (res *auths.MenuTreeListRes, err error) {
	var (
		out []system.MenuTreeListItem
	)
	out, err = service.AuthsMenu().TreeList(ctx, req)
	res = &auths.MenuTreeListRes{
		List: out,
	}
	return
}

// List 菜单列表
// @summary 菜单列表
// @description 菜单列表
// @tags 设置管理
// @accept application/json
// @produce application/json
// @params body body auth.MenuListReq
// @success 200 {object} response.ResponseRes={data=auth.MenuListRes} "code错误码 msg操作信息 data返回信息"
// @route /backend/setting/auth_menu/list [get]
func (c *cMenu) List(ctx context.Context, req *auths.MenuListReq) (res *auths.MenuListRes, err error) {
	var (
		out auths.MenuListRes
	)
	out, err = service.AuthsMenu().List(ctx, req)
	if err != nil {
		return
	}
	res = &auths.MenuListRes{
		List:  out.List,
		Total: out.Total,
	}
	return
}

// UniqueAuthList
// @summary 获取前端标识
// @description 获取前端标识
// @tags 菜单管理
// @accept application/json
// @produce application/json
// @params body body auths.MenuUniqueAuthListReq
// @success 200 {object} response.ResponseRes={data=auths.MenuUniqueAuthListRes} "code错误码 msg操作信息 data返回信息"
// @route /backend/setting/auths_menu/unique_auth_list [get]
func (c *cMenu) UniqueAuthList(ctx context.Context, req *auths.MenuUniqueAuthListReq) (res *auths.MenuUniqueAuthListRes, err error) {
	var (
		out []string
	)
	out, err = service.AuthsMenu().UniqueAuthList(ctx, req)
	if err != nil {
		return
	}
	res = &auths.MenuUniqueAuthListRes{
		List: out,
	}
	return
}
