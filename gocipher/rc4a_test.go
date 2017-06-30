package gocipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRC4ASymmetric(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	key := "Hello, World!"
	rc4a := NewRC4A(key)
	ciphertext := rc4a.Encipher(text)
	actual := rc4a.Decipher(ciphertext)
	assert.Equal(t, text, actual)
}
