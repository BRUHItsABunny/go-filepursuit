package main

import (
	"fmt"
	go_filepursuit "github.com/BRUHItsABunny/go-filepursuit"
)

func main() {
	// Initialize client
	client := go_filepursuit.GetFilePursuitClient()
	// Search
	fmt.Println(client.Search("bunny", go_filepursuit.FileTypeVideo, go_filepursuit.SortDateDescending, 0))
	// Discover
	fmt.Println(client.Discover("http://1.158.171.199:8080/legacy/get/DOC/100/Calibre_Library", 0))
	// News
	fmt.Println(client.News())
	// Suggest
	fmt.Println(client.Suggest("bunny"))
	// Submit
	fmt.Println(client.Submit("Bunny Gang", "bunny@bunny.bunny", "bunny.bunbunbun", "https://bunny.bunbun/bunny.bunbunbun"))
}
