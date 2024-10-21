package main

import (
	"hotgo/internal/global"
	_ "hotgo/internal/packed"

	_ "hotgo/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"hotgo/internal/cmd"
)

func main() {
	ctx := gctx.GetInitCtx()
	global.Init(ctx)
	cmd.Main.Run(ctx)
}
