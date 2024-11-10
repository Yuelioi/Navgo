package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Collection Collection
	Database   Database
}

type Collection struct {
	Resource string // 源文件夹
	MetaFile string // 元数据文件名
}

type Database struct {
	Resource string // 源文件
}
