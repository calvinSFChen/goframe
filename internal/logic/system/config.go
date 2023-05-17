package system

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"github.com/gogf/gf/v2/errors/gerror"
// 	"github.com/gogf/gf/v2/frame/g"
// 	"github.com/gogf/gf/v2/i18n/gi18n"
// 	"github.com/gogf/gf/v2/text/gstr"
// 	"goframe/api/v1/backend/setting/system"
// 	"goframe/internal/dao"
// 	"goframe/internal/model/entity"
// 	"goframe/internal/model/setting"
// 	"goframe/utility/utils"
// )

// type sSystemConfig struct {
// }

// func NewSystemConfig() *sSystemConfig {
// 	return &sSystemConfig{}
// }

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

// // Add 添加系统配置
// func (s *sSystemConfig) Add(ctx context.Context, req *system.BasicConfigAddReq) (err error) {
// 	var (
// 		id           = req.Id
// 		typeName     = gstr.Trim(req.TypeName)
// 		name         = gstr.Trim(req.Name)
// 		config       = gstr.Trim(req.Config)
// 		systemConfig entity.MuSystemConfig
// 	)
// 	if id != 0 {
// 		var n int
// 		n, err = dao.MuSystemConfig.Ctx(ctx).Where("id=?", id).Count()
// 		if err != nil {
// 			err = gerror.New(gi18n.T(ctx, "数据查询异常"))
// 			return
// 		}

// 		if n != 0 {
// 			err = gerror.New(gi18n.T(ctx, "ID已存在"))
// 			return
// 		}

// 		systemConfig.Id = uint(id)
// 	}
// 	systemConfig.TypeName = typeName
// 	systemConfig.Name = name
// 	systemConfig.Operator = "system"
// 	if !json.Valid([]byte(config)) {
// 		err = gerror.New(gi18n.T(ctx, "配置参数须json结构"))
// 		return
// 	}
// 	systemConfig.Config = config
// 	var lastId int64
// 	_, err = dao.MuSystemConfig.Ctx(ctx).Data(systemConfig).InsertAndGetId()
// 	if err != nil {
// 		gerror.New(gi18n.T(ctx, "操作数据库异常"))
// 		return
// 	}

// 	if lastId <= 0 {
// 		gerror.New(gi18n.T(ctx, "插入数据失败"))
// 		return
// 	}
// 	return
// }

// // Edit 编辑系统配置
// func (s *sSystemConfig) Edit(ctx context.Context, input *system.BasicConfigEditReq) (err error) {
// 	var (
// 		id           = input.Id
// 		typeName     = gstr.Trim(input.TypeName)
// 		name         = gstr.Trim(input.Name)
// 		config       = gstr.Trim(input.Config)
// 		systemConfig entity.MuSystemConfig
// 	)
// 	if id != 0 {
// 		var n int
// 		n, err = dao.MuSystemConfig.Ctx(ctx).Where("id=?", id).Count()
// 		if err != nil {
// 			err = gerror.New(gi18n.T(ctx, "数据查询异常"))
// 			return
// 		}

// 		if n != 0 {
// 			err = gerror.New(gi18n.T(ctx, "ID已存在"))
// 			return
// 		}

// 		systemConfig.Id = uint(id)
// 	}
// 	systemConfig.TypeName = typeName
// 	systemConfig.Name = name
// 	systemConfig.Operator = "system"
// 	if !json.Valid([]byte(config)) {
// 		err = gerror.New(gi18n.T(ctx, "配置参数须json结构"))
// 		return
// 	}
// 	systemConfig.Config = config
// 	var lastId int64
// 	_, err = dao.MuSystemConfig.Ctx(ctx).Data(systemConfig).InsertAndGetId()
// 	if err != nil {
// 		gerror.New(gi18n.T(ctx, "操作数据库异常"))
// 		return
// 	}

// 	if lastId <= 0 {
// 		gerror.New(gi18n.T(ctx, "插入数据失败"))
// 		return
// 	}
// 	return
// }

// func (s *sSystemConfig) GetOne(ctx context.Context, id uint64) (out *system.SystemConfigOneRes, err error) {
// 	var (
// 		systemConfig entity.MuSystemConfig

// 		config = system.SystemConfigItem1{}
// 	)
// 	if id <= 0 {
// 		err = gerror.New(utils.T(ctx, "参数异常"))
// 	}
// 	fmt.Printf("id: %+v\n", id)
// 	err = dao.MuSystemConfig.Ctx(ctx).Where("id = ?", id).Scan(&systemConfig)
// 	if err != nil {
// 		return
// 	}
// 	fmt.Println(2323)

// 	err = json.Unmarshal([]byte(systemConfig.Config), &config)
// 	if err != nil {
// 		return
// 	}
// 	out = &system.SystemConfigOneRes{
// 		MuSystemConfig: systemConfig,
// 		Config:         config,
// 	}
// 	return
// }
