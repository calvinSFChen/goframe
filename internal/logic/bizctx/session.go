package bizctx

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"goframe/internal/consts"
)

func (s *sBizCtx) GetUserInfo(ctx context.Context) (adminInfo g.Map) {
	var info *gvar.Var
	info, _ = s.GetCtx(ctx).Session.Get(consts.UserInfoKey)
	info.Scan(&adminInfo)
	return
}

func (s *sBizCtx) Operator(ctx context.Context) (operator interface{}) {
	info := s.GetUserInfo(ctx)
	if len(info) > 0 {
		operator = (info["username"]).(string)
	}
	return
}
