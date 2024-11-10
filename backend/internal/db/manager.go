package db

import (
	"backend/internal/global"
	"backend/internal/types"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
	"gorm.io/gorm/clause"
)

type Meta struct {
	Category    *types.Category     `json:"-" yaml:",inline"` // 使用 inline 标签将 Category 字段内联
	Collections []*types.Collection `json:"collections"`
}

// func (m *Meta) UnmarshalYAML(value *yaml.Node) error {
// 	type plainMeta struct {
// 		Category    *types.Category     `json:"-" yaml:",inline"`
// 		Collections []*types.Collection `json:"collections"`
// 	}

// 	var plain plainMeta
// 	if err := value.Decode(&plain); err != nil {
// 		return err
// 	}

// 	cat := &types.Category{
// 		CID:   plain.Category.CID,
// 		Title: plain.Category.Title,
// 		Order: plain.Category.Order,
// 		Path:  plain.Category.Path,
// 	}

// 	m.Category = cat

// 	m.Collections = make([]*types.Collection, len(plain.Collections))
// 	for i, c := range plain.Collections {
// 		c.Category = cat
// 		m.Collections[i] = c
// 	}

// 	return nil
// }

type Manager struct {
	Root          string
	Metas         []*Meta
	catIDs        map[string]string
	collectionIDs map[string]struct{}
}

func New(root string) *Manager {
	return &Manager{
		Root:          root,
		Metas:         make([]*Meta, 0, 1000),
		catIDs:        map[string]string{},
		collectionIDs: map[string]struct{}{},
	}
}

func (m *Manager) Read() {
	filepath.WalkDir(m.Root, func(path string, info fs.DirEntry, err error) error {
		if path == m.Root {
			return nil
		}

		if info.IsDir() {
			if strings.HasPrefix(info.Name(), "_") {
				return filepath.SkipDir
			}
			return nil
		}

		short, _ := strings.CutPrefix(path, m.Root)
		short, _ = strings.CutPrefix(short, "\\")
		relativeDirPath := filepath.Dir(short) // 父级相对路径

		depth := len(strings.Split(relativeDirPath, "\\"))

		if info.Name() == global.ConfInst.Collection.MetaFile {
			meta, err := readMeta(path)
			if err != nil {
				fmt.Printf("meta %s err: %v\n", path, err.Error())
				return err
			}

			if depth == 1 {
				// 分类级别
				m.catIDs[relativeDirPath] = meta.Category.CID
			} else {
				// 正常级别
				paths := strings.Split(relativeDirPath, "\\")
				cats := make([]string, 0)
				for _, id := range paths {
					if v, ok := m.catIDs[id]; ok {
						cats = append(cats, v)
					} else {
						cats = append(cats, meta.Category.CID)
					}
				}
				// 跳过当前没有收集网站的分类
				if meta.Collections == nil {
					return nil
				}

				// 补充目录的Path
				cat := meta.Category
				cat.Path = strings.Join(cats, "\\")
				meta.Category = cat

				for i, c := range meta.Collections {
					paths = append(cats, c.CID)
					path = strings.Join(paths, "\\")

					// 补充页面的目录与Path
					meta.Collections[i].Category = cat
					meta.Collections[i].Path = path

					// 储存存在的页面
					m.collectionIDs[path] = struct{}{}
				}
				m.Metas = append(m.Metas, meta)
			}

		}
		return nil
	})
}

func (m *Manager) Build() {
	root, _ := filepath.Abs(m.Root)
	if _, err := os.Stat(root); os.IsNotExist(err) {
		return
	}

	cats := make([]*types.Category, 0)
	collections := make([]*types.Collection, 0)
	for _, element := range m.Metas {
		cats = append(cats, element.Category)
		collections = append(collections, element.Collections...)
	}

	DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "path"}},
		UpdateAll: true,
	}).CreateInBatches(cats, 100)

	DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "path"}},
		UpdateAll: true,
	}).CreateInBatches(collections, 100)

}

func (m *Manager) Reduce() {
	var allIDs []string
	DB.Model(&types.Collection{}).Pluck("path", &allIDs)

	var toDeleteIDs []string

	for _, v := range allIDs {
		if _, ok := m.collectionIDs[v]; !ok {
			toDeleteIDs = append(toDeleteIDs, v)
		}
	}

	// 删除找到的 ID 对应的记录
	if len(toDeleteIDs) > 0 {
		result := DB.Where("path IN (?)", toDeleteIDs).Delete(&types.Collection{})
		if result.Error != nil {
			log.Fatalf("failed to delete records: %v", result.Error)
		}
		fmt.Printf("Deleted %d records\n", result.RowsAffected)
	}
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
