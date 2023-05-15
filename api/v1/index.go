package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type IndexReq struct {
	g.Meta `path:"/index" tags:"测试" method:"get" summary:"You first hello api"`
}
type IndexRes struct {
}
