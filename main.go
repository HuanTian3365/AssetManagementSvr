package main

import (
	_ "asset_management_svr/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"asset_management_svr/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
