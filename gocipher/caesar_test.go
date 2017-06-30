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
		actual := CaesarEncipher(text, test.key)
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
		actual := CaesarDecipher(text, test.key)
		assert.Equal(t, test.expected, actual)
	}
}

func TestCaesarPunctuation(t *testing.T) {
	text := "!@$%%^&*()_-+={}[]|\":;<>,./?"
	key := 14
	actual := CaesarEncipher(text, key)
	assert.Equal(t, text, actual) // Punctuation should remain unmodified
}

type caesarKeyedTest struct {
	expected string
	shift    int
	key      string
}

func TestCaesarKeyedEncipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tests := []caesarKeyedTest{
		{"helowrdabcfgijkmnpqstuvxyzHELOWRDABCFGIJKMNPQSTUVXYZ", 0, "Hello, World!"},
		{"ywxbcadefghijklmnopqrstuvzYWXBCADEFGHIJKLMNOPQRSTUVZ", 1, "ZYWXZBC"},
		{"cdefghijklmnopqrstuvwxyzabCDEFGHIJKLMNOPQRSTUVWXYZAB", 2, "ABCDEF"},
		{"efghijklmnopqrstuvwxyzabcdEFGHIJKLMNOPQRSTUVWXYZABCD", 4, ""},
		{"hijklmnopqrstuvwxyzabcdefgHIJKLMNOPQRSTUVWXYZABCDEFG", 7, "!@#$%"},
		{"hijklmnopqtuvwxyzcaesrbdfgHIJKLMNOPQTUVWXYZCAESRBDFG", 9, "Caesar"},
		{"zshiftabcdegjklmnopqruvwxyZSHIFTABCDEGJKLMNOPQRUVWXY", -1, "shift"}}
	for _, test := range tests {
		actual := CaesarKeyedEncipher(text, test.shift, test.key)
		assert.Equal(t, test.expected, actual)
	}
}

func TestCaesarKeyedDecipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tests := []caesarKeyedTest{
		{"helowrdabcfgijkmnpqstuvxyzHELOWRDABCFGIJKMNPQSTUVXYZ", 0, "Hello, World!"},
		{"tuvzywxbcadefghijklmnopqrsTUVZYWXBCADEFGHIJKLMNOPQRS", 3, "ZYWXZBC"},
		{"vwxyzabcdefghijklmnopqrstuVWXYZABCDEFGHIJKLMNOPQRSTU", 5, "ABCDEF"},
		{"stuvwxyzabcdefghijklmnopqrSTUVWXYZABCDEFGHIJKLMNOPQR", 8, ""},
		{"pqrstuvwxyzabcdefghijklmnoPQRSTUVWXYZABCDEFGHIJKLMNO", 11, "!@#$%"},
		{"jklmnopqtuvwxyzcaesrbdfghiJKLMNOPQTUVWXYZCAESRBDFGHI", 15, "CAESAR"},
		{"hiftabcdegjklmnopqruvwxyzsHIFTABCDEGJKLMNOPQRUVWXYZS", -1, "SHIFT"}}
	for _, test := range tests {
		actual := CaesarKeyedDecipher(text, test.shift, test.key)
		assert.Equal(t, test.expected, actual)
	}
}
