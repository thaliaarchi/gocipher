package gocipher

import (
	"errors"
)

/*
 * Fractionated Morse cipher
 */

var fracMorse = []string{
	"...", "..-", "../", ".-.", ".--", ".-/", "./.", "./-", ".//", "-..", "-.-", "-./", "--.",
	"---", "--/", "-/.", "-/-", "-//", "/..", "/.-", "/./", "/-.", "/--", "/-/", "//.", "//-"}

type FracMorse struct {
	key         string
	charToMorse map[rune]string
	morseToChar map[string]rune
}

func NewFracMorse(key string) (*FracMorse, error) {
	key = RemoveDuplicates(key)
	if len(key) < 26 {
		return nil, errors.New("key has fewer than 26 unique characters " + key)
	}
	charToMorse := map[rune]string{}
	morseToChar := map[string]rune{}
	chars := []rune(key)
	for i, char := range chars {
		morse := fracMorse[i]
		charToMorse[char] = morse
		morseToChar[morse] = char
	}
	return &FracMorse{key, charToMorse, morseToChar}, nil
}

// Encipher enciphers text using the fractionated Morse cipher acccording to the initialized key.
func (f *FracMorse) Encipher(text string) (string, error) {
	morse, err := NewMorse(false, false).Encode(text)
	if err != nil {
		return "", err
	}
	morse = replacePattern(morse, " / ", "//")
	morse = replacePattern(morse, " ", "/")
	if len(morse)%3 != 0 {
		morse = (morse + "//")[:3*(len(morse)/3+1)]
	}
	res := ""
	runes := []rune(morse)
	for i := 0; i < len(runes); i += 3 {
		frac := string(runes[i : i+3])
		if char, ok := f.morseToChar[frac]; ok {
			res += string(char)
		} else {
			return "", errors.New(frac + " does not exist in map")
		}
	}
	return res, nil
}

// Decipher deciphers text using the fractionated Morse cipher acccording to the initialized key.
func (f *FracMorse) Decipher(text string) (string, error) {
	morse := ""
	runes := []rune(text)
	for _, char := range runes {
		if frac, ok := f.charToMorse[char]; ok {
			morse += frac
		} else {
			return "", errors.New(string(char) + " does not exist in key " + f.key)
		}
	}
	morse = replacePattern(morse, "/", " ")
	morse = replacePattern(morse, "  ", " / ")
	text, err := NewMorse(false, false).Decode(morse)
	if err != nil {
		return "", err
	}
	return text, nil
}
