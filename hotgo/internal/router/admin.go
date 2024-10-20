package router

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/controller/admin/common"
	"hotgo/internal/controller/hello"
	"hotgo/internal/router/genrouter"
	"hotgo/internal/service"
	"hotgo/utility/simple"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Admin(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group(simple.RouterPrefix(ctx, consts.AppAdmin), func(group *ghttp.RouterGroup) {
		group.Bind(
			common.Site, // 基础路由
		)

		group.Middleware(service.Middleware().AdminAuth)
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			hello.NewV1(),
		)
	})

	// 注册生成路由
	genrouter.Register(ctx, group)
}
