package bizctx

import "goframe/internal/service"

func init() {
	service.RegisterBizCtx(NewBizCtx())
}
