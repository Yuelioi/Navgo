package main

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Model struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type User struct {
	Model `json:"-" yaml:"-"`
	Email string `json:"email"`
}

// func (m Model) MarshalYAML() (interface{}, error) {
// 	return nil, nil // 返回 nil 以跳过 Model 字段的序列化
// }

// 自定义 User 的 MarshalYAML 方法
func (u User) MarshalYAML() (interface{}, error) {
	return struct {
		Email string `json:"email"`
	}{
		Email: u.Email,
	}, nil
}

func main() {
	m := &User{
		Model: Model{Name: "John", Age: 30},
		Email: "231515",
	}

	// 这里将会调用 YamlMarshal 方法
	data, err := yaml.Marshal(m)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}

	fmt.Println(string(data)) // 输出将只包含 Email 字段
}
