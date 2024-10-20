// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/model/input/adminin"
)

type (
	IAdminSite interface {
		// AccountLogin 账号登录
		AccountLogin(ctx context.Context, in *adminin.AccountLoginInp) (res *adminin.LoginModel, err error)
	}
)

var (
	localAdminSite IAdminSite
)

func AdminSite() IAdminSite {
	if localAdminSite == nil {
		panic("implement not found for interface IAdminSite, forgot register?")
	}
	return localAdminSite
}

func RegisterAdminSite(i IAdminSite) {
	localAdminSite = i
}
