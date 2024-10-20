package common

import (
	"context"
	"hotgo/api/admin/common"
	"hotgo/internal/service"

	"github.com/gogf/gf/util/gconv"
)

var Site = cSite{}

type cSite struct{}

func (c *cSite) AccountLogin(ctx context.Context, req *common.AccountLoginReq) (res *common.AccountLoginRes, err error) {
	// login, err := service.SysConfig().GetLogin(ctx)
	// if err != nil {
	// 	return
	// }

	// if !req.IsLock && login.CaptchaSwitch == consts.StatusEnabled {
	// 	// 校验 验证码
	// 	if !captcha.Verify(req.Cid, req.Code) {
	// 		err = gerror.New("验证码错误")
	// 		return
	// 	}
	// }

	model, err := service.AdminSite().AccountLogin(ctx, &req.AccountLoginInp)
	if err != nil {
		return
	}

	err = gconv.Scan(model, &res)
	return
}
