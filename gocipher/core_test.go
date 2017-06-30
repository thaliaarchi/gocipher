package gocipher

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomKey(t *testing.T) {
	for length := 10; length < 20; length++ {
		key, err := RandomKey(length)
		if err != nil {
			t.Error("Unexpected error", err)
		}
		if len(key) != length {
			t.Errorf("Expected %q to have length of %d, but got %d", key, len(key), length)
		}
		match, err := regexp.MatchString("[A-Z]", key)
		if !match {
			t.Errorf("Expected alphabetical key with only A-Z, but got %q", key)
		}
		if err != nil {
			t.Error("Unexpected error", err)
		}
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
