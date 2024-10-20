package cmd

import (
	"context"
	"hotgo/internal/library/cache"
	"hotgo/internal/router"

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
			// 初始化http服务
			s := g.Server()

			s.Group("/", func(group *ghttp.RouterGroup) {
				// 注册后台路由
				router.Admin(ctx, group)
			})

			// 设置缓存适配器
			cache.SetAdapter(ctx)

			s.Run()
			return nil
		},
	}
)
