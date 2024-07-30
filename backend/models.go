package main

import "time"

type Category struct {
	Name        string
	DisplayName string
	Icon        string
}
type Nav struct {
	ID          string     // 唯一标识符
	Title       string     // 导航标题
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
