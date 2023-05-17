package auths

import (
	"goframe/internal/model/system"

	"github.com/gogf/gf/v2/frame/g"
)

type MenuListReq struct {
	g.Meta `path:"auths_menu/list"  tags:"菜单管理" method:"get" description:"菜单管理列表"`
	Page   int    `json:"page" default:"1" description:"页码"`
	Size   int    `json:"size" default:"20" description:"页数"`
	IsShow int8   `json:"isShow" default:"-1" description:"是否显示"`
	Title  string `json:"title" default:"" description:"名称"`
}
type MenuListRes struct {
	g.Meta `mime:"application/json" tags:"菜单管理" example:"string"`
	List   []system.MenuListItem `json:"list" description:"列表数据"`
	Total  int                   `json:"total" description: "总条数"`
}

type MenuAddReq struct {
	g.Meta     `path:"auths_menu/add"  tags:"菜单管理" method:"post" description:"添加菜单"`
	Title      string `json:"title" v:"title@required   #名称不能为空" description:"名称"`
	Pid        uint64 `json:"pid" description:"父类id"`
	Path       string `json:"path" v:"path@required #路径不为空" description:"路径"`
	UniqueAuth string `json:"unique_auth" v:"unique_auth@required #标识符不能空" description:"路径"`
	Icon       string `json:"icon" description:"图标"`
	Sort       int64  `json:"sort" v:"sort@required #排序不为空" description:"排序"`
	IsShow     uint8  `json:"isShow" v:"isShow@required|in:0,1 #请选择是否隐藏|是否隐藏类型异常" description:"是否显示"`
	Status     uint8  `json:"status" v:"status@required|in:0,1 #请选择状态|状态类型异常" description:"是否显示"`
}
type MenuAddRes struct {
	g.Meta `mime:"application/json" tags:"菜单管理" example:"string"`
}

type MenuEditReq struct {
	g.Meta     `path:"auths_menu/edit"  tags:"菜单管理" method:"post" description:"编辑菜单"`
	Id         uint   `json:"id" v:"id@required   #参数异常" description:"ID"`
	Title      string `json:"title" v:"title@required   #名称不能为空" description:"名称"`
	Pid        uint64 `json:"pid" description:"父类id"`
	Path       string `json:"path" v:"path@required #路径不为空" description:"路径"`
	UniqueAuth string `json:"unique_auth" v:"unique_auth@required #标识符不能空" description:"路径"`
	Icon       string `json:"icon" description:"图标"`
	Sort       int64  `json:"sort" v:"sort@required #排序不为空" description:"排序"`
	IsShow     uint8  `json:"isShow" v:"isShow@required|in:0,1 #请选择是否隐藏|是否隐藏类型异常" description:"是否显示"`
	Status     uint8  `json:"status" v:"status@required|in:0,1 #请选择状态|状态类型异常" description:"是否显示"`
}
type MenuEditRes struct {
	g.Meta `mime:"application/json" tags:"菜单管理" example:"string"`
}

type MenuTreeListReq struct {
	g.Meta `path:"auths_menu/tree_list"  tags:"菜单管理" method:"get" description:"菜单tree数据"`
}
type MenuTreeListRes struct {
	g.Meta `mime:"application/json" tags:"菜单管理" example:"string"`
	List   []system.MenuTreeListItem `json:"list" description:"列表数据"`
}

type MenuUniqueAuthListReq struct {
	g.Meta `path:"auths_menu/unique_auth_list"  tags:"菜单管理" method:"get" description:"菜单UniqueAuth数据"`
}
type MenuUniqueAuthListRes struct {
	g.Meta `mime:"application/json" tags:"菜单管理" example:"string"`
	List   []string `json:"list" description:"列表数据"`
}
