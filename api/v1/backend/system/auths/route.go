package auths

import (
	"goframe/internal/model/system"

	"github.com/gogf/gf/v2/frame/g"
)

type RouteListReq struct {
	g.Meta `path:"auths_route/list"  tags:"系统管理员" method:"get" description:"系统管理员列表"`
	Id     string `json:"id" default:"33333" description:"ID"`

	Page     int    `json:"page" default:"1" v:"page@min:1#页码不能为空|页码需大于或等于1" description:"页码"`
	PageSize int    `json:"pag_size" default:"10" v:"page_size@min:1#数量不能为空|数量需大于或等于1" description:"数量"`
	Title    string `json:"title" description:"名称"`
	Status   int8   `json:"status" default:"-1"  description:"是否显示"`
	Pid      uint   `json:"pid"  description:"是否显示"`
}
type RouteListRes struct {
	g.Meta `mime:"application/json" tags:"系统管理员" example:"string"`
	Total  int                     `json:"total" description:"总数量"`
	List   []system.SystemRouteOut `json:"list" description:"列表数据"`
}

type RouteAddReq struct {
	g.Meta  `path:"auths_route/add"  tags:"系统管理员" method:"post" description:"添加系统管理员"`
	Pid     uint64 `json:"pid" v:"pid@required #父级分类不为空" description:"父级分类ID"`
	Title   string `json:"title" v:"title@required #名称不能为空" description:"提交方式"`
	Methods string `json:"methods" v:"methods@required|in:POST,GET,PUT,DELETE  #提交方式不能为空|提交方式类型异常" description:"提交方式"`
	// Access   string `json:"access" v:"access@required #请选择子管理员是否可用" description:"子管理员权限开关"`
	Path       string `json:"path" v:"path@required #接口地址不能为空" description:"接口地址"`
	UniqueAuth string `json:"unique_auth" v:"unique_auth@required #前端唯一标识不能为空" description:"前端唯一标识"`
	Sort       int64  `json:"sort" v:"sort@required #排序不为空" description:"排序"`
	Status     uint8  `json:"status" v:"status@required #请选择是否隐藏" description:"是否显示"`
}

type RouteAddRes struct {
	g.Meta `mime:"application/json" tags:"系统管理员" example:"string"`
}

type RouteEditReq struct {
	g.Meta  `path:"auths_route/edit"  tags:"系统管理员" method:"post" description:"编辑系统管理员"`
	Id      string `json:"id" default:"33333" description:"ID"`
	Pid     uint64 `json:"pid" v:"pid@required #父级分类不为空" description:"父级分类ID"`
	Title   string `json:"title" v:"title@required #名称不能为空" description:"提交方式"`
	Methods string `json:"methods" v:"methods@required|in:POST,GET,PUT,DELETE  #提交方式不能为空|提交方式类型异常" description:"提交方式"`
	// Access   string `json:"access" v:"access@required #请选择子管理员是否可用" description:"子管理员权限开关"`
	Path       string `json:"path" v:"path@required #接口地址不能为空" description:"接口地址"`
	UniqueAuth string `json:"unique_auth" v:"unique_auth@required #前端唯一标识不能为空" description:"前端唯一标识"`
	Sort       int64  `json:"sort" v:"sort@required #排序不为空" description:"排序"`
	Status     uint8  `json:"status" v:"status@required #请选择是否隐藏" description:"是否显示"`
}
type RouteEditRes struct {
	g.Meta `mime:"application/json" tags:"系统管理员" example:"string"`
}
