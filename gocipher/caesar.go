package gocipher

/*
 * Caesar cipher
 */

type Caesar struct {
	key int
}

func NewCaesar(key int) *Caesar {
	return &Caesar{key}
}

// Encipher enciphers string using Caesar cipher according to key.
func (c *Caesar) Encipher(text string) string {
	return caesarEncipher(text, c.key)
}

// Decipher deciphers string using Caesar cipher according to key.
func (c *Caesar) Decipher(text string) string {
	return caesarEncipher(text, -c.key)
}

func caesarEncipher(text string, key int) string {
	if key == 0 {
		return text
	}
	return mapAlpha(text, func(i, char int) int {
		return char + key
	})
}
