package svc

import (
	"log"
	"time"

	"esay-tls-server/cmd/tls/internal/config"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config      config.Config
	LegoSerivce *LegoSerivce
	MyTlsCache  *collection.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.MustSetup(c.LogConf)
	myTlsCache, err := collection.NewCache(time.Minute, collection.WithLimit(10000))
	if err != nil {
		log.Fatal("collection new cache: ", err)
	}
	return &ServiceContext{
		Config:      c,
		LegoSerivce: NewLegoSerivce(c.LegoConf.Email, c.LegoConf.PrivateKeyStre, c.LegoConf.Env),
		MyTlsCache:  myTlsCache,
	}

}
