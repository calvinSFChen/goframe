package setting

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/text/gstr"
	"goframe/api/v1/backend/setting/rabc"
	"goframe/internal/consts"
	"goframe/internal/dao"
	"goframe/internal/model/entity"
	"goframe/internal/service"
	"goframe/utility/utils"
)

type sRabcAdmin struct {
}

func NewRabcAdmin() *sRabcAdmin {
	return &sRabcAdmin{}
}

// List 管理员列表
func (s *sRabcAdmin) List(ctx context.Context, input *rabc.AdminListReq) (total int, out []entity.MuRabcAdmin, err error) {
	t, _, _ := service.Auth().AuthJwtService().GetClaimsFromJWT(ctx)
	fmt.Println("GetIdentity", t)
	var (
		page     = input.Page
		pageSize = input.PageSize

		username = gstr.Trim(input.Username)
		phone    = gstr.Trim(input.Phone)
		email    = gstr.Trim(input.Email)
		status   = input.Status

		//systemConfigList []entity.MuRabcAdmin
	)
	model := g.Model(dao.MuRabcAdmin.Table())
	if username != "" {
		model.Where("username=?", username)
	}
	if phone != "" {
		model.Where("phone=?", phone)
	}
	if email != "" {
		model.Where("email=?", email)
	}

	if status != -1 {
		model.Where("status=?", status)
	}

	if total, err = model.Count(); err != nil {
		return
	}
	err = model.Page(page, (page-1)*pageSize).Scan(&out)
	if err != nil {
		return
	}
	return
}

// Add 添加管理员
func (s *sRabcAdmin) Add(ctx context.Context, input *rabc.AdminAddReq) (err error) {
	var (
		username = gstr.Trim(input.Username)
		password = gstr.Trim(input.Password)
		phone    = gstr.Trim(input.Phone)
		email    = gstr.Trim(input.Email)
		picUrl   = gstr.Trim(input.PicUrl)

		systemAdmin entity.MuRabcAdmin
		lastId      int64
		newPassword string
	)

	newPassword, err = utils.Encryption(password)
	systemAdmin.Username = username
	systemAdmin.Password = newPassword
	systemAdmin.Phone = phone
	if email != "" {
		systemAdmin.Email = email
	}
	systemAdmin.PicUrl = picUrl

	model := dao.MuRabcAdmin.Ctx(ctx).Safe(false)
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

// Edit 编辑管理员
func (s *sRabcAdmin) Edit(ctx context.Context, input *rabc.AdminEditReq) (err error) {
	var (
		id       = input.Id
		password = input.Password
		phone    = input.Phone
		email    = input.Email
		picUrl   = input.PicUrl

		systemAdmin entity.MuRabcAdmin
		newPassword string
	)
	if id <= 0 {
		gerror.New(gi18n.T(ctx, "参数异常"))
		return
	}
	err = dao.MuRabcAdmin.Ctx(ctx).Where("id=?", id).Scan(&systemAdmin)
	if err != nil {
		return
	}
	if password != "" {
		newPassword, err = utils.Encryption(password)
		if err != nil {
			return
		}

		systemAdmin.Password = newPassword
	}
	if phone != "" {
		systemAdmin.Phone = phone
	}
	if email != "" {
		systemAdmin.Email = email
	}
	if picUrl != "" {
		systemAdmin.PicUrl = picUrl
	}
	model := dao.MuRabcAdmin.Ctx(ctx).Safe(false).Data(systemAdmin).Where("id=?", id)
	_, err = model.Update()
	if err != nil {
		return
	}
	return
}

// Login 登录
func (s *sRabcAdmin) Login(ctx context.Context, input *rabc.AdminLoginReq) (out *rabc.AdminLoginRes, err error) {
	var (
		userName = input.Username

		adminInfo entity.MuRabcAdmin
	)
	serviceJwt := service.Auth().AuthJwtService()
	token, expireTime := serviceJwt.LoginHandler(ctx)
	if token == "" {
		err = gerror.New("登录失败")
		return
	}

	err = dao.MuRabcAdmin.Ctx(ctx).WhereOr("username=?", userName).
		WhereOr("email=?", userName).
		WhereOr("phone=?", userName).Scan(&adminInfo)
	if err != nil {
		return
	}
	err = service.BizCtx().GetCtx(ctx).Session.Set(consts.UserInfoKey, adminInfo)
	if err != nil {
		return
	}
	out = &rabc.AdminLoginRes{
		Token:      token,
		AdminInfo:  adminInfo,
		ExpireTime: expireTime.Format(consts.TimeFormat),
	}
	return
}

func (s *sRabcAdmin) Logout(ctx context.Context) (err error) {
	err = service.BizCtx().GetCtx(ctx).Session.Remove(consts.UserInfoKey)
	if err != nil {
		return
	}
	service.Auth().AuthJwtService().LogoutHandler(ctx)
	return
}
