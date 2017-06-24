package gocipher

import (
	"strings"
)

/*
 * Caesar cipher
 * Keyed Caesar cipher
 */

// CaesarEncipher enciphers string using Caesar cipher according to key.
func CaesarEncipher(text string, key int) string {
	shift := rune(key)
	runes := []rune(text)
	for i, char := range runes {
		if char >= 'A' && char <= 'Z' {
			runes[i] = modRune(char+shift-'A', 26) + 'A'
		} else if char >= 'a' && char <= 'z' {
			runes[i] = modRune(char+shift-'a', 26) + 'a'
		}
	}
	return string(runes)
}

// CaesarDecipher deciphers string using Caesar cipher according to key.
func CaesarDecipher(text string, key int) string {
	return CaesarEncipher(text, -key)
}

// CaesarKeyedEncipher enciphers string using keyed Caesar cipher according to key.
func CaesarKeyedEncipher(text string, shift int, key string) string {
	alphabet := makeKeyedAlphabet(key)
	alpha := []rune(alphabet)
	s := rune(shift)
	runes := []rune(text)
	for i, char := range runes {
		if char >= 'A' && char <= 'Z' {
			runes[i] = alpha[modRune(char+s-'A', 26)]
		} else if char >= 'a' && char <= 'z' {
			runes[i] = alpha[modRune(char+s-'a', 26)] - 'A' + 'a'
		}
	}
	return string(runes)
}

// CaesarKeyedDecipher deciphers string using keyed Caesar cipher according to key.
func CaesarKeyedDecipher(text string, shift int, key string) string {
	return CaesarKeyedEncipher(text, -shift, key)
}

func makeKeyedAlphabet(key string) string {
	chars := []rune(strings.ToUpper(key))
	alphabet := make([]rune, 26)
	pos := 0
	for _, char := range chars {
		if char >= 'A' && char <= 'Z' && !strings.ContainsRune(string(alphabet), char) {
			alphabet[pos] = char
			pos++
		}
	}
	for i := 'A'; i <= 'Z'; i++ {
		if !strings.ContainsRune(string(alphabet), i) {
			alphabet[pos] = i
			pos++
		}
	}
	return string(alphabet)
}
