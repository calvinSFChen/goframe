package system

import (
	"context"
	"database/sql"
	"fmt"

	"goframe/api/v1/backend/system/auths"
	"goframe/internal/dao"
	"goframe/internal/model/entity"
	"goframe/internal/model/system"
	"goframe/utility/utils"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
)

type sAuthsMenu struct {
}

func NewAuthsMenu() *sAuthsMenu {
	return &sAuthsMenu{}
}

// List 管理员列表
func (s *sAuthsMenu) List(ctx context.Context, input *auths.MenuListReq) (out auths.MenuListRes, err error) {

	var (
		page   = input.Page
		size   = input.Size
		isShow = input.IsShow
		title  = input.Title

		list  []system.MenuListItem
		total int
	)
	model := dao.SystemMenus.Ctx(ctx).Safe(false)
	if isShow != -1 {
		model.Where("is_show=?", isShow)
	}
	if title != "" {
		model.Where("title  like  ?", fmt.Sprintf("%v%%", title))
	}
	model.Where("pid=?", 0)

	if total, err = model.Count(); err != nil {
		return
	}

	if err = model.Offset((page - 1) * size).Limit(size).Order("sort desc").Scan(&list); err != nil {
		return
	}
	out.List = s.GetMenuItem(ctx, list)
	out.Total = total
	return
}

func (s *sAuthsMenu) GetMenuItem(ctx context.Context, list []system.MenuListItem) (res []system.MenuListItem) {
	for _, v := range list {
		dao.SystemMenus.Ctx(ctx).Where("is_show=?", 1).Where("pid=?", v.Id).Order("sort desc").Scan(&v.Children)
		res = append(res, v)
		if len(v.Children) > 0 {
			s.GetMenuItem(ctx, v.Children)
		}
	}
	return
}

// Add 添加
func (s *sAuthsMenu) Add(ctx context.Context, req *auths.MenuAddReq) (err error) {
	var (
		title      = gstr.Trim(req.Title)
		pid        = req.Pid
		path       = gstr.Trim(req.Path)
		uniqueAuth = gstr.Trim(req.UniqueAuth)
		icon       = gstr.Trim(req.Icon)
		sort       = req.Sort
		isShow     = req.IsShow
		status     = req.Status

		systemMenu entity.SystemMenus
		lastId     int64
	)
	if pid != 0 {
		var pidNumber int
		pidNumber, err = dao.SystemMenus.Ctx(ctx).Where("id=?", pid).Count()
		if err != nil {
			return
		}
		if pidNumber <= 0 {
			err = gerror.New(utils.T(ctx, "父级菜单不存在"))
			return
		}
	}

	if path == "" {
		err = gerror.New(utils.T(ctx, "路径不为空"))
		return
	}
	if icon != "" {
		systemMenu.Icon = icon
	}
	systemMenu.Title = title
	systemMenu.Pid = uint(pid)
	systemMenu.Path = path
	systemMenu.UniqueAuth = uniqueAuth
	systemMenu.Sort = int(sort)
	systemMenu.IsShow = uint(isShow)
	systemMenu.Status = uint(status)
	systemMenu.Operator = utils.GetUserName()
	lastId, err = dao.SystemMenus.Ctx(ctx).Data(systemMenu).InsertAndGetId()
	if err != nil {
		return
	}
	if lastId <= 0 {
		err = gerror.New(utils.T(ctx, "操作失败"))
		return
	}
	return
}

// Edit 编辑管理员
func (s *sAuthsMenu) Edit(ctx context.Context, req *auths.MenuEditReq) (err error) {
	var (
		id         = req.Id
		title      = gstr.Trim(req.Title)
		pid        = req.Pid
		path       = gstr.Trim(req.Path)
		uniqueAuth = gstr.Trim(req.UniqueAuth)
		icon       = gstr.Trim(req.Icon)
		sort       = req.Sort
		isShow     = req.IsShow
		status     = req.Status

		systemMenu entity.SystemMenus
	)

	// 过滤特殊代码
	//if err = ghtml.SpecialCharsMapOrStruct(req); err != nil {
	//	return
	//}
	err = dao.SystemMenus.Ctx(ctx).Where("id=", id).Scan(&systemMenu)
	if err != nil {
		return
	}
	if id <= 0 {
		gerror.New(utils.T(ctx, "参数异常"))
		return
	}
	if pid != 0 {
		var pidNumber int
		pidNumber, err = dao.SystemMenus.Ctx(ctx).Where("id=?", pid).Count()
		if err != nil {
			return
		}
		if pidNumber <= 0 {
			err = gerror.New(utils.T(ctx, "父级菜单不存在"))
			return
		}
	}
	if icon != "" {
		systemMenu.Icon = icon
	}
	systemMenu.Title = title
	systemMenu.Pid = uint(pid)
	systemMenu.Path = path
	systemMenu.UniqueAuth = uniqueAuth
	systemMenu.Sort = int(sort)
	systemMenu.IsShow = uint(isShow)
	systemMenu.Status = uint(status)
	systemMenu.Operator = utils.GetUserName()

	var result sql.Result
	result, err = dao.SystemMenus.Ctx(ctx).Where("id=?", id).Data(systemMenu).Update()
	if err != nil {
		return
	}
	if _, err = result.RowsAffected(); err != nil {
		return
	}
	return
}

// UniqueAuthList 获取unique_authk列表数据
func (s *sAuthsMenu) UniqueAuthList(ctx context.Context, input *auths.MenuUniqueAuthListReq) (out []string, err error) {
	var (
		list []entity.SystemMenus
	)
	err = dao.SystemMenus.Ctx(ctx).Fields("DISTINCT(unique_auth)").Where("unique_auth!=", "").Scan(&list)
	if err != nil {
		return
	}

	for _, v := range list {
		out = append(out, v.UniqueAuth)

	}
	return
}

// TreeList 树状列表数据
func (s *sAuthsMenu) TreeList(ctx context.Context, req *auths.MenuTreeListReq) (out []system.MenuTreeListItem, err error) {
	err = dao.SystemMenus.Ctx(ctx).Fields("id as value, title as label").Where("pid=?", 0).Where("is_del=?", 0).Scan(&out)
	if err != nil {
		return
	}
	s.TreeListItem(ctx, out)
	return
}

func (s *sAuthsMenu) TreeListItem(ctx context.Context, list []system.MenuTreeListItem) (out []system.MenuTreeListItem, err error) {
	for _, v := range list {
		err = dao.SystemMenus.Ctx(ctx).Fields("id as value, title as label").Where("pid=?", v.Value).Scan(&v.Children)
		if err != nil {
			continue
		}
		out = append(out, v)
		if len(v.Children) > 0 {
			s.TreeListItem(ctx, v.Children)
		}
	}
	return
}
