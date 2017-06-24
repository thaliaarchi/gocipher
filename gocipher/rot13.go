package gocipher

/*
 * ROT-13 cipher
 */

// Rot13Encipher encipher string using rot13 cipher
func Rot13Encipher(text string) string {
	return CaesarEncipher(text, 13)
}

// Rot13Decipher decipher string using rot13 cipher
func Rot13Decipher(text string) string {
	return Rot13Encipher(text)
}
