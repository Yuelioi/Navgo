package cache

import (
	"backend/internal/common/constants"
	"backend/internal/types"
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var InReviewCacheID = "InReviewCache"

type InReviewCache struct {
	storage      *types.Group
	inReviewMeta string
}

func NewInReviewCache() *InReviewCache {
	return &InReviewCache{}
}

func (i *InReviewCache) ID() string {
	return InReviewCacheID
}

func (i *InReviewCache) Add(id string, item any) error {
	collection, ok := item.(*types.Collection)
	if !ok {
		return errors.New("item is not a collection")
	}

	i.storage.Collections = append(i.storage.Collections, collection)

	return i.AfterModify()
}

func (i *InReviewCache) AfterModify() error {
	data, err := yaml.Marshal(i.storage)
	if err != nil {
		return err
	}

	return os.WriteFile(i.inReviewMeta, data, os.ModePerm)
}

func (i *InReviewCache) Exists(id string) bool {
	for _, c := range i.storage.Collections {
		if c.CID == id {
			return true
		}
	}
	return false
}

// Preload implements Manager.
func (i *InReviewCache) Preload() error {
	inReviewDir := filepath.Join(constants.ConfInst.Resource.Collections, "_待审核")
	inReviewMeta := filepath.Join(inReviewDir, constants.ConfInst.Resource.MetaFile)
	i.inReviewMeta = inReviewMeta

	yamlFile, err := os.ReadFile(inReviewMeta)
	if err != nil {
		return err
	}

	meta := &types.Group{}
	err = yaml.Unmarshal(yamlFile, meta)
	if err != nil {
		return err
	}
	i.storage = meta
	return nil
}

// Remove implements Manager.
func (i *InReviewCache) Remove(id string) error {
	panic("unimplemented")
}
