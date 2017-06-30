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
	n := 256
	s := rc4KSA(r.key, n)
	i := 0
	j := 0
	res := []byte(text)
	for y := 0; y < len(text); y++ {
		i = (i + 1) % n
		j = (j + s[i]) % n
		s[i], s[j] = s[j], s[i]
		res[y] ^= byte(s[(s[i]+s[j])%n])
	}
	return string(res)
}

// Decipher deciphers string using RC4 according to key
func (r *RC4) Decipher(text string) string {
	return r.Encipher(text)
}

// rc4KSA is the key-scheduling algorithm (KSA) for RC4.
// Generates a state array based on the key.
func rc4KSA(key string, n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}
	j := 0
	for i := 0; i < n; i++ {
		j = (j + s[i] + int(key[i%len(key)])) % n
		s[i], s[j] = s[j], s[i]
	}
	return s
}
