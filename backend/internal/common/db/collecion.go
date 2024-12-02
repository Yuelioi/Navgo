package db

import (
	"backend/internal/common/cache"
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
	catIDs        map[string]struct{} // 用于验证数据库差异
}

func NewCollectionManager(root string) *CollectionManager {
	return &CollectionManager{
		Root:          root,
		Metas:         make([]*types.Group, 0, 1000),
		catMetas:      map[string]*CategoryMeta{},
		collectionIDs: map[string]struct{}{},
		catIDs:        map[string]struct{}{},
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

			m.catMetas[relativeDirPath] = &CategoryMeta{
				CID:   relativeDirPath,
				Title: meta.Category.Title,
			}

			if depth == 1 {
				// 拓展分类内容
				meta.Category.CID = relativeDirPath
				meta.Category.Path = relativeDirPath
				meta.Category.Depth = depth
				meta.Category.FullTitle = meta.Category.Title
				m.catIDs[meta.Category.Path] = struct{}{}
			} else {
				// 正常导航级别

				parentTitle := ""
				parentFolderName := relativeDirPathList[0]

				cm, ok := m.catMetas[parentFolderName]
				if ok {
					parentTitle = cm.Title
				}

				// 跳过当前没有收集网站的分类
				if meta.Collections == nil {
					return nil
				}

				// 补充目录参数
				cat := meta.Category
				cat.CID = relativeDirPathList[1]
				cat.Depth = depth
				cat.Path = relativeDirPath
				cat.FullTitle = parentTitle + "/" + meta.Category.Title
				meta.Category = cat

				m.catIDs[meta.Category.Path] = struct{}{}

				if info.Name() == "comic" {
					print(1)
				}

				for i, c := range meta.Collections {

					rawUrl, err := url.Parse(c.Link)
					if err != nil {
						return err
					}

					collectionCid := rawUrl.Host
					collectionPath := filepath.Join(relativeDirPath, collectionCid)

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

	storage := make(map[string]*types.Group, 0)

	// 储存本地缓存数据
	for _, mp := range m.Metas {
		storage[mp.Category.Path] = mp
	}

	c := cache.Manager.GetController(cache.LocalCacheID)
	c.Preload(storage)

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
	var allCollectionIDs []string
	var toDeleteCollectionIDs []string

	var allCationIDs []string
	var toDeleteCatIDs []string

	DB.Model(&types.Collection{}).Pluck("path", &allCollectionIDs)
	DB.Model(&types.Category{}).Pluck("path", &allCationIDs)

	for _, v := range allCollectionIDs {
		if _, ok := m.collectionIDs[v]; !ok {
			toDeleteCollectionIDs = append(toDeleteCollectionIDs, v)
		}
	}
	for _, v := range allCationIDs {
		if _, ok := m.collectionIDs[v]; !ok {
			toDeleteCatIDs = append(toDeleteCatIDs, v)
		}
	}

	// 删除找到的 ID 对应的记录
	if len(toDeleteCollectionIDs) > 0 {
		result := DB.Where("path IN (?)", toDeleteCollectionIDs).Delete(&types.Collection{})
		if result.Error != nil {
			return fmt.Errorf("failed to delete records: %v", result.Error)
		}
		fmt.Printf("Deleted %d records\n", result.RowsAffected)
	}

	if len(toDeleteCatIDs) > 0 {
		result := DB.Where("path IN (?)", toDeleteCatIDs).Delete(&types.Collection{})
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
