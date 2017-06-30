package gocipher

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

var base64Tests = []struct {
	plaintext  string
	ciphertext string
	newFunc    func() *Base64
}{
	{"Hello, World! ~~", "SGVsbG8sIFdvcmxkISB+fg$$", func() *Base64 {
		return NewBase64(base64.StdEncoding.WithPadding('$'))
	}},
	{"Hello, World! ~~", "SGVsbG8sIFdvcmxkISB+fg==", NewBase64Std},
	{"Hello, World! ~~", "SGVsbG8sIFdvcmxkISB-fg==", NewBase64URL},
	{"Hello, World! ~~", "SGVsbG8sIFdvcmxkISB+fg", NewBase64RawStd},
	{"Hello, World! ~~", "SGVsbG8sIFdvcmxkISB-fg", NewBase64RawURL},
}

func TestBase64Encode(t *testing.T) {
	for _, test := range base64Tests {
		encoded := test.newFunc().Encode(test.plaintext)
		assert.Equal(t, test.ciphertext, encoded)
	}
}

func TestBase64Decode(t *testing.T) {
	for _, test := range base64Tests {
		decoded, err := test.newFunc().Decode(test.ciphertext)
		if err != nil {
			t.Error("Unexpected error", err)
		}
		assert.Equal(t, test.plaintext, decoded)
	}
}
