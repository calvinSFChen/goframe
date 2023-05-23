package system

import (
	"context"
	"encoding/json"
	"goframe/api/v1/backend/system/config"
	"goframe/internal/dao"
	"goframe/internal/model/entity"
	"goframe/utility/utils"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/text/gstr"
)

type sSystemConfig struct {
}

func NewSystemConfig() *sSystemConfig {
	return &sSystemConfig{}
}

// func (s *sSystemConfig) List(ctx context.Context, input *system.BasicConfigListReq) (total int, out []setting.SystemConfigRes, err error) {
// 	var (
// 		page     = input.Page
// 		pageSize = input.PageSize
// 		typeName = input.TypeName
// 		name     = input.Name

// 		systemConfigList []entity.MuSystemConfig
// 	)
// 	model := g.Model(dao.MuSystemConfig.Table())
// 	if typeName != "" {
// 		model.Where("type_name=?", typeName)
// 	}

// 	if name != "" {
// 		model.Where("name=?", name)
// 	}
// 	if total, err = model.Count(); err != nil {
// 		return
// 	}
// 	err = model.Page(page, (page-1)*pageSize).Scan(&systemConfigList)
// 	if err != nil {
// 		return
// 	}

// 	for i, v := range systemConfigList {
// 		var tmpSystemConfig = setting.SystemConfigRes{}
// 		utils.StructToStruct(v, &tmpSystemConfig)

// 		var tmpConfig interface{}
// 		err = json.Unmarshal([]byte(v.Config), &tmpConfig)
// 		if err != nil {
// 			return
// 		}

// 		out = append(out, tmpSystemConfig)
// 		out[i].Config = tmpConfig
// 	}
// 	return
// }

// Add 添加系统配置
func (s *sSystemConfig) Add(ctx context.Context, input *config.ConfigAddReq) (err error) {
	var (
		typed  = input.Type
		name   = gstr.Trim(input.Name)
		config = input.Config

		systemConfig entity.SystemConfig
	)
	if !json.Valid([]byte(config)) {
		err = gerror.New(utils.T(ctx, "配置参数须json结构异常"))
		return
	}

	systemConfig.Config = config
	systemConfig.Type = int(typed)
	systemConfig.Name = name
	systemConfig.Operator = utils.GetUserName()

	var lastId int64
	lastId, err = dao.SystemConfig.Ctx(ctx).Data(systemConfig).InsertAndGetId()
	if err != nil {
		gerror.New(gi18n.T(ctx, "操作数据库异常"))
		return
	}

	if lastId <= 0 {
		gerror.New(utils.T(ctx, "操作失败"))
		return
	}
	return
}

// Edit 编辑系统配置
func (s *sSystemConfig) Edit(ctx context.Context, input *config.ConfigEditReq) (err error) {
	var (
		id     = input.Id
		typed  = input.Type
		name   = gstr.Trim(input.Name)
		config = gstr.Trim(input.Config)

		systemConfig entity.SystemConfig
	)

	if id <= 0 {
		err = gerror.New(utils.T(ctx, "参数异常"))
		return
	}

	err = dao.SystemConfig.Ctx(ctx).Where("id=?", id).Scan(&systemConfig)
	if err != nil {
		err = gerror.New(utils.T(ctx, "操作数据库异常"))
		return
	}

	systemConfig.Type = int(typed)
	systemConfig.Name = name
	systemConfig.Operator = "system"
	if !json.Valid([]byte(config)) {
		err = gerror.New(gi18n.T(ctx, "配置参数须json结构异常"))
		return
	}
	systemConfig.Config = config
	var affected int64
	affected, err = dao.SystemConfig.Ctx(ctx).Data(systemConfig).Where("id", id).UpdateAndGetAffected(systemConfig)
	if err != nil {
		gerror.New(gi18n.T(ctx, "操作数据库异常"))
		return
	}

	if affected <= 0 {
		gerror.New(gi18n.T(ctx, "操作失败"))
		return
	}
	return
}

// GetOne 获取一条数据
func (s *sSystemConfig) GetOne(ctx context.Context, id uint64) (out *config.ConfigOneRes, err error) {
	var (
		systemConfig entity.SystemConfig
	)
	out = &config.ConfigOneRes{}

	if id <= 0 {
		err = gerror.New(utils.T(ctx, "参数异常"))
		return
	}
	err = dao.SystemConfig.Ctx(ctx).Where("id = ?", id).Scan(&systemConfig)
	if err != nil {
		return
	}
	utils.StructToStruct(systemConfig, out)
	err = json.Unmarshal([]byte(systemConfig.Config), &out.Config)
	if err != nil {
		return
	}
	return
}
