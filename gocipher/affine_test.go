package gocipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type affineTest struct {
	expected string
	key      []int
}

func TestAffineEncipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tests := []affineTest{
		{"hijklmnopqrstuvwxyzabcdefgHIJKLMNOPQRSTUVWXYZABCDEFG", []int{1, 7}},
		{"dgjmpsvybehknqtwzcfiloruxaDGJMPSVYBEHKNQTWZCFILORUXA", []int{3, 3}},
		{"afkpuzejotydinsxchmrwbglqvAFKPUZEJOTYDINSXCHMRWBGLQV", []int{5, 0}},
		{"ovcjqxelszgnubipwdkryfmtahOVCJQXELSZGNUBIPWDKRYFMTAH", []int{7, 14}},
		{"sbktcludmvenwfoxgpyhqzirajSBKTCLUDMVENWFOXGPYHQZIRAJ", []int{9, 18}},
		{"pmjgdaxurolifczwtqnkhebyvsPMJGDAXUROLIFCZWTQNKHEBYVS", []int{23, 15}}}
	for _, test := range tests {
		key, err := NewAffine(test.key[0], test.key[1])
		assert.Nil(t, err)
		actual := key.Encipher(text)
		assert.Equal(t, test.expected, actual)
	}
}

func TestAffineDecipher(t *testing.T) {
	text := "pmjgdaxurolifczwtqnkhebyvsPMJGDAXUROLIFCZWTQNKHEBYVS"
	tests := []affineTest{
		{"yfmtahovcjqxelszgnubipwdkrYFMTAHOVCJQXELSZGNUBIPWDKR", []int{7, 3}},
		{"onmlkjihgfedcbazyxwvutsrqpONMLKJIHGFEDCBAZYXWVUTSRQP", []int{3, 25}},
		{"jarizqhypgxofwnevmdulctkbsJARIZQHYPGXOFWNEVMDULCTKBS", []int{9, 12}},
		{"pmjgdaxurolifczwtqnkhebyvsPMJGDAXUROLIFCZWTQNKHEBYVS", []int{1, 0}},
		{"tmfyrkdwpibungzslexqjcvohaTMFYRKDWPIBUNGZSLEXQJCVOHA", []int{19, 18}},
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", []int{23, 15}}}
	for _, test := range tests {
		key, err := NewAffine(test.key[0], test.key[1])
		assert.Nil(t, err)
		actual := key.Decipher(text)
		assert.Equal(t, test.expected, actual)
	}
}

func TestAffinePunctuation(t *testing.T) {
	text := "!@$%%^&*()_-+={}[]|\":;<>,./?"
	key, err := NewAffine(7, 8)
	assert.Nil(t, err)
	actual := key.Encipher(text)
	assert.Equal(t, text, actual) // Punctuation should remain unmodified
}

func TestAtbashEncipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "zyxwvutsrqponmlkjihgfedcbaZYXWVUTSRQPONMLKJIHGFEDCBA"
	actual := Atbash.Encipher(text)
	assert.Equal(t, expected, actual)
}

func TestAtbashDecipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "zyxwvutsrqponmlkjihgfedcbaZYXWVUTSRQPONMLKJIHGFEDCBA"
	actual := Atbash.Decipher(text)
	assert.Equal(t, expected, actual)
}

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
