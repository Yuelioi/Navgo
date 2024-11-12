package main

import (
	"fmt"
	"net/url"
)

func main() {
	template := `  - cid: %s
    title: %s
    link: %s
    description: ""
    country: "%s"
    proxy: %t
    thumbnail:
    view: 0`

	link := "https://www.nyadm.net/"
	resp, err := generate(link)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	var proxy bool
	err = ping(link)
	if err != nil {
		proxy = true
	}

	pureUrl, _ := url.Parse(link)
	cid := pureUrl.Host

	data := fmt.Sprintf(template, cid, resp.SiteTitle, link, resp.Address.Country, proxy)
	fmt.Println(data)
}
