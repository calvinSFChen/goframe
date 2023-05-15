// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RabcRolePermission is the golang structure of table my_rabc_role_permission for DAO operations like Where/Data.
type RabcRolePermission struct {
	g.Meta       `orm:"table:my_rabc_role_permission, do:true"`
	Id           interface{} //
	RoleId       interface{} // 角色ID
	PermissionId interface{} // 权限ID
	CreateTime   *gtime.Time // 创建时间
	UpdateTime   *gtime.Time // 更新时间
}