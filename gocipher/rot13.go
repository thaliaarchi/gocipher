package gocipher

/*
 * Implements ROT-13 cipher
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
