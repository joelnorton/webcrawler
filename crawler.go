package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()
	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()
	fmt.Printf("Checking %s\n", rawCurrentURL)
	current, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}
	if cfg.baseURL.Host != current.Host {
		return
	}
	normal, _ := normalizeURL(rawCurrentURL)
	if !cfg.addPageVisit(normal) {
		return
	}
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	data := extractPageData(html, rawCurrentURL)
	cfg.setPageData(normal, data)
	for _, link := range cfg.pages[normal].OutgoingLinks {
		cfg.wg.Add(1)
		go cfg.crawlPage(link)
	}

}
