package gocipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtbashEncipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "zyxwvutsrqponmlkjihgfedcbaZYXWVUTSRQPONMLKJIHGFEDCBA"
	actual := NewAtbash().Encipher(text)
	assert.Equal(t, expected, actual)
}

func TestAtbashDecipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "zyxwvutsrqponmlkjihgfedcbaZYXWVUTSRQPONMLKJIHGFEDCBA"
	actual := NewAtbash().Decipher(text)
	assert.Equal(t, expected, actual)
}
