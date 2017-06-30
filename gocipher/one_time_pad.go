package gocipher

import "errors"

type OneTimePad struct {
	key string
}

func NewOneTimePad(key string) *OneTimePad {
	return &OneTimePad{key}
}

// Encrypt encrypts text using a one-time pad
func (otp *OneTimePad) Encrypt(text string) (string, error) {
	if len(otp.key) < len(text) {
		return "", errors.New("key must be at least as long as the plaintext")
	}
	keyChars := []rune(otp.key)
	return mapAlpha(text, func(i, char int) int {
		return char + alphaIndex(keyChars[i])
	}), nil
}

// Decrypt decrypts text using a one-time pad
func (otp *OneTimePad) Decrypt(text string) (string, error) {
	if len(otp.key) < len(text) {
		return "", errors.New("key must be at least as long as the plaintext")
	}
	keyChars := []rune(otp.key)
	return mapAlpha(text, func(i, char int) int {
		return char - alphaIndex(keyChars[i])
	}), nil
}
