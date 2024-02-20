package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_zero/study_model/user/api/internal/config"
	"go_zero/study_model/user/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.Database)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(mysqlConn),
	}
}
