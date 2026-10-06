package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(u string) (string, error) {

	url, err := url.Parse(u)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	out := url.Host + url.Path

	out = strings.ToLower(out)

	out = strings.TrimSuffix(out, "/")

	return out, nil
}
