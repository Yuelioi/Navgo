package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Database Database
}

type Database struct {
	Resource string
}
