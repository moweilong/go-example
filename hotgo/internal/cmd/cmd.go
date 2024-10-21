package cmd

import (
	"context"
	"hotgo/internal/library/cache"
	"hotgo/internal/router"
	"hotgo/internal/service"

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

			// 注册全局中间件
			s.BindMiddleware("/*any", []ghttp.HandlerFunc{
				service.Middleware().ResponseHandler, // HTTP响应预处理，在业务处理完成后，对响应结果进行格式化和错误过滤，将处理后的数据发送给请求方
			}...)

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
