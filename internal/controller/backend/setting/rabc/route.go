package rabc

import (
	"context"
	"goframe/api/v1/backend/setting/rabc"
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
// @params body body rabc.RouteAddReq
// @success 200 {object} response.ResponseRes={data=rabc.RouteAddRes} "code错误码 msg操作信息 data返回信息"
// @route /backend/setting/rabc_route/Add [post]
func (c *cRoute) Add(ctx context.Context, req *rabc.RouteAddReq) (res *rabc.RouteAddRes, err error) {
	err = service.RabcRoute().Add(ctx, req)
	return
}

// Edit 编辑路由
// @summary 编辑路由
// @description 编辑路由
// @tags 设置管理
// @accept application/json
// @produce application/json
// @params body body rabc.RouteEditReq
// @success 200 {object} response.ResponseRes={data=rabc.RouteEditRes} "code错误码 msg操作信息 data返回信息"
// @route /backend/setting/rabc_route/edit [post]
func (c *cRoute) Edit(ctx context.Context, req *rabc.RouteEditReq) (res *rabc.RouteEditRes, err error) {
	err = service.RabcRoute().Edit(ctx, req)
	return
}

// List 路由列表
// @summary 路由列表
// @description 路由列表
// @tags 设置管理
// @accept application/json
// @produce application/json
// @params body body rabc.RouteListReq
// @success 200 {object} response.ResponseRes={data=rabc.RouteListRes} "code错误码 msg操作信息 data返回信息"
// @route /backend/setting/rabc_route/list [get]
func (c *cRoute) List(ctx context.Context, req *rabc.RouteListReq) (res rabc.RouteListRes, err error) {
	res.Total, res.List, err = service.RabcRoute().List(ctx, req)
	return
}
