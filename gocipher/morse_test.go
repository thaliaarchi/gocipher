package gocipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMorseEncode(t *testing.T) {
	morse := NewMorse(true, true)
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := ".- -... -.-. -.. . ..-. --. .... .. .--- -.- .-.. -- -. --- .--. --.- .-. ... - ..- ...- .-- -..- -.-- --.."
	expected = expected + " " + expected
	actual, err := morse.Encode(text)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	assert.Equal(t, expected, actual)
}

func TestMorseDecode(t *testing.T) {
	morse := NewMorse(true, true)
	text := ".- -... -.-. -.. . ..-. --. .... .. .--- -.- .-.. -- -. --- .--. --.- .-. ... - ..- ...- .-- -..- -.-- --.."
	text = text + " " + text
	expected := "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	actual, err := morse.Decode(text)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	assert.Equal(t, expected, actual)
}
