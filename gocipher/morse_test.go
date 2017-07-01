package gocipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMorseEncode(t *testing.T) {
	morse := NewMorse()
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := ".- -... -.-. -.. . ..-. --. .... .. .--- -.- .-.. -- -. --- .--. --.- .-. ... - ..- ...- .-- -..- -.-- --.."
	actual, err := morse.Encode(text)
	assert.Nil(t, err)
	assert.Equal(t, expected+" "+expected, actual)
}

func TestMorseDecode(t *testing.T) {
	morse := NewMorse()
	text := ".- -... -.-. -.. . ..-. --. .... .. .--- -.- .-.. -- -. --- .--. --.- .-. ... - ..- ...- .-- -..- -.-- --.."
	expected := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	actual, err := morse.Decode(text)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestMorseEncodeProsigns(t *testing.T) {
	morse := NewMorse(MorseInternational, MorseNonEnglish, MorseProsigns)
	text := "AÉCChA<sos>A<SN>A"
	expected := ".- ..-.. -.-. ---- .- ...---... .- ...-. .-"
	actual, err := morse.Encode(text)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestMorseFormatBullets(t *testing.T) {
	text := "...---... / . .. ... - -- ---"
	expected := "•••–––••• / • •• ••• – –– –––"
	actual := MorseFormatBullets(text)
	assert.Equal(t, expected, actual)
}

// Test from: https://en.wikipedia.org/wiki/Morse_code#Spoken_representation
func TestMorseFormatSpoken(t *testing.T) {
	text := "-- --- .-. ... . / -.-. --- -.. ." // text for: Morse Code
	expected := "Dah-dah dah-dah-dah di-dah-dit di-di-dit dit, Dah-di-dah-dit dah-dah-dah dah-di-dit dit."
	actual := MorseFormatSpoken(text)
	assert.Equal(t, expected, actual)
}
