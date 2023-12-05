package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go_zero_study/search/internal/config"
	"go_zero_study/search/internal/model"
	"go_zero_study/user/usercenter"
)

type ServiceContext struct {
	Config config.Config
	UC     usercenter.UserCenter
	Model  model.OrdersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UC:     usercenter.NewUserCenter(zrpc.MustNewClient(c.UserProxy)),
		Model:  model.NewOrdersModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
	}
}
