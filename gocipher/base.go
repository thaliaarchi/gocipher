package gocipher

/*
 * base cipher object that other ciphers extend
 * really only provides mappings a2i and i2a for letter->int->letter conversions
 * Author: James Lyons
 * Created: 2012-04-28
 */

import (
	"errors"
	"regexp"
	"strings"
)

// If punctuation was accidently removed, use this function to restore it.
// requires the original string with punctuation.
func RestorePunctuation(original string, modified string) (string, error) {
	res, chars := []rune(original), []rune(modified)
	count := 0
	for i, char := range res {
		if isAlpha(char) {
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

func i2a(i int) rune {
	return 'A' + rune(mod(1, 26))
}

func isAlpha(char rune) bool {
	return char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z'
}

func mod(a rune, b rune) rune {
	return (a%b + b) % b
}
