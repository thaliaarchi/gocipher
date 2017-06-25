package gocipher

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplitChunks(t *testing.T) {
	text := "ABCDEFGHIJKL"
	size := 4
	expected := []string{"ADGJ", "BEHK", "CFIL"}
	actual := splitChunks(text, size)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v (text: %q, size: %d)", expected, actual, text, size)
	}

	text = "ABCD"
	size = 2
	expected = []string{"AC", "BD"}
	actual = splitChunks(text, size)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v (text: %q, size: %d)", expected, actual, text, size)
	}
}

func TestJoinChunks(t *testing.T) {
	chunks := []string{"ADGJ", "BEHK", "CFIL"}
	expected := "ABCDEFGHIJKL"
	actual := joinChunks(chunks)
	if expected != actual {
		t.Errorf("Expected %q, but got %q (chunks: %v)", expected, actual, chunks)
	}
}

func TestCartesian(t *testing.T) {
	input := [][]string{
		[]string{"ABC", "abc", "123"},
		[]string{"DEF", "def", "456"},
		[]string{"GHI", "ghi", "789"}}
	expected := []string{
		"ABC,DEF,GHI", "ABC,DEF,ghi", "ABC,DEF,789",
		"ABC,def,GHI", "ABC,def,ghi", "ABC,def,789",
		"ABC,456,GHI", "ABC,456,ghi", "ABC,456,789",
		"abc,DEF,GHI", "abc,DEF,ghi", "abc,DEF,789",
		"abc,def,GHI", "abc,def,ghi", "abc,def,789",
		"abc,456,GHI", "abc,456,ghi", "abc,456,789",
		"123,DEF,GHI", "123,DEF,ghi", "123,DEF,789",
		"123,def,GHI", "123,def,ghi", "123,def,789",
		"123,456,GHI", "123,456,ghi", "123,456,789"}
	output := cartesian(input)
	actual := make([]string, len(output))
	for i, chunks := range output {
		actual[i] = strings.Join(chunks, ",")
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v (input: %v)", expected, actual, input)
	}
}

func TestVigenerePossibilities(t *testing.T) {
	text := "ABCD"
	keyLength := 2
	expected := []string{
		"ABCD", "ACCE", "ADCF", "AECG", "AFCH", "AGCI", "AHCJ", "AICK", "AJCL", "AKCM", "ALCN", "AMCO", "ANCP",
		"AOCQ", "APCR", "AQCS", "ARCT", "ASCU", "ATCV", "AUCW", "AVCX", "AWCY", "AXCZ", "AYCA", "AZCB", "AACC",
		"BBDD", "BCDE", "BDDF", "BEDG", "BFDH", "BGDI", "BHDJ", "BIDK", "BJDL", "BKDM", "BLDN", "BMDO", "BNDP",
		"BODQ", "BPDR", "BQDS", "BRDT", "BSDU", "BTDV", "BUDW", "BVDX", "BWDY", "BXDZ", "BYDA", "BZDB", "BADC"} // ... 13 times more
	expectedLen := 676 // 26^26
	actual := VigenerePossibilities(text, keyLength)
	hasError := false
	if expectedLen != len(actual) {
		t.Errorf("Expected length %d, but got %d", expectedLen, len(actual))
	}
	for i, value := range expected {
		if value != actual[i] {
			t.Errorf("Expected %q at index %d, but got %q", value, i, actual[i])
			hasError = true
		}
	}
	if hasError {
		t.Errorf("Expected %v, but got %v (text: %q, keyLength: %d)",
			expected, actual, text, keyLength)
	}
}
