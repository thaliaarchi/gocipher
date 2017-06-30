package gocipher

import "errors"

/*
 * Rail-fence (ZigZag) cipher
 * Algorithm adapted from http://www.practicalcryptography.com/ciphers/classical-era/rail-fence/
 */

// RailfenceKey is a key for a Rail-fence cipher
type Railfence struct {
	key int
}

// NewRailfence creates a Railfence.
func NewRailfence(rails int) (*Railfence, error) {
	if rails <= 0 {
		return nil, errors.New("key must be greater than zero")
	}
	return &Railfence{rails}, nil
}

// Encipher enciphers string using Rail-fence cipher according to key
func (r *Railfence) Encipher(text string) (string, error) {
	chars := []rune(text)
	rails := r.key - 1
	if rails > 2*len(chars)-1 {
		return "", errors.New("key is too large for the text length")
	}
	if rails == 0 {
		return text, nil
	}
	res := make([]rune, len(chars))
	pos := 0
	for row := 0; row < rails; row++ {
		skip := 2 * (rails - row)
		j := 0
		for i := row; i < len(chars); {
			res[pos] = chars[i]
			pos++
			if row == 0 || j%2 == 0 {
				i += skip
			} else {
				i += 2*rails - skip
			}
			j++
		}
	}
	for i := rails; i < len(chars); i += 2 * rails {
		res[pos] = chars[i]
		pos++
	}
	return string(res), nil
}

// Decipher deciphers string using Rail-fence cipher according to key
func (r *Railfence) Decipher(text string) (string, error) {
	text = RemovePunctuation(text)
	chars := []rune(text)
	if len(chars) < 1 {
		return "", errors.New("text must not be empty")
	}
	rails := r.key - 1
	if rails > 2*len(chars)-1 {
		return "", errors.New("key is too large for the text length")
	}
	if rails == 0 {
		return text, nil
	}
	res := make([]rune, len(chars))
	pos := 0
	for line := 0; line < rails; line++ {
		skip := 2 * (rails - line)
		j := 0
		for i := line; i < len(chars); {
			res[i] = chars[pos]
			pos++
			if line == 0 || j%2 == 0 {
				i += skip
			} else {
				i += 2*rails - skip
			}
			j++
		}
	}
	for i := rails; i < len(chars); i += 2 * rails {
		res[i] = chars[pos]
		pos++
	}
	return string(res), nil
}
