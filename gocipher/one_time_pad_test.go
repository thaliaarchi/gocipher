package gocipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneTimePadEncrypt(t *testing.T) {
	text := "Hello"
	key := "XMCKL"
	expected := "Eqnvz"
	actual, err := OneTimePadEncrypt(text, key)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	assert.Equal(t, expected, actual)
}

func TestOneTimePadDecrypt(t *testing.T) {
	text := "Eqnvz"
	key := "XMCKL"
	expected := "Hello"
	actual, err := OneTimePadDecrypt(text, key)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	assert.Equal(t, expected, actual)
}
