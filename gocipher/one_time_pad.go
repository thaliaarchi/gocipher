package gocipher

import "errors"

// OneTimePadEncrypt encrypts text using a one-time pad
func OneTimePadEncrypt(text, key string) (string, error) {
	if len(key) < len(text) {
		return "", errors.New("key must be at least as long as the plaintext")
	}
	keyChars := []rune(key)
	return mapAlpha(text, func(i int, char, a, z rune) rune {
		return modRune(char+alphaIndex(keyChars[i])-a, 26) + a
	}), nil
}

// OneTimePadDecrypt decrypts text using a one-time pad
func OneTimePadDecrypt(text, key string) (string, error) {
	if len(key) < len(text) {
		return "", errors.New("key must be at least as long as the plaintext")
	}
	keyChars := []rune(key)
	return mapAlpha(text, func(i int, char, a, z rune) rune {
		return modRune(char-alphaIndex(keyChars[i])-a, 26) + a
	}), nil
}
