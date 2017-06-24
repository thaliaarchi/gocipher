package gocipher

/*
 * Atbash cipher
 */

// AtbashEncipher enciphers string using Atbash cipher.
func AtbashEncipher(text string) string {
	runes := []rune(text)
	for i, char := range runes {
		if char >= 'A' && char <= 'Z' {
			runes[i] = 'Z' - (char - 'A')
		} else if char >= 'a' && char <= 'z' {
			runes[i] = 'z' - (char - 'a')
		}
	}
	return string(runes)
}

// AtbashDecipher deciphers string using the Atbash cipher.
func AtbashDecipher(text string) string {
	return AtbashEncipher(text)
}
