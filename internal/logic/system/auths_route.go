package system

import (
	"context"
	"fmt"
	"goframe/api/v1/backend/system/auths"
	"goframe/internal/consts"
	"goframe/internal/dao"
	"goframe/internal/model/entity"
	"goframe/internal/model/system"
	"goframe/utility/utils"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
)

type sAuthsRoute struct {
}

func NewAuthsRoute() *sAuthsRoute {
	return &sAuthsRoute{}
}

// // List 管理员列表
func (s *sAuthsRoute) List(ctx context.Context, input *auths.RouteListReq) (total int, out []system.SystemRouteOut, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize
	)

	model := s.Query(ctx, input)
	if total, err = model.Count(); err != nil {
		return
	}
	model.Page(page, (page-1)*pageSize)
	if err = model.Scan(&out); err != nil {
		return
	}
	return
}

func (s *sAuthsRoute) Query(ctx context.Context, input *auths.RouteListReq) *gdb.Model {
	var (
		title    = input.Title
		status   = input.Status
		pid      = input.Pid
		rouleIds = strings.TrimSpace(input.RouleIds)
	)

	model := dao.SystemRoute.Ctx(ctx).Safe(false)
	if title != "" {
		model.Where("title like ?", fmt.Sprintf("%%%s%%", title))
	}
	if status != -1 {
		model.Where("status = ?", status)
	}

	if pid > 0 {
		model.Where("pid = ?", pid)
	}

	if rouleIds != "" {
		rouleIdsList := strings.Split(rouleIds, ",")
		model.WhereIn("id", rouleIdsList)
	}

	return model
}

// Add 添加
func (s *sAuthsRoute) Add(ctx context.Context, req *auths.RouteAddReq) (err error) {
	var (
		pid     = req.Pid
		title   = gstr.Trim(req.Title)
		methods = gstr.Trim(req.Methods)
		apiUrl  = gstr.Trim(req.ApiUrl)
		sort    = req.Sort
		status  = req.Status

		systemRoute entity.SystemRoute
	)

	var menuCount int
	menuCount, err = dao.SystemMenus.Ctx(ctx).Where("id=?", pid).Count()
	if err != nil {
		return
	}

	if menuCount <= 0 {
		err = gerror.New(utils.T(ctx, "父级分类不存在"))
		return
	}

	systemRoute.Title = title
	systemRoute.Pid = uint(pid)
	systemRoute.Methods = methods
	systemRoute.ApiUrl = apiUrl
	systemRoute.Sort = int(sort)
	systemRoute.Status = uint(status)
	systemRoute.Operator = utils.GetUserName()
	var lastId int64
	lastId, err = dao.SystemRoute.Ctx(ctx).Data(systemRoute).InsertAndGetId()
	if err != nil {
		return
	}
	if lastId <= 0 {
		err = gerror.New(utils.T(ctx, "操作失败"))
		return
	}
	return
}

// Edit 编辑
func (s *sAuthsRoute) Edit(ctx context.Context, input *auths.RouteEditReq) (err error) {
	var (
		id      = input.Id
		pid     = input.Pid
		title   = gstr.Trim(input.Title)
		methods = gstr.Trim(input.Methods)
		apiUrl  = gstr.Trim(input.ApiUrl)
		sort    = input.Sort
		status  = input.Status

		systemRoute entity.SystemRoute
	)
	fmt.Printf("input: %+v\n", input)
	if id <= 0 {
		err = gerror.New(utils.T(ctx, "参数异常"))
		return
	}

	err = dao.SystemRoute.Ctx(ctx).Where(id).Scan(&systemRoute)
	if err != nil {
		err = gerror.New(utils.T(ctx, "数据不存在"))
		return
	}
	var menuCount int
	menuCount, err = dao.SystemMenus.Ctx(ctx).Where("id=?", pid).Count()
	if err != nil {
		return
	}

	if menuCount <= 0 {
		err = gerror.New(utils.T(ctx, "父级分类不存在"))
		return
	}

	systemRoute.Title = title
	systemRoute.Pid = uint(pid)
	systemRoute.Methods = methods
	systemRoute.ApiUrl = apiUrl
	systemRoute.Sort = int(sort)
	systemRoute.Status = uint(status)
	systemRoute.Operator = utils.GetUserName()
	var affected int64
	affected, err = dao.SystemRoute.Ctx(ctx).FieldsEx("id").Data(systemRoute).Where("id=?", id).UpdateAndGetAffected()
	if err != nil {
		return
	}
	if affected <= 0 {
		err = gerror.New(utils.T(ctx, "操作失败"))
		return
	}
	return
}

func (s *sAuthsRoute) GetAuthRoutes(ctx context.Context, id int) (routeList []system.SystemRouteOut, err error) {
	var (
		adminInfo entity.SystemAdmin
	)

	err = dao.SystemAdmin.Ctx(ctx).Where("id=?", id).Scan(&adminInfo)
	if err != nil {
		return
	}

	roleIds := strings.Split(adminInfo.Roles, ",")
	var isSuperAdmin bool
	for _, v := range roleIds {
		roleId, _ := strconv.Atoi(v)
		if roleId == consts.IsSuperAdmin {
			isSuperAdmin = true
		}
	}
	if isSuperAdmin {
		roleIds = []string{}
	}

	var input *auths.RouteListReq = &auths.RouteListReq{
		Page:     1,
		PageSize: 100,
		RouleIds: strings.Join(roleIds, ","),
		Status:   1,
	}
	_, routeList, err = s.List(ctx, input)
	if err != nil {
		return
	}
	return
}
