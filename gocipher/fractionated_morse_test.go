package gocipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFracMorseEncipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "FJHDVGSLMSCCEQFDHQHOEHKTGCPFAPOJQEWCFJHDVGSLMSCCEQFDHQHOEHKTGCPFAPOJQEWC"
	key := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	frac, err := NewFracMorse(key)
	assert.Nil(t, err)
	actual, err := frac.Encipher(text)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestFracMorseDecipher(t *testing.T) {
	text := "FJHDVGSLMSCCEQFDHQHOEHKTGCPFAPOJQEWCFJHDVGSLMSCCEQFDHQHOEHKTGCPFAPOJQEWC"
	expected := "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	key := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	frac, err := NewFracMorse(key)
	assert.Nil(t, err)
	actual, err := frac.Decipher(text)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
