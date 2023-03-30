package system

import (
	"context"
	"fmt"
	"goframe/api/v1/backend/setting/system"
)

type cUserSetup struct {
}

func NewUserSetup() *cUserSetup {
	return &cUserSetup{}
}

func (c *cUserSetup) Add(ctx context.Context, req *system.UserSetupAddReq) (res *system.UserSetupAddRes, err error) {
	fmt.Println("this is ShopSetup Add")
	return
}
