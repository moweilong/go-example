package main

import (
	_ "hotgo/internal/packed"

	_ "hotgo/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"hotgo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
