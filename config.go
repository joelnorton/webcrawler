package main

import (
	"net/url"
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

func createConfig(baseURL url.URL, concurrency int, pages int) *config {
	return &config{
		pages:              make(map[string]PageData),
		baseURL:            &baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, concurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           pages,
	}
}
