package gocipher

import "strings"

/*
 * Keyed Caesar cipher
 */

type CaesarKeyed struct {
	alphabet string
	shift    int
}

func NewCaesarKeyed(key string, shift int) *CaesarKeyed {
	alphabet := KeyedAlphabetRange(strings.ToUpper(key), 'A', 'Z')
	return &CaesarKeyed{alphabet, shift}
}

// Encipher enciphers string using keyed Caesar cipher according to key.
func (c *CaesarKeyed) Encipher(text string) string {
	return caesarKeyedEncipher(text, c.alphabet, c.shift)
}

// Decipher deciphers string using keyed Caesar cipher according to key.
func (c *CaesarKeyed) Decipher(text string) string {
	return caesarKeyedEncipher(text, c.alphabet, -c.shift)
}

func caesarKeyedEncipher(text, alphabet string, shift int) string {
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
