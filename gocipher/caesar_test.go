package gocipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type caesarTest struct {
	expected string
	key      int
}

func TestCaesarEncipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tests := []caesarTest{
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 0},
		{"bcdefghijklmnopqrstuvwxyzaBCDEFGHIJKLMNOPQRSTUVWXYZA", 1},
		{"cdefghijklmnopqrstuvwxyzabCDEFGHIJKLMNOPQRSTUVWXYZAB", 2},
		{"efghijklmnopqrstuvwxyzabcdEFGHIJKLMNOPQRSTUVWXYZABCD", 4},
		{"hijklmnopqrstuvwxyzabcdefgHIJKLMNOPQRSTUVWXYZABCDEFG", 7},
		{"jklmnopqrstuvwxyzabcdefghiJKLMNOPQRSTUVWXYZABCDEFGHI", 9},
		{"zabcdefghijklmnopqrstuvwxyZABCDEFGHIJKLMNOPQRSTUVWXY", -1}}
	for _, test := range tests {
		actual := NewCaesar(test.key).Encipher(text)
		assert.Equal(t, test.expected, actual)
	}
}

func TestCaesarDecipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tests := []caesarTest{
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 0},
		{"xyzabcdefghijklmnopqrstuvwXYZABCDEFGHIJKLMNOPQRSTUVW", 3},
		{"vwxyzabcdefghijklmnopqrstuVWXYZABCDEFGHIJKLMNOPQRSTU", 5},
		{"stuvwxyzabcdefghijklmnopqrSTUVWXYZABCDEFGHIJKLMNOPQR", 8},
		{"pqrstuvwxyzabcdefghijklmnoPQRSTUVWXYZABCDEFGHIJKLMNO", 11},
		{"lmnopqrstuvwxyzabcdefghijkLMNOPQRSTUVWXYZABCDEFGHIJK", 15},
		{"bcdefghijklmnopqrstuvwxyzaBCDEFGHIJKLMNOPQRSTUVWXYZA", -1},
	}
	for _, test := range tests {
		actual := NewCaesar(test.key).Decipher(text)
		assert.Equal(t, test.expected, actual)
	}
}

func TestCaesarPunctuation(t *testing.T) {
	text := "!@$%%^&*()_-+={}[]|\":;<>,./?"
	key := 14
	actual := NewCaesar(key).Encipher(text)
	assert.Equal(t, text, actual) // Punctuation should remain unmodified
}
