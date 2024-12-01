package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Downloader struct {
	baseUrl string
	client  *http.Client
	workDir string
	links   []string

	fails []string
}

func New() *Downloader {
	client, _ := getClient()
	workDir := "./downloader"

	if _, err := os.Stat(workDir); err != nil {
		if os.IsNotExist(err) {
			// 如果目录不存在，则创建它
			err = os.MkdirAll(workDir, os.ModePerm)
			if err != nil {
				log.Fatalf("创建文件夹失败: %v", err)
			}
			fmt.Println("目录创建成功:", workDir)
		} else {
			// 其他错误，打印错误信息
			log.Fatalf("检查目录失败: %v", err)
		}
	}

	return &Downloader{
		baseUrl: "https://wallhaven.cc/search?q=id%3A5&categories=110&purity=100&atleast=1600x900&topRange=1y&sorting=toplist&order=desc&ai_art_filter=1&page=",
		client:  client,
		workDir: workDir,
		links:   []string{},
	}
}

func getClient() (*http.Client, error) {
	proxyURL, err := url.Parse("http://127.0.0.1:10809")
	if err != nil {
		fmt.Println("Error parsing proxy URL:", err)
		return nil, err
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	return &http.Client{
		Transport: transport,
	}, nil
}

func sanitizeFileName(fileName string) string {
	// 替换非法字符
	re := regexp.MustCompile(`[\\/:*?"<>|]`)
	fileName = re.ReplaceAllString(fileName, "_")

	// 去除前导和尾随的点号及空格
	fileName = strings.Trim(fileName, " .")

	// 将连续的下划线替换为单个下划线
	fileName = regexp.MustCompile(`_+`).ReplaceAllString(fileName, "_")
	fileName = regexp.MustCompile(` +`).ReplaceAllString(fileName, " ")

	return fileName
}

func removeParenthesesContent(input string) string {
	re := regexp.MustCompile(`\([^)]*\)`) // 匹配括号及其中内容
	result := re.ReplaceAllString(input, "")
	return result
}

func (d *Downloader) start(start int, limit int) {

	var wg sync.WaitGroup

	for page := range limit {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			if err := d.fetchLinks(start + p); err != nil {
				log.Printf("Error fetching links for page %d: %v", start+p, err)
			}

		}(page)
	}
	wg.Wait()
	log.Println("页面列表获取成功")

	for _, link := range d.links {
		if err := d.fetchPage(link); err != nil {
			log.Printf("Error fetching page %s: %v", link, err)
		}
		time.Sleep(time.Second)

	}

	wg.Wait()

}

func (d *Downloader) NewDoc(link string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")

	resp, err := d.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (d *Downloader) downloadImg(url, filePath string) error {
	resp, err := d.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.Create(filePath)
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

func (d *Downloader) fetchPage(link string) error {

	doc, err := d.NewDoc(link)
	if err != nil {
		return err
	}

	tags := []string{}
	doc.Find("#tags a").Each(func(i int, s *goquery.Selection) {
		tags = append(tags, removeParenthesesContent(s.Text())+" ")
	})

	tagStr := "#wallpaper "
	title := ""
	count := 0

	for _, tag := range tags {

		count += len(tag)
		if count > 200 {
			break
		}
		tag = sanitizeFileName(tag)
		tagStr += "#" + tag + " "
	}

	matches := regexp.MustCompile(`/w/(.+)`).FindStringSubmatch(link)
	if len(matches) > 1 {
		title = matches[1]
	}

	fp := filepath.Join(d.workDir, sanitizeFileName(title+" "+tagStr)+".png")

	href, exists := doc.Find("#showcase img").Attr("src")
	if !exists {
		fmt.Printf("%s has no href attribute\n", link)
		d.fails = append(d.fails, link)
		return nil
	}
	return d.downloadImg(href, fp)
}

func (d *Downloader) fetchLinks(page int) error {
	doc, err := d.NewDoc(d.baseUrl + strconv.Itoa(page))
	if err != nil {
		return err
	}

	doc.Find("#thumbs .thumb-listing-page ul > li:nth-child(n) > figure > a").Each(func(i int, s *goquery.Selection) {
		// 获取 href 属性
		href, exists := s.Attr("href")
		if !exists {
			fmt.Printf("Element %d has no href attribute\n", i)
			return
		}
		d.links = append(d.links, href)
	})

	return nil
}
