package setting

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

type BasicSetup struct {
	Name string `json:"name" description:"名称"`
	Type string `json:"type" description:"类型"`
}
