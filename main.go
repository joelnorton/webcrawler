package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 4 {
		fmt.Println("too few arguments provided")
		os.Exit(1)
	}
	if len(os.Args) > 5 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL := os.Args[1]
	maxConcurrency, _ := strconv.Atoi(os.Args[2])
	maxPages, _ := strconv.Atoi(os.Args[3])
	var filename string
	if len(os.Args) == 5 {
		filename = os.Args[4]
	} else {
		filename = "report.json"
	}

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("Cannot parse URL")
		os.Exit(1)
	}
	cfg := createConfig(*baseURL, maxConcurrency, maxPages)
	fmt.Printf("starting crawl of: %s\n", rawBaseURL)
	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()
	err = writeJSONReport(cfg.pages, filename)
	if err != nil {
		fmt.Println("Error writing file")
	} else {
		fmt.Printf("Results written to %s\n", filename)
	}
}
