package gocipher

import (
	"errors"
	"regexp"
	"strings"
)

/*
 * Morse code
 * https://morsecode.scphillips.com/translator.html
 * Mappings from: https://en.wikipedia.org/wiki/Morse_code
 */

var morseChars = [][]string{
	{"A", ".-"},
	{"B", "-..."},
	{"C", "-.-."},
	{"D", "-.."},
	{"E", "."},
	{"F", "..-."},
	{"G", "--."},
	{"H", "...."},
	{"I", ".."},
	{"J", ".---"},
	{"K", "-.-"},
	{"L", ".-.."},
	{"M", "--"},
	{"N", "-."},
	{"O", "---"},
	{"P", ".--."},
	{"Q", "--.-"},
	{"R", ".-."},
	{"S", "..."},
	{"T", "-"},
	{"U", "..-"},
	{"V", "...-"},
	{"W", ".--"},
	{"X", "-..-"},
	{"Y", "-.--"},
	{"Z", "--.."},
	{"0", "-----"},
	{"1", ".----"},
	{"2", "..---"},
	{"3", "...--"},
	{"4", "....-"},
	{"5", "....."},
	{"6", "-...."},
	{"7", "--..."},
	{"8", "---.."},
	{"9", "----."},
	{".", ".-.-.-"},
	{",", "--..--"},
	{"?", "..--.."},
	{"'", ".----."},
	{"!", "-.-.--"}, // <KW>
	{"/", "-..-."},
	{"(", "-.--."},
	{")", "-.--.-"},
	{"&", ".-..."}, // <AS>, Not in ITU-R recommendation
	{":", "---..."},
	{";", "-.-.-."},
	{"=", "-...-"},
	{"+", ".-.-."}, // <AR>
	{"-", "-....-"},
	{"_", "..--.-"}, // Not in ITU-R recommendation
	{"\"", ".-..-."},
	{"$", "...-..-"}, // <SX>, Not in ITU-R recommendation
	{"@", ".--.-."},  // <AC>
	{" ", "/"},
}

var morseProsigns = [][]string{
	{"<AA>", ".-.-"},
	{"<AR>", ".-.-."},
	{"<AS>", ".-..."},
	{"<BK>", "-...-.-"},
	{"<BT>", "-...-"},
	{"<CL>", "-.-..-.."},
	{"<CT>", "-.-.-"},
	{"<HH>", "........"}, {"<EEEEEEEE>", "........"},
	{"<KN>", "-.--."},
	{"<NJ>", "-..---"}, {"<DO>", "-..---"},
	{"<SK>", "...-.-"},
	{"<SN>", "...-."}, {"<VE>", "...-."},
	{"<VA>", "...-.-"},
	{"<SOS>", "...---..."},
}

var morseNonEnglish = [][]string{
	{"À", ".--.-"}, {"Å", ".--.-"},
	{"Ä", ".-.-"}, {"Æ", ".-.-"}, {"Ą", ".-.-"},
	{"Ć", "-.-.."}, {"Ĉ", "-.-.."}, {"Ç", "-.-.."},
	{"CH", "----"}, {"Ĥ", "----"}, {"Š", "----"},
	{"Đ", "..-.."}, {"É", "..-.."}, {"Ę", "..-.."},
	{"Ð", "..--."},
	{"È", ".-..-"}, {"Ł", ".-..-"},
	{"Ĝ", "--.-."},
	{"Ĵ", ".---."},
	{"Ń", "..-.."}, {"Ñ", "..-.."},
	{"Ó", "---."}, {"Ö", "---."}, {"Ø", "---."},
	{"Ś", "...-..."},
	{"Ŝ", "...-."}, // <SN> <VE>
	{"Þ", ".--.."},
	{"Ü", "..--"}, {"Ŭ", "..--"},
	{"Ź", "--..-."},
	{"Ż", "--..-"},
}

type Morse struct {
	useProsigns bool
	textToMorse map[string]string
	morseToText map[string]string
}

func NewMorse(useProsigns bool, useNonEnglish bool) *Morse {
	var textToMorse = map[string]string{}
	var morseToText = map[string]string{}
	mappings := morseChars
	if useProsigns {
		mappings = append(mappings, morseChars...)
	}
	if useNonEnglish {
		mappings = append(mappings, morseNonEnglish...)
	}
	for _, pair := range mappings {
		char, morse := pair[0], pair[1]
		textToMorse[char] = morse
		if _, exists := morseToText[morse]; !exists {
			morseToText[morse] = char
		}
	}
	return &Morse{useProsigns, textToMorse, morseToText}
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
		tokenLength := 1
		if m.useProsigns {
			prosign := regexp.MustCompile("^<...?>").FindString(text)
			if prosign != "" {
				tokenLength = len(prosign)
			}
		}
		tokens = append(tokens, text[:tokenLength])
		text = text[tokenLength:]
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
