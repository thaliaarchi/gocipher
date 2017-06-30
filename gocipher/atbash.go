package gocipher

/*
 * Atbash cipher
 */

type Atbash struct{}

func NewAtbash() *Atbash {
	return &Atbash{}
}

// Encipher enciphers string using Atbash cipher.
func (a *Atbash) Encipher(text string) string {
	return mapAlpha(text, func(i, char int) int {
		return 25 - char
	})
}

// Decipher deciphers string using the Atbash cipher.
func (a *Atbash) Decipher(text string) string {
	return a.Encipher(text)
}
