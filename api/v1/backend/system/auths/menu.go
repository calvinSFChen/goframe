package auths

import (
	"github.com/gogf/gf/v2/frame/g"
// 	"goframe/internal/model/setting"
)

// type MenuListReq struct {
// 	g.Meta `path:"authsmenu/list"  tags:"系统管理员" method:"get" description:"系统管理员列表"`
// 	IsShow int8 `json:"isShow" default:"-1" v:"status@required #请选择是否隐藏" description:"是否显示"`
// }
// type MenuListRes struct {
// 	g.Meta `mime:"application/json" tags:"系统管理员" example:"string"`
// 	List   []setting.MenuListItem `json:"list" description:"列表数据"`
// }

type MenuAddReq struct {
	g.Meta `path:"auths_menu/add"  tags:"系统管理员" method:"post" description:"添加系统管理员"`
	Pid    uint64 `json:"pid" v:"pid@required #参数异常" description:"父类id"`
	Title  string `json:"title" description:"名称"`
	Icon   string `json:"icon" description:"图标"`
	Sort   int64  `json:"sort" v:"sort@required #排序不为空" description:"排序"`
	IsShow uint8  `json:"isShow" v:"status@required #请选择是否隐藏" description:"是否显示"`
	Path   string `json:"path" v:"path@required #路径不为空" description:"路径"`
}
type MenuAddRes struct {
	g.Meta `mime:"application/json" tags:"系统管理员" example:"string"`
}

// type MenuEditReq struct {
// 	g.Meta `path:"authsmenu/edit"  tags:"系统管理员" method:"post" description:"编辑系统管理员"`
// 	Id     uint64 `json:"id" v:"id@required #参数异常" description:"id"`
// 	Pid    uint64 `json:"pid" v:"pid@required #参数异常" description:"父类id"`
// 	Name   string `json:"name" description:"名称"`
// 	Icon   string `json:"icon" description:"图标"`
// 	Sort   int64  `json:"sort" v:"sort@required #排序不为空" description:"排序"`
// 	Status uint8  `json:"status" v:"status@required #请选择是否隐藏" description:"是否显示"`
// 	Path   string `json:"path" v:"path@required #路径不为空" description:"路径"`
// }
// type MenuEditRes struct {
// 	g.Meta `mime:"application/json" tags:"系统管理员" example:"string"`
// }
