package main

import (
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type Collection struct {
	CID         string   `json:"cid,optional" gorm:"column:cid"`
	Title       string   `json:"title"`
	Link        string   `json:"link"`
	Order       int      `json:"order,optional" gorm:"column:order"`
	Path        string   `json:"path,optional" gorm:"column:path;unique"`
	Proxy       bool     `json:"proxy,optional" gorm:"column:proxy"`
	Description string   `json:"description,optional"`
	Thumbnail   string   `json:"thumbnail,optional"`
	Tags        []string `json:"tags,optional" gorm:"type:json"`
	View        int      `json:"view,optional"`
}

func query(link string) (*Collection, error) {

	URL, err := url.Parse(link)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// 检查响应状态码
	if res.StatusCode != 200 {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	title := doc.Find("title").Text()
	description := doc.Find("meta[name='description']").AttrOr("content", "")

	return &Collection{
		CID:         URL.Host,
		Title:       title,
		Link:        link,
		Proxy:       false,
		Description: description,
	}, nil
}

func main() {
	link := "https://www.nyadm.net/play/4469-1-11.html"
	query(link)
}
