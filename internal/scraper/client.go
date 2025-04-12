package scraper

import (
	// "fmt"
	"io/ioutil"
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

	// Baca konten dari response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
