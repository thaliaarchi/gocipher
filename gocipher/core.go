package gocipher

import (
	"crypto/rand"
	"errors"
	"math/big"
	"regexp"
	"strings"
)

// RandomKey creates a cryptographically secure pseudorandom key
func RandomKey(length int) (string, error) {
	maxNum := big.NewInt(26)
	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		nBig, err := rand.Int(rand.Reader, maxNum)
		if err != nil {
			return "", err
		}
		runes[i] = rune(nBig.Int64()) + 'A'
	}
	return string(runes), nil
}

// KeyedAlphabet creates an alphabet starting with a keyword.
// e.g. "Hello, World!", "ABCDEFGHIJKLMNOPQRSTUVWXYZ" becomes "HELOWRDABCFGIJKMNPQSTUVXYZ"
func KeyedAlphabet(key, alphabet string) string {
	alphabet = strings.ToUpper(alphabet)
	alpha := ""
	chars := []rune(strings.ToUpper(key) + alphabet)
	for _, char := range chars {
		if !strings.ContainsRune(alpha, char) && strings.ContainsRune(alphabet, char) {
			alpha += string(char)
		}
	}
	return alpha
}

// KeyedAlphabetRange creates an alphabet in a range of chars starting with a keyword.
// Uppercase and lowercase are considered different characters.
// e.g. "HELLO, WORLD!", 'A', 'Z' becomes "HELOWRDABCFGIJKMNPQSTUVXYZ"
func KeyedAlphabetRange(key string, min, max rune) string {
	alpha := ""
	chars := []rune(key)
	for _, char := range chars {
		if !strings.ContainsRune(alpha, char) && char >= min && char <= max {
			alpha += string(char)
		}
	}
	for i := min; i <= max; i++ {
		if !strings.ContainsRune(alpha, i) {
			alpha += string(i)
		}
	}
	return alpha
}

func RemoveDuplicates(text string) string {
	res := ""
	chars := []rune(text)
	for _, char := range chars {
		if !strings.ContainsRune(res, char) {
			res += string(char)
		}
	}
	return res
}

// RestorePunctuation - If punctuation was accidently removed, use this function to restore it.
// Requires the original string with punctuation.
func RestorePunctuation(original, modified string) (string, error) {
	res, chars := []rune(original), []rune(modified)
	count := 0
	for i, char := range res {
		if char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' {
			res[i] = chars[count]
			count++
		}
	}
	if count != len(modified) {
		return "", errors.New("Strings must have the same number of alphabetic chars")
	}
	return string(res), nil
}

// RemovePunctuation removes any non alphabetic chars from a string.
func RemovePunctuation(text string) string {
	return replacePattern(text, "[^A-Za-z]", "")
}

// replacePattern replaces any matches to a Regular Expression in a string.
func replacePattern(text, pattern, replace string) string {
	return regexp.MustCompile(pattern).ReplaceAllString(text, replace)
}

func mapAlpha(text string, f func(i, char int) int) string {
	runes := []rune(text)
	for i, char := range runes {
		if char >= 'A' && char <= 'Z' {
			runes[i] = rune(mod(f(i, int(char-'A')), 26)) + 'A'
		} else if char >= 'a' && char <= 'z' {
			runes[i] = rune(mod(f(i, int(char-'a')), 26)) + 'a'
		}
	}
	return string(runes)
}

// alphaIndex returns the number 0-25 corresponding to the letter
func alphaIndex(char rune) int {
	if char >= 'A' && char <= 'Z' {
		return int(char - 'A')
	} else if char >= 'a' && char <= 'z' {
		return int(char - 'a')
	}
	return -1
}
