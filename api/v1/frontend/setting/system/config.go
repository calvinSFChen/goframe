package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe/internal/model/setting"
)

type BasicConfigListReq struct {
	g.Meta   `path:"system_config/list"  tags:"系统配置管理" method:"get" description:"系统配置管理列表"`
	Page     int    `json:"page" default:"1" v:"page@min:1#页码不能为空|页码需大于或等于1" description:"页码"`
	PageSize int    `json:"pag_size" default:"10" v:"page_size@min:1#数量不能为空|数量需大于或等于1" description:"页码"`
	TypeName string `json:"type_name" default:"basic_setup" description:"配置类型"`
	Name     string `json:"name" description:"配置名称"`
}
type BasicConfigListRes struct {
	g.Meta `mime:"application/json" tags:"系统配置管理" example:"string"`
	Total  int                       `json:"total" description:"总条数"`
	List   []setting.SystemConfigRes `json:"list" description:"列表数据"`
}
