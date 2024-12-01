package cache

import (
	"backend/internal/common/constants"
	"backend/internal/types"
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"gopkg.in/yaml.v3"
)

/*
储存本地必要数据
*/

var localCacheID = "localCache"

type localCache struct {
	storage map[string]*types.Group
}

func NewLocalCache() *localCache {
	return &localCache{}
}

func (i *localCache) ID() string {
	return localCacheID
}

func (i *localCache) Add(id string, item any) error {
	collection, ok := item.(*types.Collection)
	if !ok {
		return errors.New("item is not a collection")
	}

	// 会preload数据, 因此不会出现没有的情况(仅在目录不变动的情况下)
	if group, ok := i.storage[id]; ok {
		group.Collections = append(group.Collections, collection)
	} else {
		group := &types.Group{
			Category: &types.Category{
				CID:       localCacheID,
				Title:     "",
				FullTitle: "",
				Path:      "",
			},
			Collections: []*types.Collection{},
		}
		group.Collections = append(group.Collections, collection)
	}

	return i.AfterModify(id)
}

func (i *localCache) AfterModify(id string) error {
	if group, ok := i.storage[id]; ok {
		data, err := yaml.Marshal(group)
		if err != nil {
			return err
		}

		folders := strings.Split(id, string(os.PathSeparator))
		folders = slices.Concat([]string{constants.ConfInst.Resource.Collections}, folders, []string{constants.ConfInst.Resource.MetaFile})
		metaFile := filepath.Join(folders...)
		// 文件不存在
		if _, err := os.Stat(metaFile); err == os.ErrNotExist {
			return err
		}

		return os.WriteFile(metaFile, data, os.ModePerm)
	}

	return nil

}

func (i *localCache) Exists(id string) bool {
	if _, ok := i.storage[id]; ok {
		return true
	}

	return false
}

func (i *localCache) Get(id string) (any, error) {
	if item, ok := i.storage[id]; ok {
		return item, nil
	}

	return nil, errors.New("not found")
}

// Preload implements Manager.
func (i *localCache) Preload() error {

	return nil
}

// Remove implements Manager.
func (i *localCache) Remove(id string) error {
	panic("unimplemented")
}
