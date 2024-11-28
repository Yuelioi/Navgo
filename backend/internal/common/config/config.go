package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Resource Resource
	Database Database
	System   System
	Auth     Auth
	Api      Api
}

type Resource struct {
	Collections string // 源文件夹
	Announces   string // 通知文件夹
	Comments    string // 评论文件夹
	Icons       string // 图标文件夹
	MetaFile    string // 元数据文件
	Wallpaper   string // 壁纸文件夹
}

type Auth struct {
	AccessSecret string
	AccessExpire int64
}

type Database struct {
	Resource string // 源文件
}

type System struct {
	Username string // 管理员用户名
	Password string // 管理员密码
}

type Api struct {
	Wallpaper string // 壁纸接口
}
