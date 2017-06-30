package gocipher

/*
 * Affine cipher
 */

import "errors"

// Affine is a key for an Affine cipher
type Affine struct {
	a, b, inva int
}

// NewAffine creates an Affine.
// `a` is the multiplicative part of the key (allowable values are: 1, 3, 5, 7, 9, 11, 15, 17, 19, 21, 23, and 25).
// `b` is the additive part of the key (allowable values are integers 0-25).
func NewAffine(a, b int) (*Affine, error) {
	inva := -1
	for i := 1; i < 26; i += 2 {
		if mod(a*i, 26) == 1 {
			inva = i
		}
	}
	if 0 > inva || inva > 25 {
		return nil, errors.New("invalid key: a=" + string(a) + ", no inverse exists (mod 26)")
	}
	return &Affine{a, b, inva}, nil
}

// Encipher enciphers string using Affine cipher according to key.
func (key *Affine) Encipher(text string) string {
	return mapAlpha(text, func(i, char int) int {
		return key.a*char + key.b
	})
}

// Decipher deciphers string using Affine cipher according to key.
func (key *Affine) Decipher(text string) string {
	return mapAlpha(text, func(i, char int) int {
		return key.inva * (char - key.b)
	})
}
