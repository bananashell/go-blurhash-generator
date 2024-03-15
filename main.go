package main

import (
	"fmt"
	"image/jpeg"
	"net/http"

	"github.com/bbrks/go-blurhash"
)

func main() {
	urls := []string{
		"url to image here",
	}

	for _, url := range urls {
		hash, err := generateBlurhash(url)
		if err != nil {
			println(url, "\t", "Error")
			continue
		}
		println(url, "\t", hash)
	}
}

func generateBlurhash(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("unable to fetch image")
	}

	defer response.Body.Close()
	if response.StatusCode != 200 {
		return "", fmt.Errorf("error: %d", response.StatusCode)
	}

	img, err := jpeg.Decode(response.Body)
	if err != nil {
		return "", err
	}

	hash, err := blurhash.Encode(4, 4, img)
	if err != nil {
		return "", err
	}

	return hash, nil
}
