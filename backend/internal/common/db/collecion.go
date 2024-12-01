package db

import (
	"backend/internal/common/constants"
	"backend/internal/types"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
	"gorm.io/gorm/clause"
)

type CategoryMeta struct {
	CID   string
	Title string
}

type CollectionManager struct {
	Root          string
	Metas         []*types.Group
	catMetas      map[string]*CategoryMeta
	collectionIDs map[string]struct{} // 用于验证数据库差异
}

func NewCollectionManager(root string) *CollectionManager {
	return &CollectionManager{
		Root:          root,
		Metas:         make([]*types.Group, 0, 1000),
		catMetas:      map[string]*CategoryMeta{},
		collectionIDs: map[string]struct{}{},
	}
}

func (m *CollectionManager) Read() error {
	filepath.WalkDir(m.Root, func(path string, info fs.DirEntry, err error) error {

		if info.IsDir() {
			if path == m.Root {
				return nil
			}
			if strings.HasPrefix(info.Name(), "_") {
				return filepath.SkipDir
			}
			return nil
		}

		relativeDirPath, _ := filepath.Rel(m.Root, filepath.Dir(path))
		relativeDirPathList := strings.Split(relativeDirPath, string(filepath.Separator))
		depth := len(relativeDirPathList)

		if info.Name() == constants.ConfInst.Resource.MetaFile {
			meta, err := readCollectionsMeta(path)
			if err != nil {
				fmt.Printf("meta %s err: %v\n", path, err.Error())
				return err
			}

			if depth == 1 {
				// 二级分类级别, 记录一下cid 并填写路径

				m.catMetas[relativeDirPath] = &CategoryMeta{
					CID:   meta.Category.CID,
					Title: meta.Category.Title,
				}
				meta.Category.Path = meta.Category.CID
				meta.Category.Depth = depth
			} else {
				// 正常导航级别

				cats := make([]string, 0)

				parentTitle := ""

				for _, id := range relativeDirPathList {
					if v, ok := m.catMetas[id]; ok {
						cats = append(cats, v.CID)
						parentTitle = v.Title
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
				cat.Path = filepath.Join(cats...)
				cat.FullTitle = parentTitle + "/" + meta.Category.Title

				meta.Category = cat

				for i, c := range meta.Collections {

					rawUrl, err := url.Parse(c.Link)
					if err != nil {
						return err
					}

					collectionCid := rawUrl.Host
					collectionPath := filepath.Join(append(cats, collectionCid)...)

					// 补充页面信息, 使用主机名作为ID
					meta.Collections[i].CID = collectionCid
					meta.Collections[i].Category = cat
					meta.Collections[i].Path = collectionPath
					meta.Collections[i].Order = i + 1

					// 储存存在的页面
					m.collectionIDs[collectionPath] = struct{}{}
				}
			}
			m.Metas = append(m.Metas, meta)

		}
		return nil
	})

	return nil
}

func (m *CollectionManager) Build() error {
	cats := make([]*types.Category, 0)
	collections := make([]*types.Collection, 0)
	for _, element := range m.Metas {
		cats = append(cats, element.Category)
		collections = append(collections, element.Collections...)
	}

	DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "path"}},
		UpdateAll: true,
	}).CreateInBatches(cats, 20)

	DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "path"}},
		UpdateAll: true,
	}).CreateInBatches(collections, 100)

	return nil
}

func (m *CollectionManager) Reduce() error {
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
			return fmt.Errorf("failed to delete records: %v", result.Error)
		}
		fmt.Printf("Deleted %d records\n", result.RowsAffected)
	}
	return nil
}

func readCollectionsMeta(path string) (*types.Group, error) {
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
