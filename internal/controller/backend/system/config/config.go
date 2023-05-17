package system

// import (
// 	"context"
// 	"goframe/api/v1/backend/setting/system"
// 	"goframe/internal/model/setting"
// 	"goframe/internal/service"
// )

// type cConfig struct {
// }

// var NewConfig = &cConfig{}

// // Add 添加系统配置管理
// // @Summary 添加系统配置管理
// // @Description 添加系统配置管理
// // @Tags 设置管理
// // @Accept application/json
// // @Produce application/json
// // @Param body body system.ConfigAddReq
// // @Success 200 {object} response.ResponseRes={data=system.ConfigAddRes} "code错误码，msg操作信息，data返回信息"
// // @route /backend/setting/system_config/add [post]
// func (c *cConfig) Add(ctx context.Context, req *system.BasicConfigAddReq) (res *system.BasicConfigAddRes, err error) {
// 	err = service.SystemConfig().Add(ctx, req)
// 	return
// }

// // List 系统配置列表
// // @summary 系统配置列表
// // @description 系统配置列表
// // @tags 设置管理
// // @accept application/json
// // @produce application/json
// // @param data query system.BasicConfigListReq
// // @success 200 {object} response.ResponseRes={data=system.BasicConfigListRes} "code操作码 msg操作信息 data返回信息"
// // @route /backend/setting/system_config/list [get]
// func (c *cConfig) List(ctx context.Context, req *system.BasicConfigListReq) (res *system.BasicConfigListRes, err error) {
// 	var (
// 		total int
// 		out   []setting.SystemConfigRes
// 	)
// 	total, out, err = service.SystemConfig().List(ctx, req)
// 	res = &system.BasicConfigListRes{
// 		Total: total,
// 		List:  out,
// 	}
// 	return
// }

// // Edit 编辑系统配置
// // @summary 编辑系统配置
// // @description 编辑系统配置
// // @tags 设置管理
// // @accept application/json
// // @produce application/json
// // @param data query system.BasicConfigEditReq
// // @success 200 {object} response.ResponseRes={data=system.BasicConfigEditRes} "code操作码 msg操作信息 data返回信息"
// // @route /backend/setting/system_config/list [get]
// func (c *cConfig) Edit(ctx context.Context, req *system.BasicConfigEditReq) (res *system.BasicConfigEditRes, err error) {
// 	err = service.SystemConfig().Edit(ctx, req)
// 	return
// }

// func (c *cConfig) GetOne(ctx context.Context, req *system.SystemConfigOenReq) (res *system.SystemConfigOneRes, err error) {
// 	var (
// 		id = req.Id
// 	)
// 	res, err = service.SystemConfig().GetOne(ctx, id)
// 	if err != nil {
// 		return
// 	}
// 	return
// }
