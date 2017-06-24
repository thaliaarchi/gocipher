package gocipher

import "errors"

func OneTimePadEncrypt(text, key string) (string, error) {
	if len(key) < len(text) {
		return "", errors.New("key must be at least as long as the plaintext")
	}
	keys := []rune(key)
	return mapAlpha(text, func(i int, char, a, z rune) rune {
		return char + (keys[i] - a)
	}), nil
}

func OneTimePadDecrypt(text, key string) (string, error) {
	if len(key) < len(text) {
		return "", errors.New("key must be at least as long as the plaintext")
	}
	keys := []rune(key)
	return mapAlpha(text, func(i int, char, a, z rune) rune {
		return char - (keys[i] - a)
	}), nil
}
