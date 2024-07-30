package main

type Nav struct {
	ID          string   // 唯一标识符
	Title       string   // 导航标题
	OfficialURL []string // 官网地址
	DocsURL     []string // 文档地址
	DownloadURL []string // 下载地址
	Category    []string // 分类
	Tags        []string // 标签
	Links       []string // 相关链接
	Icon        []string // 图标
	Uploader    string   // 上传者
	Description string   // 导航描述
	CreatedAt   string   // 创建时间
	UpdatedAt   string   // 更新时间
}
