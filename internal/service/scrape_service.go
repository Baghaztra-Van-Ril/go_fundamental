package service

import (
	"fmt"
	"go_fundamental/internal/scraper"
	"sync"
)

// Fungsi untuk scrape banyak URL secara paralel
func ScrapeMultiple(urls []string) []string {
	var wg sync.WaitGroup
	ch := make(chan string)
	
	// Jalankan setiap URL dalam goroutine
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			content, err := scraper.FetchContent(url)
			if err != nil {
				ch <- fmt.Sprintf("Error scraping %s: %v", url, err)
			} else {
				ch <- fmt.Sprintf("Content from %s: %s", url, content[:100]) // Tampilkan 100 karakter pertama
			}
		}(url)
	}

	// Tunggu semua goroutine selesai
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Ambil hasil dari channel
	var results []string
	for result := range ch {
		results = append(results, result)
	}

	return results
}
