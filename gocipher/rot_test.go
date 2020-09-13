package gocipher

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestROTEncipher(t *testing.T) {
	var text = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var expected = "STUVWXYZ0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789ABCDEFGHIJKLMNOPQR"
	actual := NewROT(18, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789").Encipher(strings.ToUpper(text))
	assert.Equal(t, expected, actual)
}

func TestROTDecipher(t *testing.T) {
	var text = "stuvwxyz0123456789abcdefghijklmnopqrSTUVWXYZ0123456789ABCDEFGHijklmnopqr"
	var expected = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // Loses capitalization on numbers
	actual := NewROT(18, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789").Decipher(strings.ToUpper(text))
	assert.Equal(t, expected, actual)
}

func TestROTEncipherCaps(t *testing.T) {
	var text = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var expected = "stuvwxyz0123456789abcdefghijklmnopqrSTUVWXYZ0123456789ABCDEFGHijklmnopqr"
	actual := rotEncipherCaps(text, 18, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	assert.Equal(t, expected, actual)
}

func TestROTDecipherCaps(t *testing.T) {
	var text = "stuvwxyz0123456789abcdefghijklmnopqrSTUVWXYZ0123456789ABCDEFGHijklmnopqr"
	var expected = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHijklmnopqrSTUVWXYZ0123456789" // Loses capitalization on numbers
	actual := rotEncipherCaps(text, -18, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	assert.Equal(t, expected, actual)
}

func TestROTRangeEncipher(t *testing.T) {
	var text = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var expected = "abcdefghijklmnopqrstuvwxyz0123456789NOPQRSTUVWXYZABCDEFGHIJKLM0123456789"
	actual := NewROTRange(13, 'A', 'Z').Encipher(text) // Only changes A-Z, not lowercase
	assert.Equal(t, expected, actual)
}

func TestROTRangeDecipher(t *testing.T) {
	var text = "abcdefghijklmnopqrstuvwxyz0123456789NOPQRSTUVWXYZABCDEFGHIJKLM0123456789"
	var expected = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	actual := NewROTRange(13, 'A', 'Z').Decipher(text) // Only changes A-Z, not lowercase
	assert.Equal(t, expected, actual)
}

func TestROT5Encipher(t *testing.T) {
	text := "0123456789"
	expected := "5678901234"
	actual := ROT5.Encipher(text)
	assert.Equal(t, expected, actual)
}

func TestROT5Decipher(t *testing.T) {
	text := "5678901234"
	expected := "0123456789"
	actual := ROT5.Decipher(text)
	assert.Equal(t, expected, actual)
}

func TestROT13Encipher(t *testing.T) {
	var text = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var expected = "nopqrstuvwxyzabcdefghijklmNOPQRSTUVWXYZABCDEFGHIJKLM"
	actual := ROT13.Encipher(text)
	assert.Equal(t, expected, actual)
}

func TestROT13Decipher(t *testing.T) {
	var text = "nopqrstuvwxyzabcdefghijklmNOPQRSTUVWXYZABCDEFGHIJKLM"
	var expected = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	actual := ROT13.Decipher(text)
	assert.Equal(t, expected, actual)
}

func TestROT18Encipher(t *testing.T) {
	var text = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var expected = "nopqrstuvwxyzabcdefghijklm5678901234NOPQRSTUVWXYZABCDEFGHIJKLM5678901234"
	actual := ROT18.Encipher(text)
	assert.Equal(t, expected, actual)
}

func TestROT18Decipher(t *testing.T) {
	var text = "nopqrstuvwxyzabcdefghijklm5678901234NOPQRSTUVWXYZABCDEFGHIJKLM5678901234"
	var expected = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	actual := ROT18.Decipher(text)
	assert.Equal(t, expected, actual)
}
func TestROT47Encipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "23456789:;<=>?@ABCDEFGHIJKpqrstuvwxyz{|}~!\"#$%&'()*+"
	actual := ROT47.Encipher(text)
	assert.Equal(t, expected, actual)
}

func TestROT47Decipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "23456789:;<=>?@ABCDEFGHIJKpqrstuvwxyz{|}~!\"#$%&'()*+"
	actual := ROT47.Decipher(text)
	assert.Equal(t, expected, actual)
}
