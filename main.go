package main

import (
	"fmt"
	"go_fundamental/internal/service"
)

func main() {
	urls := []string{
		"https://simarasok.umkm-pnp.com/",
		"https://kanaria.codes/",
		"https://www.tempo.co/",
	}

	// Panggil fungsi untuk scrape setaip URL
	results := service.ScrapeMultiple(urls)

	// Tampilkan hasilnya
	for i, result := range results {
		fmt.Printf("URL %d: %s\n", i+1, result)
	}	
}
