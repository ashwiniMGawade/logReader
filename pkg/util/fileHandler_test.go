package util

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected string
		err      error
	}{
		{
			name:     "Test reading a file with some content",
			filePath: "fixtures/testfile.log",
			expected: `177.71.128.21 - - [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7"
168.41.191.40 - - [09/Jul/2018:10:11:30 +0200] "GET http://example.net/faq/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (Linux; U; Android 2.3.5; en-us; HTC Vision Build/GRI40) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"`,
			err: nil,
		},
		{
			name:     "Test reading an empty file",
			filePath: "fixtures/emptyfile.log",
			expected: "",
			err:      nil,
		},
		{
			name:     "Test reading whe file does not exist",
			filePath: "nonexistent.log",
			expected: "",
			err:      fmt.Errorf("Error reading file: open nonexistent.log: no such file or directory"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			content, err := ReadFile(test.filePath)

			if test.err != nil {
				if err == nil {
					t.Errorf("Expected error to be present, but got: %v", err)
				}
			} else if err != nil {
				t.Errorf("Expected error to be nil, but got: %v", err)
			}

			if content == test.expected {
				t.Logf("Expected content %q, got %q", test.expected, content)
			}
		})
	}
}
