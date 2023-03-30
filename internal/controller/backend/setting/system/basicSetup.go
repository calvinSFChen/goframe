package system

import (
	"context"
	"fmt"
	"goframe/api/v1/backend/setting/system"
)

type cBasicSetup struct {
}

func NewBasicSetup() *cBasicSetup {
	return &cBasicSetup{}
}

func (c *cBasicSetup) Add(ctx context.Context, req *system.BasicSetupAddReq) (res *system.BasicSetupAddRes, err error) {

	fmt.Println("this is system Add")
	return
}
