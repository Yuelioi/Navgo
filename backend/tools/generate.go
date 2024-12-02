package main

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type Response struct {
	Link        string `json:"link"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func queryMeta(link string) (*Response, error) {
	_, err := url.Parse(link)
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
		return nil, errors.New("请求失败")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	title := doc.Find("title").Text()
	description := doc.Find("meta[name='description']").AttrOr("content", "")

	resp := &Response{
		Title:       title,
		Link:        link,
		Description: description,
	}

	// err = downloadFavicon(link)
	// if err != nil {
	// 	return nil, err
	// }

	return resp, err
}

func downloadFavicon(link string) error {
	resp, err := http.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	host, err := urlToHostName(link)
	if err != nil {
		return err
	}

	f, err := os.Create(host + ".png")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
