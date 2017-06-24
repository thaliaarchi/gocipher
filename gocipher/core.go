package gocipher

import (
	"errors"
	"regexp"
	"strings"
)

// KeyedAlphabet creates an alphabet starting with a keyword.
// e.g. "Hello, World!", "ABCDEFGHIJKLMNOPQRSTUVWXYZ" becomes "HELOWRDABCFGIJKMNPQSTUVXYZ"
func KeyedAlphabet(key string, alphabet string) string {
	alphabet = strings.ToUpper(alphabet)
	chars := []rune(strings.ToUpper(key) + alphabet)
	alpha := ""
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
	chars := []rune(key)
	alpha := ""
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

// RestorePunctuation - If punctuation was accidently removed, use this function to restore it.
// Requires the original string with punctuation.
func RestorePunctuation(original string, modified string) (string, error) {
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
	return removePattern(text, "[^A-Za-z]")
}

// removePattern removes any chars matching a Regular Expression from a string.
func removePattern(text string, pattern string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(text, "")
}

func mapAlpha(text string, f func(i int, char, a, z rune) rune) string {
	runes := []rune(text)
	for i, char := range runes {
		if char >= 'A' && char <= 'Z' {
			runes[i] = f(i, char, 'A', 'Z')
		} else if char >= 'a' && char <= 'z' {
			runes[i] = f(i, char, 'a', 'z')
		}
	}
	return string(runes)
}

// mod returns the modulus `a mod b` in the interval [0, b).
func mod(a int, b int) int {
	return (a%b + b) % b
}

// mod returns the modulus `a mod b` in the interval [0, b) for runes.
func modRune(a rune, b rune) rune {
	return (a%b + b) % b
}
