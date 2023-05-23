package system

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type SystemConfigRes struct {
	Id        uint        `json:"id"        description:""`
	TypeName  string      `json:"typeName"  description:"配置类型"`
	Name      string      `json:"name"      description:"名称"`
	Config    interface{} `json:"config"    description:"配置内容"`
	Operator  string      `json:"operator"   description:"操作人"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
}

type BasicConfigOut struct {
	SiteName        string `json:"site_name"  description:"网站名称"`
	SiteUrl         string `json:"site_url"      description:"基础网站地址"`
	SiteLogo        string `json:"site_logo"    description:"站点logo"`
	SiteLogoSquare  string `json:"site_logo_square"   description:"后台小logo"`
	LoginLogo       string `json:"login_logo" description:"登录logo"`
	WapLoginLogo    string `json:"wap_login_logo" description:"移动端登录logo"`
	RecordNo        string `json:"record_no" description:"备案"`
	StatisticScript string `json:"statistic_script" description:"统计代码"`
	IconPath        string `json:"icon_path" description:"icon图标"`
}
