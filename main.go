package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
)

type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func main() {

	if len(os.Args) < 4 {
		fmt.Println("too few arguments provided")
		os.Exit(1)
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL := os.Args[1]
	maxConcurrency, _ := strconv.Atoi(os.Args[2])
	maxPages, _ := strconv.Atoi(os.Args[3])
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("Cannot parse URL")
		os.Exit(1)
	}
	cfg := &config{
		pages:              make(map[string]PageData),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}
	fmt.Printf("starting crawl of: %s\n", rawBaseURL)
	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()
	for page, _ := range cfg.pages {
		fmt.Printf("Page %s\n", page)
	}
	writeJSONReport(cfg.pages, "report.json")
}

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

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, ok := cfg.pages[normalizedURL]; ok {
		return false
	}
	cfg.pages[normalizedURL] = PageData{URL: normalizedURL}

	return true

}

func (cfg *config) setPageData(normalizedURL string, data PageData) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	cfg.pages[normalizedURL] = data
}
