package gocipher

/*
 * ROT-5 cipher
 */

// Rot5Encipher enciphers string using ROT-5 cipher.
func Rot5Encipher(text string) string {
	return rot(text, '0', '9')
}

// Rot5Decipher deciphers string using ROT-5 cipher.
func Rot5Decipher(text string) string {
	return Rot5Encipher(text)
}

/*
 * ROT-13 cipher
 */

// Rot13Encipher enciphers string using ROT-13 cipher.
func Rot13Encipher(text string) string {
	return CaesarEncipher(text, 13)
}

// Rot13Decipher deciphers string using ROT-13 cipher.
func Rot13Decipher(text string) string {
	return Rot13Encipher(text)
}

/*
 * ROT-47 cipher
 */

// Rot47Encipher enciphers string using ROT-47 cipher.
func Rot47Encipher(text string) string {
	return rot(text, '!', '~')
}

// Rot47Decipher deciphers string using ROT-47 cipher.
func Rot47Decipher(text string) string {
	return Rot47Encipher(text)
}

func rot(text string, min, max rune) string {
	size := max - min + 1
	shift := rune(size / 2)
	runes := []rune(text)
	for i, char := range runes {
		if char >= min && char <= max {
			runes[i] = modRune(char+shift-min, size) + min
		}
	}
	return string(runes)
}
