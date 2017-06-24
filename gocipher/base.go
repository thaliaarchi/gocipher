package gocipher

import (
	"errors"
	"regexp"
)

// RestorePunctuation - If punctuation was accidently removed, use this function to restore it.
// Requires the original string with punctuation.
func RestorePunctuation(original string, modified string) (string, error) {
	res, chars := []rune(original), []rune(modified)
	count := 0
	for i, char := range res {
		if isAlpha, _ := isAlpha(char); isAlpha {
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

// a2i returns the number 0-25 corresponding to
func alphaIndex(char rune) int {
	if char >= 'A' && char <= 'Z' {
		return int(char - 'A')
	} else if char >= 'a' && char <= 'z' {
		return int(char - 'a')
	}
	return -1
}

func indexToRune(i int, isUpper bool) rune {
	if !isUpper {
		return 'a' + rune(mod(i, 26))
	}
	return 'A' + rune(mod(i, 26))
}

// isAlpha returns whether a rune is alphabetical and whether a rune is upper case.
// Only considers A-Z and a-z.
func isAlpha(char rune) (bool, bool) {
	isUpper := char >= 'A' && char <= 'Z'
	isAlpha := isUpper || char >= 'a' && char <= 'z'
	return isAlpha, isUpper
}

// mod returns the modulus `a mod b` in the interval [0, b).
func mod(a int, b int) int {
	return (a%b + b) % b
}

// mod returns the modulus `a mod b` in the interval [0, b) for runes.
func modRune(a rune, b rune) rune {
	return (a%b + b) % b
}
