package system

import "github.com/gogf/gf/v2/os/gtime"

type SystemAdminListOut struct {
	Id         uint        `json:"id"         description:"ID"`
	Username   string      `json:"username"   description:"后台管理员账号"`
	Email      string      `json:"email"      description:"邮箱"`
	Phone      string      `json:"phone"      description:"手机号码"`
	HeadPic    string      `json:"headPic"    description:"管理员头像"`
	Realname   string      `json:"realname"   description:"后台管理员姓名"`
	LastIp     string      `json:"last_ip"     description:"后台管理员最后一次登录ip"`
	LastTime   uint        `json:"last_time"   description:"后台管理员最后一次登录时间"`
	LoginCount uint        `json:"login_count" description:"登录次数"`
	Level      uint        `json:"level"      description:"后台管理员级别"`
	Status     uint        `json:"status"     description:"后台管理员状态 1有效0无效"`
	DivisionId int         `json:"division_id" description:"事业部id"`
	IsDel      uint        `json:"isDel"      description:"是否删除"`
	Operator   string      `json:"operator"   description:"操作人"`
	CreatedAt  *gtime.Time `json:"created_at"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updated_at"  description:"更新时间"`
}


type SystemAdmin struct {
	Id         uint        `json:"id"         description:"ID"`
	IdentityId         uint        `json:"identity_id"         description:"ID"`
	Username   string      `json:"username"   description:"后台管理员账号"`
	Realname   string      `json:"realname"   description:"后台管理员姓名"`
	Email      string      `json:"email"      description:"邮箱"`
	Phone      string      `json:"phone"      description:"手机号码"`
	HeadPic    string      `json:"headPic"    description:"管理员头像"`
	Password   string      `json:"password"   description:"后台管理员密码"`
	Roles      string      `json:"roles"      description:"后台管理员权限"`
	LastIp     string      `json:"lastIp"     description:"后台管理员最后一次登录ip"`
	LastTime   uint        `json:"lastTime"   description:"后台管理员最后一次登录时间"`
	LoginCount uint        `json:"loginCount" description:"登录次数"`
	Level      uint        `json:"level"      description:"后台管理员级别"`
	Status     uint        `json:"status"     description:"后台管理员状态 1有效0无效"`
	DivisionId int         `json:"divisionId" description:"事业部id"`
	Operator   string      `json:"operator"   description:"操作人"`
	IsDel      uint        `json:"isDel"      description:"是否删除"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:"更新时间"`
}