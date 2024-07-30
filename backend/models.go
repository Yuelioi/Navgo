package main

import "time"

type Category struct {
	Name  string
	Title string
	Icon  string
}

type DownloadInfo struct {
	Version     string
	SizeKb      int
	UpdatedAt   time.Time // 更新时间
	IsOfficial  bool      // 是否为官方
	Description string    // 描述

}

type Nav struct {
	ID          string     // 唯一标识符
	Name        string     // 导航名称(链接路径)
	Title       string     // 导航标题
	IsOfficial  bool       // 是否为官方
	Alias       []string   // 其他别名
	OfficialURL []string   // 官网地址
	DocsURL     []string   // 文档地址
	DownloadURL []string   // 下载地址
	Categories  []Category // 分类
	Tags        []string   // 标签
	Links       []string   // 相关链接
	Icons       []string   // 图标
	Description string     // 导航描述
	CreatedAt   time.Time  // 创建时间
	UpdatedAt   time.Time  // 更新时间
}
