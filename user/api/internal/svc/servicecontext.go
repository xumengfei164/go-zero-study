package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_zero_study/user/api/internal/config"
	"go_zero_study/user/api/internal/model"
)

type ServiceContext struct {
	Config config.Config
	Model  model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewUserModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
	}
}
