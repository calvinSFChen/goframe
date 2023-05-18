package controller

import (
	"context"
	"fmt"
	v1 "goframe/api/v1"
)

type cIndex struct {
}

func NewIndex() *cIndex {
	return &cIndex{}
}

func (c *cIndex) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	fmt.Printf("this is index, req: %+v\n", req)
	return
}
