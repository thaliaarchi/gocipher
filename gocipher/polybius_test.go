package gocipher

import (
	"strings"
	"testing"
)

type polybiusTest struct {
	key        string
	size       int
	chars      string
	ciphertext string
}

func TestPolybiusEncipher(t *testing.T) {
	text := "abcdefghiiklmnopqrstuvwxyzabcdefghiiklmnopqrstuvwxyz"
	tests := []polybiusTest{
		{"phqgiumeaylnofdxkrcvstzwb", 5, "ABCDE",
			"BDEEDDCEBCCDADABAEAEDBCABBCBCCAAACDCEAEBBADEEDDABEECBDEEDDCEBCCDADABAEAEDBCABBCBCCAAACDCEAEBBADEEDDABEEC"},
		{"uqfigkydlvmznxephrswaotcb", 5, "BCDEF",
			"FBFFFECDDFBDBFECBEBECBCEDBDDFCEBBCEDEEFDBBCFEFDECCDCFBFFFECDDFBDBFECBEBECBCEDBDDFCEBBCEDEEFDBBCFEFDECCDC"}}
	for _, test := range tests {
		p, err := NewPolybius(test.key, test.size, test.chars)
		if err != nil {
			t.Error("Unexpected error", err)
		}
		actual := p.Encipher(text)
		if strings.ToUpper(actual) != strings.ToUpper(test.ciphertext) {
			t.Errorf("Expected %q, but got %q", test.ciphertext, actual)
		}
	}
}

func TestPolybiusDecipher(t *testing.T) {
	text := "abcdefghiiklmnopqrstuvwxyzabcdefghiiklmnopqrstuvwxyz"
	tests := []polybiusTest{
		{"phqgiumeaylnofdxkrcvstzwb", 5, "ABCDE",
			"BDEEDDCEBCCDADABAEAEDBCABBCBCCAAACDCEAEBBADEEDDABEECBDEEDDCEBCCDADABAEAEDBCABBCBCCAAACDCEAEBBADEEDDABEEC"},
		{"uqfigkydlvmznxephrswaotcb", 5, "BCDEF",
			"FBFFFECDDFBDBFECBEBECBCEDBDDFCEBBCEDEEFDBBCFEFDECCDCFBFFFECDDFBDBFECBEBECBCEDBDDFCEBBCEDEEFDBBCFEFDECCDC"}}
	for _, test := range tests {
		p, err := NewPolybius(test.key, test.size, test.chars)
		if err != nil {
			t.Error("Unexpected error", err)
		}
		actual := p.Decipher(test.ciphertext)
		if strings.ToUpper(actual) != strings.ToUpper(text) {
			t.Errorf("Expected %q, but got %q", text, actual)
		}
	}
}
