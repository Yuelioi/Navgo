package types

/*
一些模块格式化 序列号自定义
*/

type category struct {
	CID   string `yaml:"cid,omitempty"`
	Title string `yaml:"title,omitempty"`
	Order int    `yaml:"order,omitempty"`
}

type announce struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

type comment struct {
	Nickname string `json:"nickname"`
	Content  string `json:"content" gorm:"column:content;unique"`
	Date     string `json:"date"`
}

type collection struct {
	Title       string   `json:"title"`
	Link        string   `json:"link"`
	Description string   `json:"description"`
	Proxy       bool     `json:"proxy,omitempty" yaml:",omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

//  暂时没用到
func (c Category) MarshalYAML() (interface{}, error) {
	return struct {
		CID   string `yaml:"cid,omitempty"`
		Title string `yaml:"title,omitempty"`
		Order int    `yaml:"order,omitempty"`
	}{
		CID:   c.CID,
		Title: c.Title,
		Order: c.Order,
	}, nil
}

// 写入到待审核文件
func (g Group) MarshalYAML() (interface{}, error) {
	collections := make([]*collection, len(g.Collections))

	for i, c := range g.Collections {
		collections[i] = &collection{
			Title:       c.Title,
			Link:        c.Link,
			Description: c.Description,
			Proxy:       c.Proxy,
			Tags:        c.Tags,
		}
	}

	return struct {
		Category    *category `json:"category" yaml:",inline"`
		Collections []*collection
	}{
		Category: &category{
			CID:   g.Category.CID,
			Title: g.Category.Title,
			Order: g.Category.Order,
		},
		Collections: collections,
	}, nil
}

func (g AnnouncesData) MarshalYAML() (interface{}, error) {
	announces := make([]*announce, len(g.Announces))

	for i, c := range g.Announces {
		announces[i] = &announce{
			Title:   c.Title,
			Content: c.Content,
			Date:    c.Date,
		}
	}

	return struct {
		Announces []*announce
	}{
		Announces: announces,
	}, nil
}

func (g CommentsResponse) MarshalYAML() (interface{}, error) {
	datas := make([]*comment, len(g.Comments))

	for i, c := range g.Comments {
		datas[i] = &comment{
			Nickname: c.Nickname,
			Content:  c.Content,
			Date:     c.Date,
		}
	}

	return struct {
		Comments []*comment
	}{
		Comments: datas,
	}, nil
}

func (g Comment) MarshalYAML() (interface{}, error) {

	return comment{
		Nickname: g.Nickname,
		Content:  g.Content,
		Date:     g.Date,
	}, nil
}
