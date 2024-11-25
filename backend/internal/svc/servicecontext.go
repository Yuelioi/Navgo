package svc

import (
	"backend/internal/common/config"
)

type ServiceContext struct {
	Config config.Config
	Auth   struct {
		AccessSecret string
		AccessExpire int64
	}
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
