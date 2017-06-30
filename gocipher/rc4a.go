package gocipher

/*
 * RC4A cipher
 * A variant of RC4 proposed by Souradyuti Paul and Bart Preneel.
 */

type RC4A struct {
	key string
}

func NewRC4A(key string) *RC4A {
	return &RC4A{key}
}

// Encipher enciphers string using RC4A according to key
func (r *RC4A) Encipher(text string) string {
	n := 256
	s1 := rc4KSA(r.key, n)
	s2 := rc4KSA(r.key, n)
	i := 0
	j1 := 0
	j2 := 0
	res := []byte(text)
	for y := 0; y < len(text); y++ {
		i = (i + 1) % n
		j1 = (j1 + s1[i]) % n
		s1[i], s1[j1] = s1[j1], s1[i]
		res[y] ^= byte(s2[(s1[i]+s1[j1])%n])
		y++
		j2 = (j2 + s2[i]) % n
		s2[i], s2[j2] = s2[j2], s2[i]
		res[y] ^= byte(s1[(s2[i]+s2[j2])%n])
	}
	return string(res)
}

// Decipher deciphers string using RC4A according to key
func (r *RC4A) Decipher(text string) string {
	return r.Encipher(text)
}
