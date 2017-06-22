package gocipher

import "testing"

type caesarTest struct {
	expected string
	key      int
}

// Test known plaintext->ciphertext pairs
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
		output := CaesarEncipher(text, test.key)
		if output != test.expected {
			t.Errorf("Expected %q, but got %q (key: %d)", test.expected, output, test.key)
		}
	}
}

// Test known ciphertext->plaintext pairs
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
		output := CaesarDecipher(text, test.key)
		if output != test.expected {
			t.Errorf("Expected %q, but got %q (key: %d)", test.expected, output, test.key)
		}
	}
}

// Punctuation should remain unmodified
func TestCaesarPunctuation(t *testing.T) {
	text := "!@$%%^&*()_-+={}[]|\":;<>,./?"
	key := 14
	actual := CaesarEncipher(text, key)
	if text != actual {
		t.Errorf("Expected %q, but got %q (key %d)", text, actual, key)
	}
}
