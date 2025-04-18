package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	LogConf  logx.LogConf
	LegoConf Lego
}

type Lego struct {
	Email          string
	PrivateKeyStre string
	Env            string
	DomainMap      map[string]map[string]string
}
