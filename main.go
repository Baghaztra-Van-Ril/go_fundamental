package main

import (
	"fmt"
	"go_fundamental/internal/service"
)

func main() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://news.ycombinator.com",
	}

	// Panggil fungsi untuk scrape banyak URL
	results := service.ScrapeMultiple(urls)

	// Tampilkan hasilnya
	for _, result := range results {
		fmt.Println(result)
	}
}
