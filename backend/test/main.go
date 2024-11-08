package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	CID   string `json:"cid,optional" gorm:"column:cid;unique"`
	Title string `json:"title"`
	Order int    `json:"order,optional"`
}

type Collection struct {
	CID         string   `json:"cid,optional" gorm:"column:cid;unique"`
	Title       string   `json:"title"`
	Link        string   `json:"link"`
	CategoryID  string   `json:"-" gorm:"column:category_id;index"` // 外键字段
	Category    Category `json:"category" gorm:"foreignKey:CategoryID;references:CID"`
	Description string   `json:"description,optional"`
	Country     string   `json:"country,optional"`
	Thumbnail   string   `json:"thumbnail,optional"`
	// Tags        []string `json:"tags,optional" gorm:"type:json"`
	View int `json:"view,optional"`
}

type Meta struct {
	Category    Category     `json:"-" yaml:",inline"` // 使用 inline 标签将 Category 字段内联
	Collections []Collection `json:"collections"`
}

func main() {
	rootRel := "..\\resource\\collections"
	root, _ := filepath.Abs(rootRel)
	if _, err := os.Stat(root); os.IsNotExist(err) {
		return
	}
	metas := readDir(root)

	cats := make([]Category, 0)
	collections := make([]Collection, 0)
	for _, element := range metas {
		cats = append(cats, element.Category)
		collections = append(collections, element.Collections...)
	}

	DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "cid"}},
		UpdateAll: true,
	}).CreateInBatches(collections, 100)

	DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "cid"}},
		UpdateAll: true,
	}).CreateInBatches(cats, 100)

}

var DB *gorm.DB
var err error

func init() {
	if _, err = os.Stat("./core.db"); os.IsNotExist(err) {
		dir := filepath.Dir("./core.db")

		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("failed to create directory: %v", err)
		}

		file, err := os.Create("./core.db")
		if err != nil {
			panic("failed to create file")
		}
		defer file.Close()
	}

	DB, err = gorm.Open(sqlite.Open("./core.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = DB.AutoMigrate(&Category{}, &Collection{})

}

func readDir(root string) []*Meta {
	metas := make([]*Meta, 0, 1000)
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if info.Name() == "meta.yaml" {
			meta, err := readMeta(path)

			for i, c := range meta.Collections {
				c.Category = meta.Category
				meta.Collections[i] = c
			}

			if err != nil {
				fmt.Printf("read %s err: %v\n", path, err.Error())
				return err
			}
			metas = append(metas, meta)
		}
		return nil
	})
	return metas

}

func readMeta(path string) (*Meta, error) {
	meta := &Meta{}
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, meta)
	if err != nil {
		return nil, err
	}
	return meta, nil
}
