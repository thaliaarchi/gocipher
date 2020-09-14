package gocipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOTPEncrypt(t *testing.T) {
	text := "Hello"
	key := "XMCKL"
	expected := "Eqnvz"
	actual, err := NewOTP(key).Encrypt(text)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestOTPDecrypt(t *testing.T) {
	text := "Eqnvz"
	key := "XMCKL"
	expected := "Hello"
	actual, err := NewOTP(key).Decrypt(text)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
