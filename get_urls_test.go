package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "3 links",
			inputURL: "https://crawler-test.com",
			inputBody: `<html><body>
        <h1>Test Title</h1>
        <p>This is the first paragraph.</p>
        <a href="/link1">Link 1</a>
        <img src="/image1.jpg" alt="Image 1">
    </body></html>`,
			expected: []string{
				"https://crawler-test.com/link1",
			},
		},
		{
			name:     "2 links",
			inputURL: "https://crawler-test.com",
			inputBody: `
				<!DOCTYPE html>
				<html>
				<head><title>Link Test</title></head>
				<body>
					<main>
						<!-- 1. Absolute Link (Contains the full protocol and domain) -->
						<a href="https://crawler-test.com">Absolute Link</a>

						<!-- 2. Relative Link (Relative to the current directory/path) -->
						<a href="about-us.html">Relative Link</a>
								</main>
				</body>
				</html>`,
			expected: []string{
				"https://crawler-test.com",
				"https://crawler-test.com/about-us.html",
			},
		}, {
			name:     "1 links",
			inputURL: "https://crawler-test.com",
			inputBody: `
				<!DOCTYPE html>
				<html>
				<head><title>Link Test</title></head>
				<body>
					<main>

						<!-- 3. Root-Relative Link (Relative to the domain's root folder) -->
						<a href="/contact">Root-Relative Link</a>
					</main>
				</body>
				</html>`,
			expected: []string{
				"https://crawler-test.com/contact",
			},
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			url, _ := url.Parse(tc.inputURL)
			actual, err := getURLsFromHTML(tc.inputBody, url)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}

		})
	}
}

func TestGetImagesFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "3 links",
			inputURL: "https://crawler-test.com",
			inputBody: `
				<!DOCTYPE html>
				<html>
				<head><title>Image Test</title></head>
				<body>
					<main>
						<!-- 1. Absolute Image Path (Full protocol and domain) -->
						<img src="https://crawler-test.com" alt="Absolute Logo">

						<!-- 2. Relative Image Path (Relative to the current directory) -->
						<img src="images/banner.jpg" alt="Relative Banner">

						<!-- 3. Root-Relative Image Path (Relative to the domain's root) -->
						<img src="/icons/avatar.svg" alt="Root-Relative Avatar">
					</main>
				</body>
				</html>`,
			expected: []string{
				"https://crawler-test.com",
				"https://crawler-test.com/images/banner.jpg",
				"https://crawler-test.com/icons/avatar.svg",
			},
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			url, _ := url.Parse(tc.inputURL)
			actual, err := getImagesFromHTML(tc.inputBody, url)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
