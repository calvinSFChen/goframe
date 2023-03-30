package system

import (
	"context"
	"fmt"
	"goframe/api/v1/backend/setting/system"
)

type cShopSetup struct {
}

func NewShopSetup() *cShopSetup {
	return &cShopSetup{}
}

func (c *cShopSetup) Add(ctx context.Context, req *system.ShopSetupAddReq) (res *system.ShopSetupAddRes, err error) {
	fmt.Println("this is ShopSetup Add")
	return
}
