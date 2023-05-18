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
