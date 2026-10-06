package main

import (
	"net/url"
)

type PageData struct {
	URL            string   `json:"url"`
	Heading        string   `json:"heading"`
	FirstParagraph string   `json:"first_paragraph"`
	OutgoingLinks  []string `json:"outgoing_links"`
	ImageURLs      []string `json:"image_urls"`
}

func extractPageData(html, pageURL string) PageData {
	baseURL, _ := url.Parse(pageURL)
	//fmt.Printf("baseURL: %s\n", baseURL.Host)
	heading := getHeadingFromHTML(html)
	firstParagraph := getFirstParagraphFromHTML(html)
	outgoingLinks, _ := getURLsFromHTML(html, baseURL)
	images, _ := getImagesFromHTML(html, baseURL)
	return PageData{
		URL:            pageURL,
		Heading:        heading,
		FirstParagraph: firstParagraph,
		OutgoingLinks:  outgoingLinks,
		ImageURLs:      images,
	}

	// blah := PageData{URL:https://crawler-test.com Heading:Test Title FirstParagraph:This is the first paragraph. OutgoingLinks:[https://crawler-test.com/link1] ImageURLs:[https://crawler-test.com/image1.jpg]}
	// blah := PageData{URL:https://crawler-test.com Heading:Test Title FirstParagraph:This is the first paragraph. OutgoingLinks:[ https://crawler-test.com/link1] ImageURLs:[ https://crawler-test.com/image1.jpg]}
}
