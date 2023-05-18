package auths

import (
	"goframe/internal/model/entity"
	"goframe/internal/model/system"

	"github.com/gogf/gf/v2/frame/g"
)

type AdminListReq struct {
	g.Meta   `path:"auths_admin/list"  tags:"系统管理员" method:"get" description:"系统管理员列表"`
	Page     int    `json:"page" default:"1" v:"page@min:1#页码不能为空|页码需大于或等于1" description:"页码"`
	PageSize int    `json:"pag_size" default:"10" v:"page_size@min:1#数量不能为空|数量需大于或等于1" description:"数量"`
	Username string `json:"username" description:"账号"`
	Realname string `json:"realname" description:"姓名"`
	Phone    string `json:"phone" description:"手机号码"`
	Email    string `json:"email" description:"邮箱"`
	Status   int    `json:"status" default:"-1" v:"status@in:-1,0,1,2 #状态参数异常" description:"状态"`
}

type AdminListRes struct {
	g.Meta `mime:"application/json" tags:"系统管理员" example:"string"`
	Total  int                  `json:"total" description:"总条数"`
	List   []system.SystemAdminListOut `json:"list" description:"列表数据"`
}



type AdminAddReq struct {
	g.Meta   `path:"auths_admin/add"  tags:"系统管理员" method:"post" description:"添加系统管理员"`
	Username string `json:"username" v:"username@required #姓名不能为空" description:"姓名"`
	Realname string `json:"realname" v:"real_name@required   #姓名不能为空" description:"姓名"`
	Phone    string `json:"phone" v:"phone@phone #手机号码格式有误" description:"手机号码"`
	Email    string `json:"email" v:"email@email #邮箱格式有误" description:"邮箱"`
	HeadPic  string `json:"head_pic" description:"头像"`
	Password string `json:"password" v:"password@required|length:6,10 #密码不能为|密码长度为6~10" description:"密码"`
}
type AdminAddRes struct {
	g.Meta `mime:"application/json" tags:"系统管理员" example:"string"`
}

type AdminEditReq struct {
	g.Meta   `path:"auths_admin/edit"  tags:"系统管理员" method:"post" description:"编辑系统管理员"`
	Id       uint   `json:"id" v:"id@required #参数异常" description:"ID"`
	Realname string `json:"realname" v:"real_name@required   #姓名不能为空" description:"姓名"`
	Phone    string `json:"phone" v:"phone@phone #手机号码格式有误" description:"手机号码"`
	Email    string `json:"email" v:"email@email #邮箱格式有误" description:"邮箱"`
	HeadPic  string `json:"head_pic" description:"头像"`
	Password string `json:"password" v:"password@required|length:6,10 #密码不能为|密码长度为6~10" description:"密码"`
}
type AdminEditRes struct {
	g.Meta `mime:"application/json" tags:"系统管理员" example:"string"`
}

type AdminLoginReq struct {
	g.Meta   `path:"auths_admin/login"  tags:"系统管理员" method:"post" description:"编辑系统管理员"`
	Username string `json:"username" v:"id@required #名称错误" description:"名称"`
	Password string `json:"password" v:"password@length:6,10 #密码长度6~10" description:"密码"`
}
type AdminLoginRes struct {
	g.Meta     `mime:"application/json" tags:"系统管理员" example:"string"`
	Token      string                `json:"token" description:"token"`
	ExpireTime string                `json:"expire_time" description:"过期时间"`
	AdminInfo  entity.SystemAdmin    `json:"admin_info" description:"用户信息"`
	Menus      []system.MenuListItem `json:"menus"`
	UniqueAuth []string              `json:"unique_auth"`
}

type AdminLogoutReq struct {
	g.Meta `path:"auths_admin/logout"  tags:"系统管理员" method:"get" description:"编辑系统管理员"`
}
type AdminLogoutRes struct {
	g.Meta `mime:"application/json" tags:"系统管理员" example:"string"`
}
