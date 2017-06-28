package gocipher

import "testing"

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
		key, err := NewAffineKey(test.key[0], test.key[1])
		if err != nil {
			t.Error("Key creation error")
		}
		actual := AffineEncipher(text, key)
		if test.expected != actual {
			t.Errorf("Expected %q, but got %q (text: %q, key %v)", test.expected, actual, text, key)
		}
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
		key, err := NewAffineKey(test.key[0], test.key[1])
		if err != nil {
			t.Error("Key creation error")
		}
		actual := AffineDecipher(text, key)
		if test.expected != actual {
			t.Errorf("Expected %q, but got %q (text: %q, key %v)", test.expected, actual, text, test.key)
		}
	}
}

func TestAffinePunctuation(t *testing.T) {
	text := "!@$%%^&*()_-+={}[]|\":;<>,./?"
	key, err := NewAffineKey(7, 8)
	if err != nil {
		t.Error("Key creation error")
	}
	actual := AffineEncipher(text, key)
	if text != actual { // Punctuation should remain unmodified
		t.Errorf("Expected %q, but got %q (text: %q, key %v)", text, actual, text, key)
	}
}
