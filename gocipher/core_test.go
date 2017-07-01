package gocipher

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomKey(t *testing.T) {
	for length := 10; length < 20; length++ {
		key, err := RandomKey(length)
		assert.Nil(t, err)
		assert.Equal(t, len(key), length)
		match, err := regexp.MatchString("[A-Z]", key)
		assert.Nil(t, err)
		assert.True(t, match, "Key must be alphabetical with only A-Z")
	}
}

func TestKeyedAlphabet(t *testing.T) {
	key := "Hello, World!"
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "HELOWRDABCFGIJKMNPQSTUVXYZ"
	actual := KeyedAlphabet(key, alphabet)
	assert.Equal(t, expected, actual)
}

func TestKeyedAlphabetRange(t *testing.T) {
	key := "HELLO, WORLD!"
	min, max := 'A', 'Z'
	expected := "HELOWRDABCFGIJKMNPQSTUVXYZ"
	actual := KeyedAlphabetRange(key, min, max)
	assert.Equal(t, expected, actual)
}
