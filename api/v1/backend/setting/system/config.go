package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe/internal/model/entity"
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

type BasicConfigAddReq struct {
	g.Meta   `path:"system_config/add"  tags:"系统配置管理" method:"post" description:"添加系统配置管理"`
	Id       uint64 `json:"id" description:"配置ID"`
	TypeName string `json:"type_name" v:"type_name@required#配置类型不能为空" description:"配置类型"`
	Name     string `json:"name" v:"name@required#配置名称不能为空" description:"配置名称"`
	Config   string `json:"config" description:"配置"`
}
type BasicConfigAddRes struct {
	g.Meta `mime:"application/json" tags:"系统配置管理" example:"string"`
}

type BasicConfigEditReq struct {
	g.Meta   `path:"system_config/edit"  tags:"系统配置管理" method:"post" description:"添加系统配置管理"`
	Id       uint64 `json:"id" description:"配置ID"`
	TypeName string `json:"type_name" v:"type_name@required#配置类型不能为空" description:"配置类型"`
	Name     string `json:"name" v:"name@required#配置名称不能为空" description:"配置名称"`
	Config   string `json:"config" description:"配置"`
}
type BasicConfigEditRes struct {
	g.Meta `mime:"application/json" tags:"系统配置管理" example:"string"`
}

type SystemConfigOenReq struct {
	g.Meta `path:"system_config/get-one"  tags:"系统配置管理" method:"get" description:"添加系统配置管理"`
	Id     uint64 `json:"id" description:"配置ID"`
}

type SystemConfigOneRes struct {
	entity.MuSystemConfig
	Config interface{} `json:"config" description:"配置"`
}

type SystemConfigItem1 struct {
	SiteOpen  string `json:"site_open"  description:"站点开启"`
	SiteName  string `json:"site_name"  description:"网站名称"`
	SiteUrl   string `json:"site_url"  description:"网站地址"`
	LoginLogo string `json:"login_logo"  description:"后台登录页LOGO"`
	SiteLogo  string `json:"site_logo"  description:"站点logo"`
}