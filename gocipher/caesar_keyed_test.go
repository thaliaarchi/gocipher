package gocipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		actual := NewCaesarKeyed(test.key, test.shift).Encipher(text)
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
		actual := NewCaesarKeyed(test.key, test.shift).Decipher(text)
		assert.Equal(t, test.expected, actual)
	}
}
