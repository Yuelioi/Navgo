package db

import (
	"backend/internal/types"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
	"gorm.io/gorm/clause"
)

type AnnounceManager struct {
	Root          string // 储存的meta文件路径
	announcesMeta []*types.Announce
	announceIDs   map[string]struct{} // 用于验证数据库差异
}

func NewAnnounceManager(root string) *AnnounceManager {
	return &AnnounceManager{
		Root:        root,
		announceIDs: map[string]struct{}{},
	}
}

func (m *AnnounceManager) init() error {
	root := filepath.Dir(m.Root)

	if _, err := os.Stat(root); os.IsNotExist(err) {
		err := os.MkdirAll(root, os.ModePerm)
		if err != nil {
			return err
		}
	}

	if _, err := os.Stat(m.Root); os.IsNotExist(err) {
		f, err := os.OpenFile(m.Root, os.O_CREATE, os.ModePerm)
		if err != nil {
			return err
		}
		defer f.Close()

		datas := &types.AnnouncesData{
			Announces: []*types.Announce{{
				Title:   "第一个公告",
				Content: "欢迎使用公告",
				Date:    "2024/11/16",
			}},
		}

		data, err := yaml.Marshal(datas)
		if err != nil {
			return err
		}
		_, err = f.Write(data)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *AnnounceManager) Read() error {

	// 1.初始化判断
	err := m.init()
	if err != nil {
		return err
	}

	// 2. 读取meta文件
	meta := &types.AnnouncesData{}
	yamlFile, err := os.ReadFile(m.Root)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, meta)
	if err != nil {
		return err
	}

	for _, a := range meta.Announces {
		m.announceIDs[a.Content] = struct{}{}
	}
	m.announcesMeta = meta.Announces
	return nil
}

func (m *AnnounceManager) Build() error {
	DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "content"}},
		UpdateAll: true,
	}).Create(m.announcesMeta)
	return nil
}

func (m *AnnounceManager) Reduce() error {
	var allIDs []string
	DB.Model(&types.Announce{}).Pluck("content", &allIDs)

	var toDeleteIDs []string

	for _, v := range allIDs {
		if _, ok := m.announceIDs[v]; !ok {
			toDeleteIDs = append(toDeleteIDs, v)
		}
	}

	// 删除找到的 ID 对应的记录
	if len(toDeleteIDs) > 0 {
		result := DB.Where("content in (?)", toDeleteIDs).Delete(&types.Announce{})
		if result.Error != nil {
			return fmt.Errorf("failed to delete records: %v", result.Error)
		}
		fmt.Printf("Deleted %d records\n", result.RowsAffected)
	}
	return nil
}
