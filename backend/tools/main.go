package main

import (
	"fmt"
)

func main() {
	template := `  -title: %s
    link: %s
    description: %s
    proxy: %t`

	link := "https://tools.miku.ac/"
	meta, err := queryMeta(link)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	var proxy bool
	err = ping(link)
	if err != nil {
		proxy = true
	}

	data := fmt.Sprintf(template, meta.Title, link, meta.Description, proxy)
	fmt.Println(data)
}
