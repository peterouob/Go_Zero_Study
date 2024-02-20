package svc

import (
	"go_zero/common/init_gorm"
	"go_zero/study_model/user_gorm/api/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     init_gorm.InitGorm(c.Mysql.Database),
	}
}
