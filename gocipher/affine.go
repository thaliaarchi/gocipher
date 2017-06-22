package gocipher

/*
 * implements affine substitution cipher
 * Author: James Lyons
 * Created: 2012-04-28
 * http://www.practicalcryptography.com/ciphers/affine-cipher/
 */

import "errors"

/*The Affine Cipher has two components to the key, numbers *a* and *b*.
This cipher encrypts a letter according to the following equation::

		c = (a*p + b)%26

where c is the ciphertext letter, p the plaintext letter.
*b* is an integer 0-25, *a* is an integer that has an inverse (mod 26).
Allowable values for *a* are: 1,3,5,7,9,11,15,17,19,21,23,25
For more info on the Affine cipher see
http://www.practicalcryptography.com/ciphers/affine-cipher/.

:param a: The multiplicative part of the key. Allowable values are: 1,3,5,7,9,11,15,17,19,21,23,25
:param b: The additive part of the key. Allowable values are integers 0-25
*/

type AffineKey struct {
	a, b, inva int
}

func NewAffineKey(a, b int) (*AffineKey, error) {
	inva := -1
	for i := 1; i < 26; i += 2 {
		if mod(a*i, 26) == 1 {
			inva = i
		}
	}
	if 0 > inva || inva > 25 {
		return nil, errors.New("invalid key: a=" + string(a) + ", no inverse exists (mod 26)")
	}
	return &AffineKey{a, b, inva}, nil
}

// AffineEncipher - Encipher string using affine cipher according to key.
func AffineEncipher(text string, key *AffineKey) string {
	runes := []rune(text)
	for i, char := range runes {
		if isAlpha, isUpper := isAlpha(char); isAlpha {
			runes[i] = i2a(key.a*a2i(char)+key.b, isUpper)
		}
	}
	return string(runes)
}

// AffineDecipher - Decipher string using affine cipher according to key.
func AffineDecipher(text string, key *AffineKey) string {
	runes := []rune(text)
	for i, char := range runes {
		if isAlpha, isUpper := isAlpha(char); isAlpha {
			runes[i] = i2a(key.inva*(a2i(char)-key.b), isUpper)
		}
	}
	return string(runes)
}
