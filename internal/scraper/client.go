package scraper

import (
	"fmt"
	"io"
	"net/http"
)

// Fungsi untuk ambil konten dari URL
func FetchContent(url string) (string, error) {
	// Kirim request GET ke URL
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code: %d", resp.StatusCode)
	}

	// Baca konten dari response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
