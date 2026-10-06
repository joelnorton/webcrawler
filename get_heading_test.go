package main

import "testing"

func TestGetHeadingFromHTMLBasic(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name:      "h1",
			inputBody: "<html><body><h1>Test Title</h1></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "h2",
			inputBody: "<html><body><h2>Test Title2</h2></body></html>",
			expected:  "Test Title2",
		},
		{
			name:      "none",
			inputBody: "<html><body><h3>Test Title2</h3></body></html>",
			expected:  "",
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getHeadingFromHTML(tc.inputBody)

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	inputBody := `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := "Main paragraph."

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}

	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name: "main p",
			inputBody: `<html><body>
				<p>Outside paragraph.</p>
				<main>
					<p>Main paragraph.</p>
					<p>Second paragraph.</p>
				</main>
			</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name: "no main",
			inputBody: `<html><body>
				<p>Outside paragraph.</p>
					<p>Main paragraph.</p>
					<p>Second paragraph.</p>
			</body></html>`,
			expected: "Outside paragraph.",
		},
		{
			name: "no p",
			inputBody: `<html><body>
				<main>
				Main paragraph.
				</main>
			</body></html>`,
			expected: "",
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputBody)

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
