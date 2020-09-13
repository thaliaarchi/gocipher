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

// Encipher enciphers string using ROT cipher with alphabet according to key.
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
func rotEncipherCaps(text string, key int, alphabet string) string {
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

type ROTRange struct {
	key int
	min rune
	max rune
}

func NewROTRange(key int, min, max rune) *ROTRange {
	return &ROTRange{key, min, max}
}

// Encipher enciphers string using ROT cipher with ranged alphabet according to key.
// Uppercase and lowercase are considered different characters.
func (r *ROTRange) Encipher(text string) string {
	return rotEncipherRange(text, r.key, r.min, r.max)
}

// Decipher deciphers string using ROT cipher with ranged alphabet according to key.
// Uppercase and lowercase are considered different characters.
func (r *ROTRange) Decipher(text string) string {
	return rotEncipherRange(text, -r.key, r.min, r.max)
}

// rotEncipherRange enciphers string using ROT cipher with ranged alphabet according to key.
// Uppercase and lowercase are considered different characters.
func rotEncipherRange(text string, key int, min, max rune) string {
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

// ROT5 is the ROT-5 symmetric cipher.
// For example, "1234567890" becomes "5678901234".
var ROT5 = &ROTRange{5, '0', '9'}

// ROT13 is the ROT-13 symmetric cipher.
// For example, "ABCDEFGHIJKLM" becomes "NOPQRSTUVWXYZ".
var ROT13 = Affine{1, 13, 1}

// ROT18 is the ROT-18 symmetric cipher.
// For example, "ABCXYZ012" becomes "STUFGHijk"
var ROT18 = rot18{}

type rot18 struct{}

func (r *rot18) Encipher(text string) string {
	return ROT13.Encipher(ROT5.Encipher(text))
}

func (r *rot18) Decipher(text string) string {
	return r.Encipher(text)
}

// ROT47 is the ROT-47 symmetric cipher.
// For example, "ABCabc" becomes "pqr234".
var ROT47 = &ROTRange{47, '!', '~'}
