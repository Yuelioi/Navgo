package db

import (
	"backend/internal/types"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
	"gorm.io/gorm/clause"
)

type CommentManager struct {
	Root         string // 储存的meta文件路径
	commentsMeta []*types.Comment
	commentIDs   map[string]struct{} // 用于验证数据库差异
}

func NewCommentManager(root string) *CommentManager {
	return &CommentManager{
		Root:       root,
		commentIDs: map[string]struct{}{},
	}
}

func (m *CommentManager) init() error {
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

		datas := &types.CommentsResponse{
			Comments: []*types.Comment{{
				Nickname: "system",
				Content:  "test",
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

func (m *CommentManager) Read() error {

	// 1.初始化判断
	err := m.init()
	if err != nil {
		return err
	}

	// 2. 读取meta文件
	meta := &types.CommentsResponse{}
	yamlFile, err := os.ReadFile(m.Root)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, meta)
	if err != nil {
		return err
	}

	for _, a := range meta.Comments {
		m.commentIDs[a.Content] = struct{}{}
	}
	m.commentsMeta = meta.Comments
	return nil
}

func (m *CommentManager) Build() error {
	DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "content"}},
		UpdateAll: true,
	}).CreateInBatches(m.commentsMeta, 10)
	return nil
}

func (m *CommentManager) Reduce() error {
	var allIDs []string
	DB.Model(&types.Comment{}).Pluck("content", &allIDs)

	var toDeleteIDs []string

	for _, v := range allIDs {
		if _, ok := m.commentIDs[v]; !ok {
			toDeleteIDs = append(toDeleteIDs, v)
		}
	}

	// 删除找到的 ID 对应的记录
	if len(toDeleteIDs) > 0 {
		result := DB.Where("content in (?)", toDeleteIDs).Delete(&types.Comment{})
		if result.Error != nil {
			return fmt.Errorf("failed to delete records: %v", result.Error)
		}
		fmt.Printf("Deleted %d records\n", result.RowsAffected)
	}
	return nil
}
