package controller

import (
	"context"
	"fmt"
	"goframe/api/v1"
)

type cIndex struct {
}

var Index = cIndex{}

func (c *cIndex) Index(ctx *context.Context, req *v1.Req) (res *v1.Res, err error) {
	fmt.Println("this is index")
	return
}
