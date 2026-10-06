package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	a := doc.Find("a")
	var links []string
	a.Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		url, _ := url.Parse(link)
		if url.Host == "" {
			url.Host = baseURL.Host
		}
		if !strings.HasPrefix(url.Path, "/") {
			url.Path = "/" + url.Path
		}
		fullURL, _ := normalizeURL(url.Host + url.Path)
		links = append(links, strings.TrimSpace(baseURL.Scheme+"://"+fullURL))
	})
	return links, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	a := doc.Find("img")
	var images []string
	a.Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("src")
		url, _ := url.Parse(link)
		if url.Host == "" {
			url.Host = baseURL.Host
		}
		if !strings.HasPrefix(url.Path, "/") {
			url.Path = "/" + url.Path
		}
		fullURL, _ := normalizeURL(url.Host + url.Path)
		images = append(images, strings.TrimSpace(baseURL.Scheme+"://"+fullURL))
	})
	return images, nil
}
