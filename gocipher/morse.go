package gocipher

import (
	"errors"
	"sort"
	"strings"
)

/*
 * Morse code
 * https://morsecode.scphillips.com/translator.html
 */

type Morse struct {
	textToMorse  map[string]string
	morseToText  map[string]string
	textLengths  []int // For multi-character tokens like prosigns and CH
	morseLengths []int // For multi-character Morse tokens in Wabun
}

func NewMorse(alphabets ...MorseAlphabet) *Morse {
	textToMorse := map[string]string{}
	morseToText := map[string]string{}
	textLengthMap := map[int]bool{}
	morseLengthMap := map[int]bool{}
	textLengths := []int{}
	morseLengths := []int{}

	// Load alphabet mappings. If none provided, use International Morse
	mappings := [][]string{{" ", "/"}}
	if len(alphabets) < 1 {
		mappings = append(mappings, morseInternational...)
	}
	for _, alphabet := range alphabets {
		mappings = append(mappings, morseAlphabets[alphabet]...)
	}

	// Add each character, morse translation, and the lengths to maps.
	for _, item := range mappings {
		chars := item[:len(item)-1]
		morse := item[len(item)-1]
		for _, char := range chars {
			char = strings.ToUpper(char)
			textToMorse[char] = morse
			textLengthMap[len([]rune(char))] = true
		}
		if _, exists := morseToText[morse]; !exists {
			morseToText[morse] = chars[0]
			length := len(strings.Split(morse, " "))
			morseLengthMap[length] = true
		}
	}

	// Convert length maps to descending sorted slices.
	for length := range textLengthMap {
		if length > 0 {
			textLengths = append(textLengths, length)
		}
	}
	for length := range morseLengthMap {
		if length > 0 {
			morseLengths = append(morseLengths, length)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(textLengths)))
	sort.Sort(sort.Reverse(sort.IntSlice(morseLengths)))

	return &Morse{textToMorse, morseToText, textLengths, morseLengths}
}

// Encode converts text into Morse code
func (m *Morse) Encode(text string) (string, error) {
	return morseFunc(tidyMorseText(text), m.textToMorse, m.textLengths, "", " ")
}

// Decode converts Morse code into text
func (m *Morse) Decode(morse string) (string, error) {
	return morseFunc(tidyMorse(morse), m.morseToText, m.morseLengths, " ", "")
}

func morseFunc(text string, tokenMap map[string]string, lengths []int, splitSep, joinSep string) (string, error) {
	if text == "" {
		return "", nil
	}
	tokens := strings.Split(text, splitSep)
	res := []string{}
	err := []string{}
	hasError := false
	for i := 0; i < len(tokens); i++ {
		found := false
		// Check each token length starting at the largest.
		for _, length := range lengths {
			if i+length > len(tokens) {
				continue
			}
			token := strings.Join(tokens[i:i+length], splitSep)
			if char, ok := tokenMap[token]; ok {
				res = append(res, char)
				err = append(err, token)
				i += length - 1
				found = true
				break
			}
		}
		if !found {
			res = append(res, "#")
			err = append(err, "#"+tokens[i]+"#")
			hasError = true
		}
	}
	result := strings.Join(res, joinSep)
	if hasError {
		return result, errors.New("error in input: " + strings.Join(err, splitSep))
	}
	return result, nil
}

// MorseFormatBullets replaces the .- characters in Morse text with •– so the characters are vertially centered.
func MorseFormatBullets(morse string) string {
	morse = strings.Replace(morse, "-", "–", -1)
	morse = strings.Replace(morse, ".", "•", -1)
	return morse
}

// MorseFormatSpoken replaces dots and dashes with dits and dahs in Morse text.
// https://en.wikipedia.org/wiki/Morse_code#Spoken_representation
func MorseFormatSpoken(morse string) string {
	words := strings.Split(morse, " / ")
	for i, word := range words {
		letters := strings.Split(word, " ")
		for j, letter := range letters {
			chars := []rune(letter)
			syllables := make([]string, len(letter))
			for k, char := range chars {
				if char == '-' {
					syllables[k] = "dah"
				} else if char == '.' {
					if k == len(chars)-1 {
						syllables[k] = "dit"
					} else {
						syllables[k] = "di"
					}
				}
			}
			letters[j] = strings.Join(syllables, "-")
		}
		word = strings.Join(letters, " ")
		chars := []rune(word)
		words[i] = strings.ToUpper(string(chars[0])) + string(chars[1:])
	}
	return strings.Join(words, ", ") + "."
}

func tidyMorseText(text string) string {
	text = strings.ToUpper(text)
	text = strings.TrimSpace(text)
	text = replacePattern(text, "\\s+", " ")
	return text
}

func tidyMorse(morse string) string {
	morse = strings.Replace(morse, "|", "/", -1)
	morse = strings.Replace(morse, "/", " / ", -1)
	morse = replacePattern(morse, "\\s+", " ")
	morse = replacePattern(morse, "(/ )+/", "/")
	morse = strings.Replace(morse, "_", "-", -1)
	morse = strings.TrimSpace(morse)
	return morse
}
