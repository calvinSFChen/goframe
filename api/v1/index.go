package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type Req struct {
	g.Meta `path:"/" tags:"测试" method:"all" summary:"You first hello api"`
}
type Res struct {
}
