package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func walk(dir string) error {
	// 使用 filepath.Walk 函数来遍历目录
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 打印文件或目录的路径
		fmt.Println(path)
		return nil
	})
}

func main() {

	// 定义要遍历的根目录
	root := "E:\\Scripting\\nav.yuelili.com\\backend\\resource"

	// 调用 walk 函数开始遍历
	err := walk(root)
	if err != nil {
		log.Fatal(err)
	}
}
