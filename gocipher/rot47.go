package gocipher

/*
 * ROT-47 cipher
 */

// Rot47Encipher encipher string using ROT-47 cipher.
func Rot47Encipher(text string) string {
	shift := rune(47)
	runes := []rune(text)
	for i, char := range runes {
		if char >= '!' && char <= '~' {
			runes[i] = modRune(char+shift-'!', 94) + '!'
		}
	}
	return string(runes)
}

// Rot47Decipher decipher string using ROT-47 cipher.
func Rot47Decipher(text string) string {
	return Rot47Encipher(text)
}
