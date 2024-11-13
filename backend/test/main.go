package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

type CC struct {
	ID        string `json:"id"`
	Thumbnail string `json:"thumbnail"`
}

func download(url string, name string) error {
	if strings.HasPrefix(url, "data:") {
		return nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("请求错误: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("请求失败: %d", resp.StatusCode)
	}

	fileName := fmt.Sprintf("%s.png", name)
	f, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("创建文件错误: %v", err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return fmt.Errorf("下载错误: %v", err)
	}

	fmt.Printf("图片 %s 下载成功\n", fileName)
	return nil
}

func main() {
	var cs []*CC
	data, err := os.ReadFile("demo.json")
	if err != nil {
		fmt.Printf("读取文件错误: %v\n", err)
		return
	}

	if err := json.Unmarshal(data, &cs); err != nil {
		fmt.Printf("解析 JSON 错误: %v\n", err)
		return
	}

	var w sync.WaitGroup
	sem := make(chan struct{}, 5) // 控制并发数为 5

	for _, c := range cs {
		if c.ID != "" && c.Thumbnail != "" {
			fmt.Println(c.ID, c.Thumbnail)
			w.Add(1)
			go func(c *CC) {
				defer w.Done()
				sem <- struct{}{} // 获取一个许可
				if err := download(c.Thumbnail, c.ID); err != nil {
					fmt.Printf("下载错误: %v\n", err)
				}
				<-sem // 释放一个许可
			}(c)
		}
	}

	w.Wait()
}
