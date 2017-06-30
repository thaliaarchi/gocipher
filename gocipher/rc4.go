package gocipher

/*
 * RC4 cipher
 */

type RC4 struct {
	key string
}

func NewRC4(key string) *RC4 {
	return &RC4{key}
}

// Encipher enciphers string using RC4 according to key
func (r *RC4) Encipher(text string) string {
	s := make([]int, 256)
	for i := 0; i < 256; i++ {
		s[i] = i
	}
	j := 0
	for i := 0; i < 256; i++ {
		j = (j + s[i] + int(r.key[i%len(r.key)])) % 256
		s[i], s[j] = s[j], s[i]
	}
	i := 0
	j = 0
	res := []byte(text)
	for y := 0; y < len(text); y++ {
		i = (i + 1) % 256
		j = (j + s[i]) % 256
		s[i], s[j] = s[j], s[i]
		res[y] = byte(int(text[y]) ^ s[(s[i]+s[j])%256])
	}
	return string(res)
}

// Decipher deciphers string using RC4 according to key
func (r *RC4) Decipher(text string) string {
	return r.Encipher(text)
}
