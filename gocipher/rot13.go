package gocipher

/*
 * implements rot13 cipher
 * Author: James Lyons
 * Created: 2014-02-09
 * http://www.practicalcryptography.com/ciphers/rot13-cipher/
 */

// Rot13Encipher - Encipher string using rot13 cipher
func Rot13Encipher(text string) string {
	return CaesarEncipher(text, 13)
}

// Rot13Decipher - Decipher string using rot13 cipher
func Rot13Decipher(text string) string {
	return Rot13Encipher(text)
}
