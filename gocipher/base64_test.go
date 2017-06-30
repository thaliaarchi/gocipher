package gocipher

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64Encode(t *testing.T) {
	testBase64Encode(t, "Hello, World! ~~", "SGVsbG8sIFdvcmxkISB+fg$$", func(text string) string {
		return Base64Encode(text, base64.StdEncoding.WithPadding('$'))
	})
}
func TestBase64Decode(t *testing.T) {
	testBase64Decode(t, "SGVsbG8sIFdvcmxkISB+fg$$", "Hello, World! ~~", func(text string) (string, error) {
		return Base64Decode(text, base64.StdEncoding.WithPadding('$'))
	})
}
func TestBase64EncodeStd(t *testing.T) {
	testBase64Encode(t, "Hello, World! ~~", "SGVsbG8sIFdvcmxkISB+fg==", Base64EncodeStd)
}
func TestBase64DecodeStd(t *testing.T) {
	testBase64Decode(t, "SGVsbG8sIFdvcmxkISB+fg==", "Hello, World! ~~", Base64DecodeStd)
}
func TestBase64EncodeURL(t *testing.T) {
	testBase64Encode(t, "Hello, World! ~~", "SGVsbG8sIFdvcmxkISB-fg==", Base64EncodeURL)
}
func TestBase64DecodeURL(t *testing.T) {
	testBase64Decode(t, "SGVsbG8sIFdvcmxkISB-fg==", "Hello, World! ~~", Base64DecodeURL)
}
func TestBase64EncodeRawStd(t *testing.T) {
	testBase64Encode(t, "Hello, World! ~~", "SGVsbG8sIFdvcmxkISB+fg", Base64EncodeRawStd)
}
func TestBase64DecodeRawStd(t *testing.T) {
	testBase64Decode(t, "SGVsbG8sIFdvcmxkISB+fg", "Hello, World! ~~", Base64DecodeRawStd)
}
func TestBase64EncodeRawURL(t *testing.T) {
	testBase64Encode(t, "Hello, World! ~~", "SGVsbG8sIFdvcmxkISB-fg", Base64EncodeRawURL)
}
func TestBase64DecodeRawURL(t *testing.T) {
	testBase64Decode(t, "SGVsbG8sIFdvcmxkISB-fg", "Hello, World! ~~", Base64DecodeRawURL)
}

func testBase64Encode(t *testing.T, text, expected string, f func(string) string) {
	actual := f(text)
	assert.Equal(t, expected, actual)
}

func testBase64Decode(t *testing.T, text, expected string, f func(string) (string, error)) {
	actual, err := f(text)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	assert.Equal(t, expected, actual)
}
