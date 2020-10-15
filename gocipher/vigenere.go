package gocipher

import (
	"math"
	"strings"

	"github.com/andrewarchi/gocipher/mod"
)

var englishFreq = [26]float64{
	0.080406052, 0.014846488, 0.033437737, 0.038169583, 0.124920625,
	0.024031234, 0.018693758, 0.050533014, 0.075692775, 0.001587737,
	0.005405135, 0.04068986, 0.025117606, 0.072336292, 0.076406929,
	0.02135891, 0.001204689, 0.062794207, 0.065127666, 0.092755648,
	0.027297018, 0.010532516, 0.016756642, 0.002348569, 0.016649801, 0.000899507}

// Vigenere is a Vigenère cipher key.
type Vigenere struct {
	shifts []int8
}

// NewVigenere constructs a Vigenère cipher key.
func NewVigenere(key []int8) *Vigenere {
	shifts := make([]int8, len(key))
	for i := range key {
		shifts[i] = mod.ModInt8(key[i], 26)
	}
	return &Vigenere{shifts}
}

// Encrypt encrypts the plaintext with the Vigenère cipher.
func (v *Vigenere) Encrypt(plain string) string {
	return v.crypt(plain, false)
}

// Decrypt decrypts the plaintext with the Vigenère cipher.
func (v *Vigenere) Decrypt(cipher string) string {
	return v.crypt(cipher, true)
}

func (v *Vigenere) crypt(text string, decrypt bool) string {
	var b strings.Builder
	b.Grow(len(text))
	j := 0
	for i := 0; i < len(text); i++ {
		ch := text[i]
		if !isLetter(ch) {
			b.WriteByte(ch)
			continue
		}
		n := toOrd(ch)
		if decrypt {
			n += byte(26 - v.shifts[j])
		} else {
			n += byte(v.shifts[j])
		}
		b.WriteByte(n%26 + getCase(ch))
		j = (j + 1) % len(v.shifts)
	}
	return b.String()
}

// VigenereCrack finds the key with the letter frequencies closest to
// English.
func VigenereCrack(cipher string, keyLen int) []int8 {
	freqs := make([][26]float64, keyLen)
	j := 0
	cipherLen := 0
	for i := 0; i < len(cipher); i++ {
		ch := cipher[i]
		if !isLetter(ch) {
			continue
		}
		freqs[j][toOrd(ch)]++
		j = (j + 1) % keyLen
		cipherLen++
	}

	for k := 0; k < keyLen; k++ {
		n := float64((cipherLen + keyLen - k) / keyLen)
		for i := 0; i < 26; i++ {
			freqs[k][i] /= n
		}
	}

	key := make([]int8, keyLen)
	for k := 0; k < keyLen; k++ {
		best := math.MaxFloat64
		for i := 0; i < 26; i++ {
			chi := chiSquared(&freqs[k], &englishFreq, i)
			if chi < best {
				best = chi
				key[k] = int8(i)
			}
		}
	}
	return key
}

func chiSquared(observeFreq, expectFreq *[26]float64, shift int) float64 {
	var sum float64
	for i := 0; i < 26; i++ {
		f := observeFreq[(i+shift)%26] - expectFreq[i]
		sum += f * f / expectFreq[i]
	}
	return sum
}

func isLetter(ch byte) bool {
	upper := toUpper(ch)
	return 'A' <= upper && upper <= 'Z'
}

func toUpper(ch byte) byte {
	return ch &^ 0x20 // a-z -> A-Z
}

func toOrd(ch byte) byte {
	return (ch - 1) & 0x1f // a-z or A-Z -> 0-25
}

// getCase returns A for uppercase, a for lowercase, invalid otherwise.
func getCase(ch byte) byte {
	return ch&^0x1f + 1
}
