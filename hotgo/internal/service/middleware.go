// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// AdminAuth 后台鉴权中间件
		AdminAuth(r *ghttp.Request)
		// IsExceptLogin 是否是不需要登录的路由地址
		IsExceptLogin(ctx context.Context, appName string, path string) bool
		// ResponseHandler HTTP响应预处理
		ResponseHandler(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
