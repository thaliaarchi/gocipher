package gocipher

/*
 * implements atbash cipher
 * Author: James Lyons
 * Created: 2014-02-09
 */

// AtbashEncipher -  Encipher string using Atbash cipher.
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

// AtbashDecipher - Decipher string using the Atbash cipher.
func AtbashDecipher(text string) string {
	return AtbashEncipher(text)
}
