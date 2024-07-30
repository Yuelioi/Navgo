package main

import (
	"encoding/json"
	"os"
)

func ReadNavFromFile(filename string) ([]Nav, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	navs := make([]Nav, 0)

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&navs)
	if err != nil {
		return nil, err
	}

	return navs, nil
}

func WriteNavToFile(filename string, navs []Nav) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // 设置缩进以便于阅读
	err = encoder.Encode(navs)
	if err != nil {
		return err
	}

	return nil
}
