package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	heading := doc.Find("h1")

	if heading.Length() != 0 {
		return heading.Text()
	}
	heading = doc.Find("h2")
	if heading.Length() != 0 {
		return heading.Text()
	}

	return ""
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	main := doc.Find("main")
	if main.Length() != 0 {
		inside := main.Find("p")
		if inside.Length() != 0 {
			return strings.TrimSpace(inside.First().Text())
		}
	}

	para := doc.Find("p")
	if para.Length() == 0 {
		return ""
	}

	return strings.TrimSpace(para.First().Text())
}
