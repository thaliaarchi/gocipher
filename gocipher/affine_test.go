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
