package gocipher

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, strings.ToUpper(test.ciphertext), strings.ToUpper(actual))
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
		assert.Equal(t, strings.ToUpper(text), strings.ToUpper(actual))
	}
}
