package gocipher

/*
 * ROT-5 cipher
 */

// Rot5Encipher encipher string using ROT-5 cipher.
func Rot5Encipher(text string) string {
	shift := rune(5)
	runes := []rune(text)
	for i, char := range runes {
		if char >= '0' && char <= '9' {
			runes[i] = modRune(char+shift-'0', 10) + '0'
		}
	}
	return string(runes)
}

// Rot5Decipher decipher string using ROT-5 cipher.
func Rot5Decipher(text string) string {
	return Rot5Encipher(text)
}
