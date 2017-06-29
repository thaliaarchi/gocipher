package gocipher

/*
 * Atbash cipher
 */

// AtbashEncipher enciphers string using Atbash cipher.
func AtbashEncipher(text string) string {
	return mapAlpha(text, func(i, char int) int {
		return 25 - char
	})
}

// AtbashDecipher deciphers string using the Atbash cipher.
func AtbashDecipher(text string) string {
	return AtbashEncipher(text)
}
