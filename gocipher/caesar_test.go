package gocipher

import "testing"

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
		output := CaesarEncipher(text, test.key)
		if output != test.expected {
			t.Errorf("Expected %q, but got %q (key: %d)", test.expected, output, test.key)
		}
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
		output := CaesarDecipher(text, test.key)
		if output != test.expected {
			t.Errorf("Expected %q, but got %q (key: %d)", test.expected, output, test.key)
		}
	}
}

func TestCaesarPunctuation(t *testing.T) {
	text := "!@$%%^&*()_-+={}[]|\":;<>,./?"
	key := 14
	actual := CaesarEncipher(text, key)
	if text != actual { // Punctuation should remain unmodified
		t.Errorf("Expected %q, but got %q (key %d)", text, actual, key)
	}
}

type caesarKeyedTest struct {
	expected string
	shift    int
	key      string
}

func TestCaesarKeyedEncipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tests := []caesarKeyedTest{
		{"heloabcdfgijkmnpqrstuvwxyzHELOABCDFGIJKMNPQRSTUVWXYZ", 0, "Hello"},
		{"orldabcefghijkmnpqstuvxyzwORLDABCEFGHIJKMNPQSTUVXYZW", 1, "World"},
		{"wxbcadefghijklmnopqrstuvzyWXBCADEFGHIJKLMNOPQRSTUVZY", 2, "ZYWXZBC"},
		{"efghijklmnopqrstuvwxyzabcdEFGHIJKLMNOPQRSTUVWXYZABCD", 4, ""},
		{"hijklmnopqrstuvwxyzabcdefgHIJKLMNOPQRSTUVWXYZABCDEFG", 7, "!@#$%"},
		{"hijklmnopqtuvwxyzcaesrbdfgHIJKLMNOPQTUVWXYZCAESRBDFG", 9, "CAESAR"},
		{"zshiftabcdegjklmnopqruvwxyZSHIFTABCDEGJKLMNOPQRUVWXY", -1, "SHIFT"}}
	for _, test := range tests {
		output := CaesarKeyedEncipher(text, test.shift, test.key)
		if output != test.expected {
			t.Errorf("Expected %q, but got %q (key: %q, shift: %d)", test.expected, output, test.key, test.shift)
		}
	}
}

func TestCaesarKeyedDecipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tests := []caesarKeyedTest{
		{"heloabcdfgijkmnpqrstuvwxyzHELOABCDFGIJKMNPQRSTUVWXYZ", 0, "Hello"},
		{"xyzworldabcefghijkmnpqstuvXYZWORLDABCEFGHIJKMNPQSTUV", 3, "World"},
		{"rstuvzywxbcadefghijklmnopqRSTUVZYWXBCADEFGHIJKLMNOPQ", 5, "ZYWXZBC"},
		{"stuvwxyzabcdefghijklmnopqrSTUVWXYZABCDEFGHIJKLMNOPQR", 8, ""},
		{"pqrstuvwxyzabcdefghijklmnoPQRSTUVWXYZABCDEFGHIJKLMNO", 11, "!@#$%"},
		{"jklmnopqtuvwxyzcaesrbdfghiJKLMNOPQTUVWXYZCAESRBDFGHI", 15, "CAESAR"},
		{"hiftabcdegjklmnopqruvwxyzsHIFTABCDEGJKLMNOPQRUVWXYZS", -1, "SHIFT"}}
	for _, test := range tests {
		output := CaesarKeyedDecipher(text, test.shift, test.key)
		if output != test.expected {
			t.Errorf("Expected %q, but got %q (key: %q, shift: %d)", test.expected, output, test.key, test.shift)
		}
	}
}
