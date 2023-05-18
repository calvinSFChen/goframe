package auths

import (
	"context"
	"goframe/api/v1/backend/system/auths"

	"github.com/gogf/gf/v2/errors/gerror"

	// "goframe/internal/model/entity"

	"goframe/internal/model/system"
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
// @route /backend/setting/system_auth/list [get]
func (c *cAdmin) List(ctx context.Context, req *auths.AdminListReq) (res *auths.AdminListRes, err error) {
	var (
		total int
		out   []system.SystemAdminListOut
	)
	total, out, err = service.AuthsAdmin().List(ctx, req)
	res = &auths.AdminListRes{
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
// @route /backend/system/auths_admin/add [post]
func (c *cAdmin) Add(ctx context.Context, req *auths.AdminAddReq) (res *auths.AdminAddRes, err error) {
	err = service.AuthsAdmin().Add(ctx, req)
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
// @route /backend/system/auths_admin/edit [post]
func (c *cAdmin) Edit(ctx context.Context, req *auths.AdminEditReq) (res *auths.AdminEditRes, err error) {
	err = service.AuthsAdmin().Edit(ctx, req)
	return
}

// Login 登录
// @summary 管理员登录
// @description 管理员登录
// @tags 设置管理
// @accept application/json
// @produce application/json
// @params body body auths.AdminLoginReq
// @success 200 {object} response.ResponseRes={data=auths.AdminLoginRes} "code错误码 msg操作信息 data返回信息"
// @route /backend/setting/auths_admin/login [post]
func (c *cAdmin) Login(ctx context.Context, req *auths.AdminLoginReq) (res *auths.AdminLoginRes, err error) {
	res, err = service.AuthsAdmin().Login(ctx, req)
	if err != nil {
		return
	}

	// listReq := &auths.MenuListReq{
	// 	IsShow: 1,
	// }

	// res.Menus, err = service.authsMenu().List(ctx, listReq)
	// if err != nil {
	// 	return
	// }

	// for _, v := range res.Menus {
	// 	res.Uniqueauths = append(res.Uniqueauths, v.UniqueAuth)
	// }
	// res.UniqueAuth, err = service.authsMenu().GetAllUniquePath(ctx)
	err = gerror.NewCode(cerrors.CodeSuccessNil, "登录成功")

	return
}

// Logout 退出登录
// @summary 退出登录
// @description 退出登录
// @tags 设置管理
// @accept application/json
// @produce application/json
// @params body body auth.AdminLogoutReq
// @success 200 {object} response.ResponseRes={data=auth.AdminLogoutReq} "code错误码 msg操作信息 data返回信息"
// @route /backend/setting/auth_admin/logout [get]
func (c *cAdmin) Logout(ctx context.Context, req *auths.AdminLogoutReq) (res *auths.AdminLogoutRes, err error) {
	service.AuthsAdmin().Logout(ctx)
	err = gerror.NewCode(cerrors.CodeExpire, "退出成功")
	return
}
