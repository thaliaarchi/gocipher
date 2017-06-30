package gocipher

/*
 * Trithemius cipher
 */

type Trithemius struct{}

func NewTrithemius() *Trithemius {
	return &Trithemius{}
}

// Encipher enciphers string using Trithemius cipher.
func (t *Trithemius) Encipher(text string) string {
	return mapAlpha(text, func(i, char int) int {
		return char + i
	})
}

// Decipher deciphers string using Trithemius cipher.
func (t *Trithemius) Decipher(text string) string {
	return mapAlpha(text, func(i, char int) int {
		return char - i
	})
}
