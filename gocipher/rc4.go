package gocipher

// RC4 is an instance of the RC4 cipher.
type RC4 struct {
	key string
}

func NewRC4(key string) *RC4 {
	return &RC4{key}
}

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

// RC4A is an instance of the RC4A cipher, a variant of RC4 proposed by
// Souradyuti Paul and Bart Preneel.
type RC4A struct {
	key string
}

func NewRC4A(key string) *RC4A {
	return &RC4A{key}
}

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

func (r *RC4A) Decipher(text string) string {
	return r.Encipher(text)
}
