package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserProxy  zrpc.RpcClientConf
	Mysql      Mysql
	CacheRedis cache.CacheConf
	Auth       Auth
}
type Mysql struct {
	DataSource string
}

type Auth struct {
	AccessSecret string
	AccessExpire int64
}
