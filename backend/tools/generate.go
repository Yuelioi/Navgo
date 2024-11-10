package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type Response struct {
	Code            int     `json:"code"`
	Cache           bool    `json:"cache"`
	Lang            string  `json:"lang"`
	Address         Address `json:"address"`
	PrimaryIP       string  `json:"primary_ip"`
	SiteTitle       string  `json:"site_title"`
	SiteKeywords    string  `json:"site_keywords"`
	SiteDescription string  `json:"site_description"`
	FaviconIcon     string  `json:"favicon_icon"`
	ServerID        int     `json:"server_id"`
	URL             string  `json:"url"`
}

// Address 结构体用于解析地址信息
type Address struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func generate(link string) (*Response, error) {

	url := "https://apiv2.iotheme.cn/webinfo/get.php"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("url", link)
	_ = writer.WriteField("type", "json")
	_ = writer.WriteField("key", "WyJNakF5TVRJMk16Z3hNalkzIiwiYTIxMVNHTktOR05LVDB0RVJscDRWMDVFTkRWaGRYVlBTR0o0UjJ4cSJd")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("referer", "https://nav.yuelili.com/")
	req.Header.Add("Cookie", "PHPSESSID=0ufc4ll9vsk34r4cmmhi4ftfr3")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resp, err
}
