package main

import (
	"fmt"
	"sort"
)

type Item struct {
	Category Category
	Groups   []*Group `json:"groups"`
}

type Group struct {
	Category    *Category     `json:"-" yaml:",inline"` // 使用 inline 标签将 Category 字段内联
	Collections []*Collection `json:"collections"`
}

type Category struct {
	Cid   string `json:"cid,optional" gorm:"column:cid"`
	Depth int    `json:"depth,optional"`
	Title string `json:"title"`
	Order int    `json:"order,optional"`
	Path  string `json:"path,optional" gorm:"column:path;unique"`
}

type Collection struct {
	CID         string    `json:"cid,optional" gorm:"column:cid"`
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	Path        string    `json:"path,optional" gorm:"column:path;unique"`
	CategoryID  string    `json:"-" gorm:"column:category_id;index"` // 外键字段
	Category    *Category `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	Description string    `json:"description,optional"`
	Country     string    `json:"country,optional"`
	Thumbnail   string    `json:"thumbnail,optional"`
	Tags        []string  `json:"tags,optional" gorm:"type:json"`
	View        int       `json:"view,optional"`
}

type ByCategoryOrder []Item

func (a ByCategoryOrder) Len() int           { return len(a) }
func (a ByCategoryOrder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCategoryOrder) Less(i, j int) bool { return a[i].Category.Order < a[j].Category.Order }

func main() {
	items := []Item{
		{
			Category: Category{
				Cid:   "cg",
				Depth: 1,
				Title: "后期",
				Order: 1,
				Path:  "cg",
			},
			Groups: []*Group{},
		},

		{
			Category: Category{
				Cid:   "web",
				Depth: 1,
				Title: "前端",
				Order: 3,
				Path:  "web",
			},
			Groups: []*Group{},
		},
		{
			Category: Category{
				Cid:   "app",
				Depth: 1,
				Title: "应用",
				Order: 2,
				Path:  "app",
			},
			Groups: []*Group{},
		},
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Category.Order < items[j].Category.Order
	})

	// 打印排序后的结果
	for _, item := range items {
		fmt.Printf("Title: %s, Order: %d\n", item.Category.Title, item.Category.Order)
	}
}
