package setting

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"goframe/api/v1/backend/setting/rabc"
	"goframe/internal/dao"
	"goframe/internal/model/entity"
	"goframe/internal/model/setting"
	"goframe/utility/utils"
)

type sRabcMenu struct {
}

func NewRabcMenu() *sRabcMenu {
	return &sRabcMenu{}
}

// List 管理员列表
func (s *sRabcMenu) List(ctx context.Context, input *rabc.MenuListReq) (out []setting.MenuListItem, err error) {

	var (
		isShow = input.IsShow

		list []setting.MenuListItem
	)
	model := dao.MuRabcMenus.Ctx(ctx).Safe(false)
	if isShow != -1 {
		model.Where("is_show=?", isShow)
	}
	err = model.Where("pid=?", 0).Order("sort desc").Scan(&list)
	//fmt.Printf("list: %+v\n", list)

	if err != nil {
		return
	}
	out = s.GetMenuItem(ctx, list)
	return
}

func (s *sRabcMenu) GetMenuItem(ctx context.Context, list []setting.MenuListItem) (res []setting.MenuListItem) {
	for _, v := range list {
		dao.MuRabcMenus.Ctx(ctx).Where("is_show=?", 1).Where("pid=?", v.Id).Order("sort desc").Scan(&v.Children)
		res = append(res, v)
		if len(v.Children) > 0 {
			s.GetMenuItem(ctx, v.Children)
		}
	}
	return
}

// Add 添加
func (s *sRabcMenu) Add(ctx context.Context, req *rabc.MenuAddReq) (err error) {
	var (
		pid    = req.Pid
		title  = gstr.Trim(req.Title)
		icon   = gstr.Trim(req.Icon)
		sort   = req.Sort
		path   = gstr.Trim(req.Path)
		isShow = req.IsShow

		systemMenu entity.MuRabcMenus
		lastId     int64
	)
	if pid > 0 {
		var pidNumber int
		pidNumber, err = dao.MuRabcMenus.Ctx(ctx).Where("pid=?", pid).Count()
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

	systemMenu.Pid = uint(pid)
	systemMenu.Title = title
	systemMenu.Sort = int(sort)
	systemMenu.IsShow = int(isShow)
	lastId, err = dao.MuRabcMenus.Ctx(ctx).Data(systemMenu).InsertAndGetId()
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
func (s *sRabcMenu) Edit(ctx context.Context, req *rabc.MenuEditReq) (err error) {
	var (
		id     = req.Id
		pid    = req.Pid
		title  = gstr.Trim(req.Name)
		icon   = gstr.Trim(req.Icon)
		sort   = req.Sort
		path   = gstr.Trim(req.Path)
		isShow = req.Status

		systemMenu entity.MuRabcMenus
	)
	// 过滤特殊代码
	//if err = ghtml.SpecialCharsMapOrStruct(req); err != nil {
	//	return
	//}

	err = dao.MuRabcMenus.Ctx(ctx).Where("id=", id).Scan(&systemMenu)
	if err != nil {
		return
	}
	if id <= 0 {
		gerror.New(utils.T(ctx, "参数异常"))
		return
	}
	systemMenu.Pid = uint(pid)
	systemMenu.Icon = icon
	systemMenu.Title = title
	systemMenu.Path = path
	systemMenu.Sort = int(sort)
	systemMenu.IsShow = int(isShow)
	model := dao.MuRabcMenus.Ctx(ctx)
	_, err = model.Where("id=?", id).Data(systemMenu).Update()
	if err != nil {
		return
	}
	return
}

func (s *sRabcMenu) GetAllUniquePath(ctx context.Context) (uniqueAuthList []string, err error) {
	var (
		list string
	)
	models := dao.MuRabcMenus.Ctx(ctx).Fields("id").Where("unique_auth!=", "").Scan(&list)
	fmt.Printf("uniqueList: %+v", list)

	if text := models.Error(); text != "" {
		return
	}
	return
}
