package gocipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFracMorseEncipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "FJHDVGSLMSCCEQFDHQHOEHKTGCPFAPOJQEWCFJHDVGSLMSCCEQFDHQHOEHKTGCPFAPOJQEWC"
	key := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	f, err := NewFracMorse(key)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	actual, err := f.Encipher(text)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	assert.Equal(t, expected, actual)
}

func TestFracMorseDecipher(t *testing.T) {
	text := "FJHDVGSLMSCCEQFDHQHOEHKTGCPFAPOJQEWCFJHDVGSLMSCCEQFDHQHOEHKTGCPFAPOJQEWC"
	expected := "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	key := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	f, err := NewFracMorse(key)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	actual, err := f.Decipher(text)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	assert.Equal(t, expected, actual)
}
