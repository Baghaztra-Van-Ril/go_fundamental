package service

import (
	"fmt"
	"go_fundamental/internal/scraper"
	"golang.org/x/net/html"
	"strings"
	"sync"
)

// Fungsi untuk mengambil judul dari laman html yang didapatkan
func ExtractTitle(htmlContent string) (string, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return "", fmt.Errorf("error parsing HTML: %v", err)
	}

	var title string
	var found bool

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		// Jika title sudah ditemukan (dari perulangan sebelumnya), fungsi tidak dilanjutkan
		if found {
			return
		}

		// Jika title ditemukan, simpan, lalu hentikan fungsi ini
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			title = n.FirstChild.Data
			found = true
			return
		}

		// lompat ke tag berikutnya (next sibling)
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	// Jalankan rekursi mulai dari doc (hasil scrap awal)
	traverse(doc)
	return title, nil
}


// Penerapan Goroutine dan Channel
// func ScrapeMultiple(urls []string) []string {
// 	// Untuk menunggu semuanya
// 	var wg sync.WaitGroup
// 	// mengirimkan hasil ke main
// 	ch := make(chan string)

// 	// Mulai cari setiap URL
// 	for _, url := range urls {
// 		wg.Add(1)
// 		go func(url string) {
// 			defer wg.Done()
// 			content, err := scraper.FetchContent(url)
// 			if err != nil {
// 				ch <- fmt.Sprintf("Error scraping %s: %v", url, err)
// 			} else {
// 				title, _ := ExtractTitle(content)
// 				ch <- fmt.Sprintf("URL: %s\nTitle: %s\n", url, title)
// 			}
// 		}(url)
// 	}

// 	// Tunggu semuanya selesai
// 	go func() {
// 		wg.Wait()
// 		close(ch)
// 	}()

// 	// Ambil hasilnya dengan channel
// 	var results []string
// 	for result := range ch {
// 		results = append(results, result)
// 	}

// 	// Serahkan hasilnya
// 	return results
// }

// Fungsi untuk scraping URL dengan worker pool
func ScrapeMultiple(urls []string) []string {
	var wg sync.WaitGroup
	taskChan := make(chan string)   // Untuk mengirim URL
	resultChan := make(chan string) // Untuk menerima hasil scraping

	numWorkers := 5 // Jumlah maksimal worker yang berjalan

	// Buat worker goroutine
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range taskChan {
				content, err := scraper.FetchContent(url)
				if err != nil {
					resultChan <- fmt.Sprintf("Error scraping %s: %v", url, err)
				} else {
					title, _ := ExtractTitle(content) // Ambil title dan metadata
					resultChan <- fmt.Sprintf("URL: %s\nTitle: %s\n", url, title)
				}
			}
		}()
	}

	// Kirim URL ke taskChan
	go func() {
		for _, url := range urls {
			taskChan <- url
		}
		close(taskChan) // Setelah semua URL dikirim, tutup channel
	}()

	// Tunggu hingga semua worker selesai
	go func() {
		wg.Wait()
		close(resultChan) // Tutup resultChan setelah semua selesai
	}()

	// Ambil semua hasil dari resultChan
	var results []string
	for result := range resultChan {
		results = append(results, result)
	}

	return results
}
