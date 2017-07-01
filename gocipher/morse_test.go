package gocipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMorseEncode(t *testing.T) {
	morse := NewMorse()
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := ".- -... -.-. -.. . ..-. --. .... .. .--- -.- .-.. -- -. --- .--. --.- .-. ... - ..- ...- .-- -..- -.-- --.."
	expected += " " + expected
	actual, err := morse.Encode(text)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	assert.Equal(t, expected, actual)
}

func TestMorseDecode(t *testing.T) {
	morse := NewMorse()
	text := ".- -... -.-. -.. . ..-. --. .... .. .--- -.- .-.. -- -. --- .--. --.- .-. ... - ..- ...- .-- -..- -.-- --.."
	expected := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	actual, err := morse.Decode(text)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	assert.Equal(t, expected, actual)
}

func TestMorseEncodeProsigns(t *testing.T) {
	morse := NewMorse(MorseInternational, MorseNonEnglish, MorseProsigns)
	text := "AÉCCHA<SOS>A<UndERSTOOD>A"
	expected := ".- ..-.. -.-. ---- .- ...---... .- ...-. .-"
	actual, err := morse.Encode(text)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	assert.Equal(t, expected, actual)
}

func TestMorseToBulletEnDash(t *testing.T) {
	text := "...---... / . .. ... - -- ---"
	expected := "•••–––••• / • •• ••• – –– –––"
	actual := MorseToBulletEnDash(text)
	assert.Equal(t, expected, actual)
}
