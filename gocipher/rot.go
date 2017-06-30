package gocipher

import (
	"strings"
	"unicode"
)

/*
 * ROT
 * ROT-5 cipher
 * ROT-13 cipher
 * ROT-18 cipher
 * ROT-47 cipher
 */

type ROT struct {
	key      int
	alphabet string
}

func NewROT(key int, alphabet string) *ROT {
	return &ROT{key, alphabet}
}

// Encipher enciphers string using ROT cipher with alphabet according to key.
func (r *ROT) Encipher(text string) string {
	return rotEncipher(text, r.key, r.alphabet)
}

// Decipher deciphers string using ROT cipher with alphabet according to key.
func (r *ROT) Decipher(text string) string {
	return rotEncipher(text, -r.key, r.alphabet)
}

func rotEncipher(text string, key int, alphabet string) string {
	size := len(alphabet)
	alphaRunes := []rune(alphabet)
	runes := []rune(text)
	for i, char := range runes {
		if pos := strings.IndexRune(alphabet, char); pos != -1 {
			runes[i] = alphaRunes[mod(pos+key, size)]
		}
	}
	return string(runes)
}

// ROTEncipherCaps enciphers string using ROT cipher with alphabet according to key.
// Preserves capitalization.
func ROTEncipherCaps(text string, key int, alphabet string) string {
	size := len(alphabet)
	alphabet = strings.ToLower(alphabet)
	alphaRunes := []rune(alphabet)
	runes := []rune(text)
	for i, char := range runes {
		charLower := unicode.ToLower(char)
		if pos := strings.IndexRune(alphabet, charLower); pos != -1 {
			shifted := alphaRunes[mod(pos+key, size)]
			if unicode.IsUpper(char) {
				shifted = unicode.ToUpper(shifted)
			}
			runes[i] = shifted
		}
	}
	return string(runes)
}

// ROTDecipherCaps deciphers string using ROT cipher with alphabet according to key.
// Preserves capitalization.
func ROTDecipherCaps(text string, key int, alphabet string) string {
	return ROTEncipherCaps(text, -key, alphabet)
}

// ROTEncipherRange enciphers string using ROT cipher with ranged alphabet according to key.
// Uppercase and lowercase are considered different characters.
func ROTEncipherRange(text string, key int, min, max rune) string {
	size := max - min + 1
	shift := rune(key)
	runes := []rune(text)
	for i, char := range runes {
		if char >= min && char <= max {
			runes[i] = modRune(char+shift-min, size) + min
		}
	}
	return string(runes)
}

// ROTDecipherRange deciphers string using ROT cipher with ranged alphabet according to key.
// Uppercase and lowercase are considered different characters.
func ROTDecipherRange(text string, key int, min, max rune) string {
	return ROTEncipherRange(text, -key, min, max)
}

// ROT5Encipher enciphers string using ROT-5 cipher. Identical to ROT5Decipher.
// e.g. "1234567890" becomes "5678901234".
func ROT5Encipher(text string) string {
	return ROTEncipherRange(text, 5, '0', '9')
}

// ROT5Decipher deciphers string using ROT-5 cipher. Identical to ROT5Encipher.
// e.g. "5678901234" becomes "1234567890".
func ROT5Decipher(text string) string {
	return ROT5Encipher(text)
}

// ROT13Encipher enciphers string using ROT-13 cipher. Identical to ROT13Decipher.
// e.g. "ABCDEFGHIJKLM" becomes "NOPQRSTUVWXYZ".
func ROT13Encipher(text string) string {
	return caesarEncipher(text, 13)
}

// ROT13Decipher deciphers string using ROT-13 cipher. Identical to ROT13Encipher.
// e.g. "NOPQRSTUVWXYZ" becomes "ABCDEFGHIJKLM".
func ROT13Decipher(text string) string {
	return ROT13Encipher(text)
}

// ROT18Encipher enciphers string using ROT-18 cipher. Identical to ROT18Decipher.
// e.g. "ABCXYZ012" becomes "STUFGHijk".
func ROT18Encipher(text string) string {
	return ROT13Encipher(ROT5Encipher(text))
}

// ROT18Decipher deciphers string using ROT-18 cipher. Identical to ROT18Encipher.
// e.g. "STUFGHIJK" becomes "ABCXYZ012".
func ROT18Decipher(text string) string {
	return ROT18Encipher(text)
}

// ROT47Encipher enciphers string using ROT-47 cipher. Identical to ROT47Decipher.
// e.g. "ABCabc" becomes "pqr234".
func ROT47Encipher(text string) string {
	return ROTEncipherRange(text, 47, '!', '~')
}

// ROT47Decipher deciphers string using ROT-47 cipher. Identical to ROT47Encipher.
// e.g. "pqr234" becomes "ABCabc".
func ROT47Decipher(text string) string {
	return ROT47Encipher(text)
}
