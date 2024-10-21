package sys

import (
	"context"
	"hotgo/internal/library/token"
	"hotgo/internal/model"
	"hotgo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterSysConfig(NewSysConfig())
}

type sSysConfig struct{}

func NewSysConfig() *sSysConfig {
	return &sSysConfig{}
}

// InitConfig 初始化系统配置
func (s *sSysConfig) InitConfig(ctx context.Context) {
	if err := s.LoadConfig(ctx); err != nil {
		g.Log().Fatalf(ctx, "InitConfig fail：%+v", err)
	}
}

// LoadConfig 加载系统配置
func (s *sSysConfig) LoadConfig(ctx context.Context) (err error) {
	tc, err := s.GetLoadToken(ctx)
	if err != nil {
		return
	}
	token.SetConfig(tc)

	return
}

// GetLoadToken 获取本地token配置
func (s *sSysConfig) GetLoadToken(ctx context.Context) (conf *model.TokenConfig, err error) {
	err = g.Cfg().MustGet(ctx, "token").Scan(&conf)
	return
}
