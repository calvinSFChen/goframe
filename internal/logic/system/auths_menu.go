package system

import (
	"context"
	"database/sql"
	"fmt"

	"goframe/api/v1/backend/system/auths"
	"goframe/internal/dao"
	"goframe/internal/model/entity"
	"goframe/internal/model/system"
	"goframe/internal/service"
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
	out.List = s.GetMenuItem(ctx, list, []int{})
	out.Total = total
	return
}

func (s *sAuthsMenu) GetMenuItem(ctx context.Context, list []system.MenuListItem, inIds []int) (res []system.MenuListItem) {
	for _, v := range list {
		model := dao.SystemMenus.Ctx(ctx).Where("is_show=?", 1).Where("pid=?", v.Id)
		model.Order("sort desc").Scan(&v.Children)
		if len(v.Children) > 0 && len(inIds) > 0 {
			var tmp = system.MenuListItem{}
			for _, v2 := range v.Children {
				var flag bool
				for _, id := range inIds {
					if v2.Id == id {
						flag = true
						break
					}
				}
				if flag {
					tmp.Children = append(tmp.Children, v2)
				}
			}

			if len(tmp.Children) > 0 {
				v.Children = tmp.Children
			}
		}
		res = append(res, v)
		if len(v.Children) > 0 {
			s.GetMenuItem(ctx, v.Children, inIds)
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
		apiUrl     = gstr.Trim(req.ApiUrl)
		uniqueAuth = gstr.Trim(req.UniqueAuth)
		icon       = gstr.Trim(req.Icon)
		sort       = req.Sort
		isShow     = req.IsShow

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
	systemMenu.ApiUrl = apiUrl
	systemMenu.UniqueAuth = uniqueAuth
	systemMenu.Sort = int(sort)
	systemMenu.IsShow = uint(isShow)
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
		apiUrl     = gstr.Trim(req.ApiUrl)
		uniqueAuth = gstr.Trim(req.UniqueAuth)
		icon       = gstr.Trim(req.Icon)
		sort       = req.Sort
		isShow     = req.IsShow
		systemMenu entity.SystemMenus
	)

	// 过滤特殊代码
	//if err = ghtml.SpecialCharsMapOrStruct(req); err != nil {
	//	return
	//}
	fmt.Printf("menu: %+v\n", req)
	if id <= 0 {
		err = gerror.New(utils.T(ctx, "参数异常"))
		return
	}
	err = dao.SystemMenus.Ctx(ctx).Where("id=", id).Scan(&systemMenu)
	if err != nil {
		err = gerror.New(utils.T(ctx, "参数异常"))
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
	systemMenu.ApiUrl = apiUrl
	systemMenu.UniqueAuth = uniqueAuth
	systemMenu.Sort = int(sort)
	systemMenu.IsShow = uint(isShow)
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
	err = dao.SystemMenus.Ctx(ctx).Fields("id as value, title as label").Where("pid=?", 0).Scan(&out)
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

func (s *sAuthsMenu) GetAuthMenus(ctx context.Context, id int) (out []system.MenuListItem, uniqueAuth []string, routeList []system.SystemRouteOut, err error) {
	var (
		apiList []entity.SystemMenus
	)

	routeList, err = service.AuthsRoute().GetAuthRoutes(ctx, id)
	if err != nil {
		return
	}

	err = dao.SystemMenus.Ctx(ctx).Where("unique_auth!=?", "").Scan(&apiList)
	if err != nil {
		return
	}

	var (
		menuIds []int
	)
	for _, v := range apiList {
		var flag bool = true
		for _, item := range routeList {
			if v.ApiUrl == item.ApiUrl {
				flag = true
			}
		}

		if flag {
			menuIds = append(menuIds, int(v.Id))
			uniqueAuth = append(uniqueAuth, v.UniqueAuth)
		}
	}

	// 组装前端数据
	err = dao.SystemMenus.Ctx(ctx).WhereIn("id", menuIds).Where("pid=?", 0).Scan(&out)
	if err != nil {
		return
	}
	out = s.GetMenuItem(ctx, out, menuIds)
	return
}
