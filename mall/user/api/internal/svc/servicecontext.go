package svc

import (
	"go-zero-demo2/mall/user/api/internal/config"
	"go-zero-demo2/mall/user/api/internal/middleware"
	"go-zero-demo2/mall/user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config        config.Config
	UserinfoModel model.UserinfoModel
	RouteExample  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		UserinfoModel: model.NewUserinfoModel(conn, c.CacheRedis),
		RouteExample:  middleware.NewRouteExampleMiddleware().Handle,
	}
}
