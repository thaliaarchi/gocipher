package gocipher

// LetterNumberEncrypt - Converts letters to the corresponding number.
// e.g. "ABC...XYZ" becomes []int{1, 2, 3 ... 24, 25, 26}
func LetterNumberEncrypt(text string) []int {
	runes := []rune(text)
	numbers := make([]int, len(runes))
	for i, rune := range runes {
		if rune >= 'A' && rune <= 'Z' {
			numbers[i] = int(rune - 'A' + 1)
		} else if rune >= 'a' && rune <= 'z' {
			numbers[i] = int(rune - 'a' + 1)
		}
	}
	return numbers
}

// LetterNumberDecrypt - Converts numbers to the corresponding letter.
// e.g. []int{1, 2, 3 ... 24, 25, 26} becomes "ABC...XYZ"
func LetterNumberDecrypt(numbers []int) string {
	runes := make([]rune, len(numbers))
	for i, number := range numbers {
		runes[i] = rune(number + 'A' - 1)
	}
	return string(runes)
}
