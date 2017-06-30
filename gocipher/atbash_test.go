package gocipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtbashEncipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "zyxwvutsrqponmlkjihgfedcbaZYXWVUTSRQPONMLKJIHGFEDCBA"
	actual := AtbashEncipher(text)
	assert.Equal(t, expected, actual)
}

func TestAtbashDecipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "zyxwvutsrqponmlkjihgfedcbaZYXWVUTSRQPONMLKJIHGFEDCBA"
	actual := AtbashDecipher(text)
	assert.Equal(t, expected, actual)
}
