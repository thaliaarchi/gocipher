package gocipher

/*
 * Trithemius cipher
 */

// TrithemiusEncipher enciphers string using Trithemius cipher.
func TrithemiusEncipher(text string) string {
	return mapAlpha(text, func(i, char int) int {
		return char + i
	})
}

// TrithemiusDecipher deciphers string using Trithemius cipher.
func TrithemiusDecipher(text string) string {
	return mapAlpha(text, func(i, char int) int {
		return char - i
	})
}
