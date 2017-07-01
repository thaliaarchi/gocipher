package gocipher

import (
	"errors"
	"strings"
)

/*
 * Morse code
 * https://morsecode.scphillips.com/translator.html
 */

type Morse struct {
	textToMorse map[string]string
	morseToText map[string]string
	prosigns    []string
}

func NewMorse(alphabets ...MorseAlphabet) *Morse {
	var textToMorse = map[string]string{}
	var morseToText = map[string]string{}
	var prosigns = []string{}
	mappings := [][]string{{" ", "/"}}
	if len(alphabets) == 0 {
		mappings = append(mappings, morseInternational...)
	}
	for _, alphabet := range alphabets {
		mappings = append(mappings, morseAlphabets[alphabet]...)
	}
	for _, item := range mappings {
		chars := item[:len(item)-1]
		morse := item[len(item)-1]
		if _, exists := morseToText[morse]; !exists {
			morseToText[morse] = chars[0]
		}
		for _, char := range chars {
			char = strings.ToUpper(char)
			textToMorse[char] = morse
			if len([]rune(char)) > 1 {
				prosigns = append(prosigns, char)
			}
		}
	}
	return &Morse{textToMorse, morseToText, prosigns}
}

// Encode converts text into Morse code
func (m *Morse) Encode(text string) (string, error) {
	text = tidyMorseText(text)
	if text == "" {
		return "", nil
	}
	morse := ""
	err := ""
	hasError := false
	tokens := []string{}
	for len(text) > 0 {
		textRunes := []rune(text)
		tokenLength := 1
		for _, prosign := range m.prosigns {
			prosignLength := len([]rune(prosign))
			if prosign == string(textRunes[:prosignLength]) {
				tokenLength = prosignLength
			}
		}
		tokens = append(tokens, string(textRunes[:tokenLength]))
		text = string(textRunes[tokenLength:])
	}
	for _, token := range tokens {
		if char, ok := m.textToMorse[token]; ok {
			morse += char + " "
			err += token
		} else {
			morse += "# "
			err += "#" + token + "#"
			hasError = true
		}
	}
	morse = morse[:len(morse)-1]
	if hasError {
		return morse, errors.New("error in input: " + err)
	}
	return morse, nil
}

// Decode converts Morse code into text
func (m *Morse) Decode(morse string) (string, error) {
	morse = tidyMorse(morse)
	if morse == "" {
		return "", nil
	}
	text := ""
	err := ""
	hasError := false
	tokens := strings.Split(morse, " ")
	for _, token := range tokens {
		if char, ok := m.morseToText[token]; ok {
			text += char
			err += token + " "
		} else {
			text += "#"
			err += "#" + token + "# "
			hasError = true
		}
	}
	if hasError {
		err = err[:len(err)-1]
		return text, errors.New("error in input: " + err)
	}
	return text, nil
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
