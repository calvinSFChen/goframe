package system

import "github.com/gogf/gf/v2/os/gtime"

type RouteItem struct {
	Module string      `json:"module" description:"模块"`
	Path   string      `json:"path" description:"地址"`
	Method string      `json:"method" description:"请求方式"`
	IsAuth bool        `json:"is_auth" description:"是否验证"`
	Func   interface{} `json:"func" description:"路由方法"`
}

type SystemRouteOut struct {
	Id         uint        `json:"id"         description:"菜单ID"`
	Pid        uint        `json:"pid"        description:"父级id"`
	Title      string      `json:"title"      description:"名称"`
	Methods    string      `json:"methods"    description:"提交方式POST GET PUT DELETE"`
	Access     uint        `json:"access"     description:"子管理员是否可用"`
	Path       string      `json:"path"       description:"路径"`
	UniqueAuth string      `json:"unique_auth" description:"前台唯一标识"`
	Status     uint        `json:"status"     description:"0关闭1开启"`
	Sort       int         `json:"sort"       description:"排序"`
	Operator   string      `json:"operator"   description:"操作人"`
	CreatedAt  *gtime.Time `json:"created_at"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updated_at"  description:"更新时间"`
}
