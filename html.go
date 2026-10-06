package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "BootCrawler/1.0")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode > 400 {
		return "", fmt.Errorf("Error status: %d %s", res.StatusCode, res.Status)
	}
	if !strings.Contains(res.Header.Get("content-type"), "text/html") {
		return "", errors.New("Content not HTML")
	}
	html, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errors.New("Cannot read HTML")
	}
	return string(html), nil
}
