package common

import (
	"hotgo/internal/model/input/adminin"

	"github.com/gogf/gf/v2/frame/g"
)

// LoginLogoutReq 注销登录
type LoginLogoutReq struct {
	g.Meta `path:"/site/logout" method:"post" tags:"后台基础" summary:"注销登录"`
}

type LoginLogoutRes struct{}

// AccountLoginReq 提交账号登录
type AccountLoginReq struct {
	g.Meta `path:"/site/accountLogin" method:"post" tags:"后台基础" summary:"账号登录"`
	adminin.AccountLoginInp
}

type AccountLoginRes struct {
	*adminin.LoginModel
}
