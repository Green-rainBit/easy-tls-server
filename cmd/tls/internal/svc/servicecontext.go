package svc

import (
	"esay-tls-server/cmd/tls/internal/config"

	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config      config.Config
	LegoSerivce *LegoSerivce
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.MustSetup(c.LogConf)
	return &ServiceContext{
		Config:      c,
		LegoSerivce: NewLegoSerivce(c.LegoConf.Email, c.LegoConf.PrivateKeyStre, c.LegoConf.Env),
	}
}
