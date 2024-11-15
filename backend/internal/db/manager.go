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

type Manager struct {
	Root          string
	Metas         []*types.Group
	catIDs        map[string]string
	collectionIDs map[string]struct{} // 用于验证数据库差异
}

func New(root string) *Manager {
	return &Manager{
		Root:          root,
		Metas:         make([]*types.Group, 0, 1000),
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
				// 二级分类级别, 记录一下cid 并填写路径
				m.catIDs[relativeDirPath] = meta.Category.CID
				meta.Category.Path = meta.Category.CID
				meta.Category.Depth = depth
			} else {
				// 正常导航级别
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
				cat.Depth = depth
				cat.Path = strings.Join(cats, "\\")
				meta.Category = cat

				for i, c := range meta.Collections {
					paths = append(cats, c.CID)
					path = strings.Join(paths, "\\")

					// 补充页面的目录与Path
					meta.Collections[i].Category = cat
					meta.Collections[i].Path = path
					meta.Collections[i].Order = i + 1

					// 储存存在的页面
					m.collectionIDs[path] = struct{}{}
				}
			}
			m.Metas = append(m.Metas, meta)

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
	}).CreateInBatches(cats, 10)

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

func readMeta(path string) (*types.Group, error) {
	meta := &types.Group{}
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
