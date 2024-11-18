package main

import (
	"fmt"
	"path"
)

// GetDepth 计算路径的深度
func GetDepth(root, relativeDirPath string) int {
	depth := 0
	for relativeDirPath != "" {
		relativeDirPath, _ = path.Split(relativeDirPath)
		depth++
	}
	return depth - 1 // 减去根目录
}

func main() {
	// 测试路径
	root := "C:\\Users\\User"
	relativeDirPathWindows := "C:\\Users\\User\\Documents\\Projects"
	// relativeDirPathLinux := "/home/user/documents/projects"

	fmt.Printf(root)

	// 计算路径深度
	depthWindows := GetDepth(root, relativeDirPathWindows)
	// depthLinux := GetDepth(relativeDirPathLinux)

	fmt.Printf("Depth of Windows path '%s': %d\n", relativeDirPathWindows, depthWindows)
	// fmt.Printf("Depth of Linux path '%s': %d\n", relativeDirPathLinux, depthLinux)
}
