package system

import (
	"context"
	"fmt"
	"goframe/api/v1/backend/system/auths"
	"goframe/internal/consts"
	"goframe/internal/dao"
	"goframe/internal/model/entity"
	"goframe/internal/model/system"
	"goframe/internal/service"
	"goframe/utility/utils"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/text/gstr"
)

type sAuthsAdmin struct {
}

func NewAuthsAdmin() *sAuthsAdmin {
	return &sAuthsAdmin{}
}

// List 管理员列表
func (s *sAuthsAdmin) List(ctx context.Context, input *auths.AdminListReq) (total int, out []system.SystemAdminListOut, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize

		model *gdb.Model
	)

	model = s.Query(input)
	if total, err = model.Count(); err != nil {
		return
	}
	err = model.Fields("id, username, realname, phone, email, operator, createdAt, updatedAt").Page(page, (page-1)*pageSize).Scan(&out)
	if err != nil {
		return
	}
	return
}

func (s *sAuthsAdmin) Query(input *auths.AdminListReq) *gdb.Model {
	var (
		username = gstr.Trim(input.Username)
		realname = gstr.Trim(input.Realname)
		phone    = gstr.Trim(input.Phone)
		email    = gstr.Trim(input.Email)
		status   = input.Status
	)
	model := g.Model(dao.SystemAdmin.Table())
	if username != "" {
		model.Where("username=?", username)
	}
	if phone != "" {
		model.Where("phone=?", phone)
	}
	if email != "" {
		model.Where("email=?", email)
	}

	if realname != "" {
		model.Where("realname like ?", fmt.Sprintf("%%%s%%", realname))
	}
	if status != -1 {
		model.Where("status=?", status)
	}

	return model
}

// Add 添加管理员
func (s *sAuthsAdmin) Add(ctx context.Context, input *auths.AdminAddReq) (err error) {
	var (
		username = gstr.Trim(input.Username)
		realname = gstr.Trim(input.Realname)
		password = gstr.Trim(input.Password)
		phone    = gstr.Trim(input.Phone)
		email    = gstr.Trim(input.Email)
		headPic  = gstr.Trim(input.HeadPic)

		systemAdmin entity.SystemAdmin
		newPassword string
	)

	newPassword = utils.Encryption(password)
	if newPassword == "" {
		err = gerror.New(utils.T(ctx, "密码加密失败"))
		return
	}

	if err = s.IsExitsUser(ctx, 0, username); err != nil {
		return
	}

	systemAdmin.Username = username
	systemAdmin.Password = newPassword
	if phone != "" {
		if err = s.IsExitsUser(ctx, 0, phone); err != nil {
			return
		}
		systemAdmin.Phone = phone
	}
	if email != "" {
		if err = s.IsExitsUser(ctx, 0, email); err != nil {
			return
		}
		systemAdmin.Email = email
	}
	if headPic != "" {
		systemAdmin.HeadPic = headPic
	}
	systemAdmin.Realname = realname
	systemAdmin.Operator = utils.GetUserName()

	model := dao.SystemAdmin.Ctx(ctx).Safe(false)
	var lastId int64
	lastId, err = model.Data(systemAdmin).InsertAndGetId()
	if err != nil {
		return
	}

	if lastId <= 0 {
		err = gerror.New(gi18n.T(ctx, "操作失败"))
		return
	}
	return
}

// IsExitsUser 判断用户是否存在
func (s *sAuthsAdmin) IsExitsUser(ctx context.Context, id uint, name string) (err error) {
	var (
		exitsCount int
		flag       = true
		errText    string
	)

	model := dao.SystemAdmin.Ctx(ctx).Safe(false)

	if utils.IsPhone(name) {
		model.Where("phone=?", name)
		errText = "手机号码已存在"
		flag = false
	}
	if utils.IsEmail(name) {
		model.Where("email=?", name)
		errText = "邮箱已存在"
		flag = false
	}

	if flag {
		model.Where("username=?", name)
		errText = "用户名已存在"
		flag = false
	}
	if !flag {
		if id > 0 {
			model.Where("id!=?", id)
		}
		exitsCount, err = model.Count()
		if err != nil {
			return
		}
		if exitsCount > 0 {
			err = gerror.New(utils.T(ctx, errText))
			return
		}
	}
	return
}

// Add 添加管理员
func (s *sAuthsAdmin) Edit(ctx context.Context, input *auths.AdminEditReq) (err error) {
	var (
		id       = input.Id
		realname = gstr.Trim(input.Realname)
		password = gstr.Trim(input.Password)
		phone    = gstr.Trim(input.Phone)
		email    = gstr.Trim(input.Email)
		headPic  = gstr.Trim(input.HeadPic)

		systemAdmin entity.SystemAdmin
		newPassword string
	)
	if id <= 0 {
		err = gerror.New(utils.T(ctx, "参数异常"))
		return
	}

	err = dao.SystemAdmin.Ctx(ctx).Where("id=?", id).Scan(&systemAdmin)
	if err != nil {
		err = gerror.New(utils.T(ctx, "数据不存在"))
		return
	}
	newPassword = utils.Encryption(password)
	if newPassword == "" {
		err = gerror.New(utils.T(ctx, "密码加密失败"))
		return
	}
	systemAdmin.Password = newPassword
	if phone != "" {
		if err = s.IsExitsUser(ctx, id, phone); err != nil {
			return
		}
		systemAdmin.Phone = phone
	}
	if email != "" {
		if err = s.IsExitsUser(ctx, id, email); err != nil {
			return
		}
		systemAdmin.Email = email
	}
	if headPic != "" {
		systemAdmin.HeadPic = headPic
	}
	systemAdmin.Realname = realname
	systemAdmin.Operator = utils.GetUserName()

	model := dao.SystemAdmin.Ctx(ctx).Safe(false).Where("id=?", id)
	var affected int64
	affected, err = model.Data(systemAdmin).UpdateAndGetAffected()
	if err != nil {
		return
	}

	if affected <= 0 {
		err = gerror.New(gi18n.T(ctx, "操作失败"))
		return
	}
	return
}

// Login 登录
func (s *sAuthsAdmin) Login(ctx context.Context, input *auths.AdminLoginReq) (out *auths.AdminLoginRes, err error) {
	var (
		userName = input.Username

		adminInfo entity.SystemAdmin
	)
	serviceJwt := service.Auth().AuthJwtService()
	token, expireTime := serviceJwt.LoginHandler(ctx)
	if token == "" {
		err = gerror.New("登录失败")
		return
	}

	err = dao.SystemAdmin.Ctx(ctx).WhereOr("username=?", userName).
		WhereOr("phone=?", userName).Scan(&adminInfo)
	if err != nil {
		return
	}
	err = service.BizCtx().GetCtx(ctx).Session.Set(consts.UserInfoKey, adminInfo)
	if err != nil {
		return
	}
	out = &auths.AdminLoginRes{
		Token:      token,
		AdminInfo:  adminInfo,
		ExpireTime: expireTime.Format(consts.TimeFormat),
	}
	return
}

// Logout 退出登录
func (s *sAuthsAdmin) Logout(ctx context.Context) (err error) {
	err = service.BizCtx().GetCtx(ctx).Session.Remove(consts.UserInfoKey)
	if err != nil {
		return
	}
	service.Auth().AuthJwtService().LogoutHandler(ctx)
	return
}
