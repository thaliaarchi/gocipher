package gocipher

import "testing"

func TestMorseEncode(t *testing.T) {
	morse := NewMorse(true)
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := ".- -... -.-. -.. . ..-. --. .... .. .--- -.- .-.. -- -. --- .--. --.- .-. ... - ..- ...- .-- -..- -.-- --.."
	expected = expected + " " + expected
	actual, err := morse.Encode(text)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestMorseDecode(t *testing.T) {
	morse := NewMorse(true)
	text := ".- -... -.-. -.. . ..-. --. .... .. .--- -.- .-.. -- -. --- .--. --.- .-. ... - ..- ...- .-- -..- -.-- --.."
	text = text + " " + text
	expected := "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	actual, err := morse.Decode(text)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}
