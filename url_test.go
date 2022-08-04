package url_test

import (
	"testing"

	"github.com/hueristiq/url"
)

func TestDefaultScheme(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{input: "localhost", output: "http://localhost"},
		{input: "example.com", output: "http://example.com"},
		{input: "https://example.com", output: "https://example.com"},
		{input: "://example.com", output: "http://example.com"},
		{input: "//example.com", output: "http://example.com"},
	}

	for _, tt := range tests {
		URL := url.DefaultScheme(tt.input, "http")

		if URL != tt.output {
			t.Errorf(`"%s": got "%s", want "%v"`, tt.input, URL, tt.output)
		}
	}
}
