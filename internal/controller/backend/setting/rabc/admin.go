package rabc

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"goframe/api/v1/backend/setting/rabc"
	"goframe/internal/model/entity"
	"goframe/internal/service"
	"goframe/utility/cerrors"
)

type cAdmin struct {
}

var NewAdmin = &cAdmin{}

// List 管理员列表
// @summary 管理员列表
// @description 管理员列表
// @tags 设置管理员
// @accept application/json
// @produce application/json
// @params body body system.AdminListReq
// @success 200 {object} response.ResponseRes={data=system.AdminListRes} "code错误码 msg操作信息 data返回信息"
// @route /backend/setting/system_rabc/list [get]
func (c *cAdmin) List(ctx context.Context, req *rabc.AdminListReq) (res *rabc.AdminListRes, err error) {
	var (
		total int
		out   []entity.MuRabcAdmin
	)
	total, out, err = service.RabcAdmin().List(ctx, req)
	res = &rabc.AdminListRes{
		Total: total,
		List:  out,
	}
	return
}

// Add 添加管理员
// @summary 添加管理员
// @description 添加管理员
// @tags 设置管理员
// @accept application/json
// @produce application/json
// @params body body system.AdminAddReq
// @success 200 {object} response.ResponseRes={data=system.AdminAddRes} "code错误码 msg操作信息 data返回信息"
// @route /backend/setting/rabc_admin/add [post]
func (c *cAdmin) Add(ctx context.Context, req *rabc.AdminAddReq) (res *rabc.AdminAddRes, err error) {
	err = service.RabcAdmin().Add(ctx, req)
	return
}

// Edit 编辑管理员
// @summary 编辑管理员
// @description 编辑管理员
// @tags 设置管理
// @accept application/json
// @produce application/json
// @params body body system.AdminEditReq
// @success 200 {object} response.ResponseRes={data=system.AdminEditRes} "code错误码 msg操作信息 data返回信息"
// @route /backend/setting/rabc_admin/edit [post]
func (c *cAdmin) Edit(ctx context.Context, req *rabc.AdminEditReq) (res *rabc.AdminEditRes, err error) {
	err = service.RabcAdmin().Edit(ctx, req)
	return
}

// Login 登录
// @summary 管理员登录
// @description 管理员登录
// @tags 设置管理
// @accept application/json
// @produce application/json
// @params body body rabc.AdminLoginReq
// @success 200 {object} response.ResponseRes={data=rabc.AdminLoginRes} "code错误码 msg操作信息 data返回信息"
// @route /backend/setting/rabc_admin/login [post]
func (c *cAdmin) Login(ctx context.Context, req *rabc.AdminLoginReq) (res *rabc.AdminLoginRes, err error) {
	res, err = service.RabcAdmin().Login(ctx, req)
	if err != nil {
		return
	}

	listReq := &rabc.MenuListReq{
		IsShow: 1,
	}

	res.Menus, err = service.RabcMenu().List(ctx, listReq)
	if err != nil {
		return
	}

	for _, v := range res.Menus {
		res.UniqueAuth = append(res.UniqueAuth, v.UniqueAuth)
	}
	res.UniqueAuth, err = service.RabcMenu().GetAllUniquePath(ctx)
	err = gerror.NewCode(cerrors.CodeSuccessNil, "登录成功")

	return
}

// Logout 退出登录
// @summary 退出登录
// @description 退出登录
// @tags 设置管理
// @accept application/json
// @produce application/json
// @params body body rabc.AdminLogoutReq
// @success 200 {object} response.ResponseRes={data=rabc.AdminLogoutReq} "code错误码 msg操作信息 data返回信息"
// @route /backend/setting/rabc_admin/logout [get]
func (c *cAdmin) Logout(ctx context.Context, req *rabc.AdminLogoutReq) (res *rabc.AdminLogoutRes, err error) {
	service.RabcAdmin().Logout(ctx)
	err = gerror.NewCode(cerrors.CodeExpire, "退出成功")
	return
}
