package rabc

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe/internal/model/entity"
)

type RouteListReq struct {
	g.Meta   `path:"rabc_route/list"  tags:"系统管理员" method:"get" description:"系统管理员列表"`
	Page     int    `json:"page" default:"1" v:"page@min:1#页码不能为空|页码需大于或等于1" description:"页码"`
	PageSize int    `json:"pag_size" default:"10" v:"page_size@min:1#数量不能为空|数量需大于或等于1" description:"数量"`
	Name     string `json:"name" description:"名称"`
	Status   int8   `json:"status" default:"-1"  description:"是否显示"`
}
type RouteListRes struct {
	g.Meta `mime:"application/json" tags:"系统管理员" example:"string"`
	Total  int                  `json:"total" description:"总数量"`
	List   []entity.MuRabcRoute `json:"list" description:"列表数据"`
}

type RouteAddReq struct {
	g.Meta `path:"rabc_route/add"  tags:"系统管理员" method:"post" description:"添加系统管理员"`
	MenuId uint64 `json:"menu_id" v:"menu_id@required #参数异常" description:"菜单ID"`
	Method string `json:"method" v:"method@required|in:POST,GET,PUT,DELETE  #提交方式不为空|提交方式类型异常" description:"提交方式"`
	Path   string `json:"path" v:"path@required #路径不为空" description:"路径"`
	Name   string `json:"name" v:"name@required #名称不为空" description:"名称"`
	Sort   int64  `json:"sort" v:"sort@required #排序不为空" description:"排序"`
	Status uint8  `json:"status" v:"status@required #请选择是否隐藏" description:"是否显示"`
}

type RouteAddRes struct {
	g.Meta `mime:"application/json" tags:"系统管理员" example:"string"`
}

type RouteEditReq struct {
	g.Meta `path:"rabc_route/edit"  tags:"系统管理员" method:"post" description:"编辑系统管理员"`
	Id     uint64 `json:"id" v:"id@required #参数异常" description:"id"`
	MenuId uint64 `json:"menu_id" v:"menu_id@required #参数异常" description:"菜单ID"`
	Method string `json:"method" v:"method@required|in:POST,GET,PUT,DELETE  #提交方式不为空|提交方式类型异常" description:"提交方式"`
	Path   string `json:"path" v:"path@required #路径不为空" description:"路径"`
	Name   string `json:"name" v:"name@required #名称不为空" description:"名称"`
	Sort   int64  `json:"sort" v:"sort@required #排序不为空" description:"排序"`
	Status uint8  `json:"status" v:"status@required #请选择是否隐藏" description:"是否显示"`
}
type RouteEditRes struct {
	g.Meta `mime:"application/json" tags:"系统管理员" example:"string"`
}
