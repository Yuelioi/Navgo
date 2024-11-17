package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Resource Resource
	Database Database
}

type Resource struct {
	Collections string // 源文件夹
	Announces   string // 通知文件夹
	Comments    string // 评论文件夹
	Icons       string // 图标文件夹
	MetaFile    string // 元数据文件名
}

type Database struct {
	Resource string // 源文件
}
