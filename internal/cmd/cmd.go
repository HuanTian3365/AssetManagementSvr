package cmd

import (
	"asset_management_svr/internal/controller/asset"
	"asset_management_svr/internal/middleware"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/asset", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.ResponseHandler)
				group.Bind(
					asset.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
