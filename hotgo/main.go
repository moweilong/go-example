package main

import (
	_ "hotgo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"hotgo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
