package system

// import (
// 	"context"
// 	"github.com/gogf/gf/v2/errors/gerror"
// 	"github.com/gogf/gf/v2/text/gstr"
// 	"goframe/api/v1/backend/setting/rabc"
// 	"goframe/internal/dao"
// 	"goframe/internal/model/entity"
// 	"goframe/utility/utils"
// )

type sAuthsRoute struct {
}

func NewAuthsRoute() *sAuthsRoute {
	return &sAuthsRoute{}
}

// // List 管理员列表
// func (s *sRabcRoute) List(ctx context.Context, input *rabc.RouteListReq) (total int, out []entity.MuRabcRoute, err error) {
// 	var (
// 		page     = input.Page
// 		pageSize = input.PageSize
// 		name     = input.Name
// 		status   = input.Status
// 	)

// 	model := dao.MuRabcRoute.Ctx(ctx).Safe(false)
// 	if name != "" {
// 		model.Where("name=?", name)
// 	}
// 	if status != -1 {
// 		model.Where("status=?", status)
// 	}

// 	if total, err = model.Count(); err != nil {
// 		return
// 	}
// 	model.Page(page, (page-1)*pageSize)
// 	if err = model.Scan(&out); err != nil {
// 		return
// 	}
// 	return
// }

// // Add 添加
// func (s *sRabcRoute) Add(ctx context.Context, req *rabc.RouteAddReq) (err error) {
// 	var (
// 		menuId = req.MenuId
// 		method = gstr.Trim(req.Method)
// 		name   = gstr.Trim(req.Name)
// 		path   = gstr.Trim(req.Path)
// 		sort   = req.Sort
// 		status = req.Status

// 		rabcRoute entity.MuRabcRoute
// 	)

// 	var exitMenuId int
// 	exitMenuId, err = dao.MuRabcMenus.Ctx(ctx).Where(dao.MuRabcMenus.Columns().Id, menuId).Count()
// 	if err != nil {
// 		return
// 	}
// 	if exitMenuId <= 0 {
// 		err = gerror.New(utils.T(ctx, "菜单分类不存在"))
// 		return
// 	}
// 	rabcRoute.Name = name
// 	rabcRoute.MenuId = uint(menuId)
// 	rabcRoute.Method = method
// 	rabcRoute.Path = path
// 	rabcRoute.Sort = int(sort)
// 	rabcRoute.Status = int(status)
// 	var lastId int64
// 	lastId, err = dao.MuRabcRoute.Ctx(ctx).Data(rabcRoute).InsertAndGetId()
// 	if err != nil {
// 		return
// 	}
// 	if lastId <= 0 {
// 		err = gerror.New(utils.T(ctx, "操作失败"))
// 		return
// 	}
// 	return
// }

// // Edit 编辑管理员
// func (s *sRabcRoute) Edit(ctx context.Context, req *rabc.RouteEditReq) (err error) {
// 	var (
// 		id     = req.Id
// 		menuId = req.MenuId
// 		method = gstr.Trim(req.Method)
// 		name   = gstr.Trim(req.Name)
// 		path   = gstr.Trim(req.Path)
// 		sort   = req.Sort
// 		status = req.Status

// 		rabcRoute entity.MuRabcRoute
// 	)

// 	if id <= 0 {
// 		err = gerror.New(utils.T(ctx, "参数异常"))
// 		return
// 	}
// 	err = dao.MuRabcRoute.Ctx(ctx).Where("id=?", id).Scan(&rabcRoute)
// 	if err != nil {
// 		return
// 	}
// 	rabcRoute.Name = name
// 	rabcRoute.MenuId = uint(menuId)
// 	rabcRoute.Method = method
// 	rabcRoute.Path = path
// 	rabcRoute.Sort = int(sort)
// 	rabcRoute.Status = int(status)
// 	_, err = dao.MuRabcRoute.Ctx(ctx).Data(rabcRoute).Where("id=?", id).Update()
// 	if err != nil {
// 		return
// 	}
// 	return
// }
