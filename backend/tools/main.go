package main

import "fmt"

func main() {
	template := `  - title: %s
    link: %s
    description: %s`

	link := "https://www.nmod.net/"
	meta, err := queryMeta(link)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	err = ping(link)
	if err != nil {
		return
	}

	data := fmt.Sprintf(template, meta.Title, link, meta.Description)
	fmt.Println(data)
}
