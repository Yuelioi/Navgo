package main

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

func walk(dir string) error {
	// 使用 filepath.Walk 函数来遍历目录
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 打印文件或目录的路径
		fmt.Println(path)

		f, _ := os.Create(filepath.Join(path, "meta.yaml"))
		content := fmt.Sprintf(`title: %s
cid: tool
order: 1
collections:`, info.Name())
		defer f.Close()
		f.Write([]byte(content))
		return nil
	})
}

type CategoryMap map[string][]Collection

type Collection struct {
	CID   string `json:"cid,optional" gorm:"column:cid;unique"`
	Title string `json:"title"`
	Link  string `json:"link"`
	// CategoryID string `json:"-" gorm:"column:category_id;index"` // 外键字段
	// Category    Category `json:"category" gorm:"foreignKey:CategoryID;references:CID"`
	Description string `json:"description,optional"`
	Country     string `json:"country,optional"`
	Proxy       bool   `json:"proxy,optional"`
	Favicon     string `json:"favicon,optional"`
	// Tags        []string `json:"tags,optional" gorm:"type:json"`
	View int `json:"view,optional"`
}

// err => 需要代理
func ping(url string) error {
	// 解析URL以获取主机名和端口
	uri, err := urlToHostPort(url)
	if err != nil {
		return err
	}

	// 设置超时时间
	timeout := 5 * time.Second

	// 尝试建立TCP连接
	conn, err := net.DialTimeout("tcp", uri, timeout)
	if err != nil {
		return err
	}

	// 关闭连接
	conn.Close()
	return nil
}

func urlToHostPort(link string) (string, error) {
	// 默认端口
	defaultPort := "80"

	// 如果是HTTPS链接，使用443端口
	if strings.HasPrefix(link, "https://") {
		defaultPort = "443"
	}

	// 提取主机名
	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	host := u.Hostname()

	// 获取端口号
	port := u.Port()
	if port == "" {
		port = defaultPort
	}

	// 返回主机名和端口号
	return net.JoinHostPort(host, port), nil
}

func urlToHostName(link string) (string, error) {
	// 提取主机名
	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}

	return u.Hostname(), nil
}

func rebuild() {
	r := "C:\\Users\\yl\\Downloads\\output.yaml"

	data, _ := os.ReadFile(r)
	groups := &CategoryMap{}
	yaml.Unmarshal(data, groups)

	for _, c := range *groups {
		for index, site := range c {
			parseUlr, err := url.Parse(site.Link)
			if err != nil {
				fmt.Println(site.Link, err)
				continue
			}
			c[index].CID = parseUlr.Hostname()
			err = ping(site.Link)
			c[index].Proxy = false
			if err != nil {
				fmt.Println("ping err", site.Link, err)
				c[index].Proxy = true
			}

		}
	}

	data2, _ := yaml.Marshal(groups)
	os.WriteFile(r, data2, os.ModePerm)
}
