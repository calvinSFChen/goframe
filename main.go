package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"

	_ "goframe/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	_ "goframe/internal/logic"

	"goframe/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
