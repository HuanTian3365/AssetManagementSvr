package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	"asset_management_svr/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
