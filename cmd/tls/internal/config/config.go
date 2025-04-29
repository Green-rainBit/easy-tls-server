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
	Env            string `json:",optional"`
	DomainMap      map[string]map[string]string
}
