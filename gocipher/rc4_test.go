package gocipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var rc4Tests = []struct {
	key        string
	plaintext  string
	ciphertext string
}{
	// Tests samples from: http://en.wikipedia.org/wiki/RC4#Test_vectors
	{"Key", "Plaintext", "\xBB\xF3\x16\xE8\xD9\x40\xAF\x0A\xD3"},
	{"Wiki", "pedia", "\x10\x21\xBF\x04\x20"},
	{"Secret", "Attack at dawn", "\x45\xA0\x1F\x64\x5F\xC3\x5B\x38\x35\x52\x54\x4B\x9B\xF5"}}

func TestRC4(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	key := "Hello, World!"
	rc4 := NewRC4(key)
	actual := rc4.Decipher(rc4.Encipher(text))
	assert.Equal(t, text, actual)
}

func TestRC4Encipher(t *testing.T) {
	for _, test := range rc4Tests {
		actual := NewRC4(test.key).Encipher(test.plaintext)
		assert.Equal(t, test.ciphertext, actual)
	}
}

func TestRC4Decipher(t *testing.T) {
	for _, test := range rc4Tests {
		actual := NewRC4(test.key).Decipher(test.ciphertext)
		assert.Equal(t, test.plaintext, actual)
	}
}
