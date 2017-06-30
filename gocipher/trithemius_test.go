package gocipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrithemiusEncipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "acegikmoqsuwyacegikmoqsuwyACEGIKMOQSUWYACEGIKMOQSUWY"
	actual := TrithemiusEncipher(text)
	assert.Equal(t, expected, actual)
}

func TestTrithemiusDecipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "aaaaaaaaaaaaaaaaaaaaaaaaaaAAAAAAAAAAAAAAAAAAAAAAAAAA"
	actual := TrithemiusDecipher(text)
	assert.Equal(t, expected, actual)
}
