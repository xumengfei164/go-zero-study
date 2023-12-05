package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	CacheRedis cache.CacheConf
	Mysql      Mysql
	Auth       Auth
}
type Mysql struct {
	DataSource string
}
type Auth struct {
	AccessSecret string
	AccessExpire int64
}
