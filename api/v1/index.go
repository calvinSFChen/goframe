package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type IndexReq struct {
	g.Meta `path:"/index" tags:"测试" method:"post" summary:"You first hello api"`
	Id     string `json:"id" description:"ID"`
}
type IndexRes struct {
}
