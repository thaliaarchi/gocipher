package gocipher

import "errors"

// OneTimePadEncrypt encrypts text using a one-time pad
func OneTimePadEncrypt(text, key string) (string, error) {
	if len(key) < len(text) {
		return "", errors.New("key must be at least as long as the plaintext")
	}
	keyChars := []rune(key)
	return monoalphabetic(text, func(i, char int) int {
		return char + alphaIndex(keyChars[i])
	}), nil
}

// OneTimePadDecrypt decrypts text using a one-time pad
func OneTimePadDecrypt(text, key string) (string, error) {
	if len(key) < len(text) {
		return "", errors.New("key must be at least as long as the plaintext")
	}
	keyChars := []rune(key)
	return monoalphabetic(text, func(i, char int) int {
		return char - alphaIndex(keyChars[i])
	}), nil
}
