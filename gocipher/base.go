package gocipher

import (
	"errors"
	"regexp"
	"strings"
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

func RemovePunctuation(text string) string {
	return removePattern(text, "[^A-Z]")
}

func removePattern(text string, pattern string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(strings.ToUpper(text), "")
}

func a2i(char rune) int {
	if char >= 'A' && char <= 'Z' {
		return int(char - 'A')
	} else if char >= 'a' && char <= 'z' {
		return int(char - 'a')
	}
	return -1
}

func i2a(i int, isUpper bool) rune {
	if !isUpper {
		return 'a' + rune(mod(i, 26))
	}
	return 'A' + rune(mod(i, 26))
}

func isAlpha(char rune) (bool, bool) {
	isUpper := char >= 'A' && char <= 'Z'
	isAlpha := isUpper || char >= 'a' && char <= 'z'
	return isAlpha, isUpper
}

func mod(a int, b int) int {
	return (a%b + b) % b
}

func modRune(a rune, b rune) rune {
	return (a%b + b) % b
}
