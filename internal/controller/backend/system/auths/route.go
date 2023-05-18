package auths

import (
	"context"
	"goframe/api/v1/backend/system/auths"
	"goframe/internal/service"
)

type cRoute struct {
}

var NewRoute = &cRoute{}

// Add 添加路由
// @summary 菜单列表
// @description 菜单列表
// @tags 设置管理
// @accept application/json
// @produce application/json
// @params body body auth.RouteAddReq
// @success 200 {object} response.ResponseRes={data=auth.RouteAddRes} "code错误码 msg操作信息 data返回信息"
// @route /backend/system/auths_route/Add [post]
func (c *cRoute) Add(ctx context.Context, req *auths.RouteAddReq) (res *auths.RouteAddRes, err error) {
	err = service.AuthsRoute().Add(ctx, req)
	return
}

// Edit 编辑路由
// @summary 编辑路由
// @description 编辑路由
// @tags 设置管理
// @accept application/json
// @produce application/json
// @params body body auths.RouteEditReq
// @success 200 {object} response.ResponseRes={data=auth.RouteEditRes} "code错误码 msg操作信息 data返回信息"
// @route /backend/system/auths_route/edit [post]
func (c *cRoute) Edit(ctx context.Context, req *auths.RouteEditReq) (res *auths.RouteEditRes, err error) {
	err = service.AuthsRoute().Edit(ctx, req)
	return
}

// List 路由列表
// @summary 路由列表
// @description 路由列表
// @tags 设置管理
// @accept application/json
// @produce application/json
// @params body body auths.RouteListReq
// @success 200 {object} response.ResponseRes={data=auth.RouteListRes} "code错误码 msg操作信息 data返回信息"
// @route /backend/system/auths_route/list [get]
func (c *cRoute) List(ctx context.Context, req *auths.RouteListReq) (res auths.RouteListRes, err error) {
	res.Total, res.List, err = service.AuthsRoute().List(ctx, req)
	return
}
