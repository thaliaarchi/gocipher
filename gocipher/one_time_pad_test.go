package gocipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneTimePadEncrypt(t *testing.T) {
	text := "Hello"
	key := "XMCKL"
	expected := "Eqnvz"
	actual, err := NewOneTimePad(key).Encrypt(text)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	assert.Equal(t, expected, actual)
}

func TestOneTimePadDecrypt(t *testing.T) {
	text := "Eqnvz"
	key := "XMCKL"
	expected := "Hello"
	actual, err := NewOneTimePad(key).Decrypt(text)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	assert.Equal(t, expected, actual)
}
