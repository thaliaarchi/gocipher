package gocipher

import (
	"errors"
	"strconv"
	"strings"
)

/*
 * Polybius square cipher
 */

type Polybius struct {
	key   string
	size  int
	chars string
}

func NewPolybius(key string, size int, chars string) (*Polybius, error) {
	key = strings.ToUpper(key)
	chars = strings.ToUpper(chars)[:size]
	if len(key) != size*size {
		return nil, errors.New("key must have length of size*size, has length " + strconv.Itoa(len(key)))
	}
	if len(chars) != size {
		return nil, errors.New("chars must have length of size, has length " + strconv.Itoa(len(chars)))
	}
	return &Polybius{key, size, chars}, nil
}

// Encipher enciphers string using Polybius square cipher according to initialised key.
func (p *Polybius) Encipher(text string) string {
	chars := []rune(strings.ToUpper(text))
	res := ""
	for _, char := range chars {
		res += p.encipherChar(char)
	}
	return res
}

// Decipher deciphers string using Polybius square cipher according to initialised key.
func (p *Polybius) Decipher(text string) string {
	chars := []rune(strings.ToUpper(text))
	res := ""
	for i := 0; i < len(chars); i += 2 {
		res += p.decipherPair(chars[i : i+2])
	}
	return res
}

func (p *Polybius) encipherChar(char rune) string {
	index := strings.IndexRune(p.key, char)
	row := index / p.size
	col := index % p.size
	chars := []rune(p.chars)
	return string([]rune{chars[row], chars[col]})
}

func (p *Polybius) decipherPair(pair []rune) string {
	row := strings.IndexRune(p.chars, pair[0])
	col := strings.IndexRune(p.chars, pair[1])
	return string([]rune(p.key)[row*p.size+col])
}
