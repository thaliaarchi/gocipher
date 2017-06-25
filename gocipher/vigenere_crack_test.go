package gocipher

import (
	"reflect"
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

	chunks = []string{"AC", "BD"}
	expected = "ABCD"
	actual = joinChunks(chunks)
	if expected != actual {
		t.Errorf("Expected %q, but got %q (chunks: %v)", expected, actual, chunks)
	}
}

func TestCartesianProduct(t *testing.T) {
	input := [][]string{
		[]string{"ABC", "abc", "123"},
		[]string{"DEF", "def", "456"},
		[]string{"GHI", "ghi", "789"}}
	expected := [][]string{
		[]string{"ABC", "DEF", "GHI"}, []string{"ABC", "DEF", "ghi"}, []string{"ABC", "DEF", "789"},
		[]string{"ABC", "def", "GHI"}, []string{"ABC", "def", "ghi"}, []string{"ABC", "def", "789"},
		[]string{"ABC", "456", "GHI"}, []string{"ABC", "456", "ghi"}, []string{"ABC", "456", "789"},
		[]string{"abc", "DEF", "GHI"}, []string{"abc", "DEF", "ghi"}, []string{"abc", "DEF", "789"},
		[]string{"abc", "def", "GHI"}, []string{"abc", "def", "ghi"}, []string{"abc", "def", "789"},
		[]string{"abc", "456", "GHI"}, []string{"abc", "456", "ghi"}, []string{"abc", "456", "789"},
		[]string{"123", "DEF", "GHI"}, []string{"123", "DEF", "ghi"}, []string{"123", "DEF", "789"},
		[]string{"123", "def", "GHI"}, []string{"123", "def", "ghi"}, []string{"123", "def", "789"},
		[]string{"123", "456", "GHI"}, []string{"123", "456", "ghi"}, []string{"123", "456", "789"}}
	actual := cartesianProduct(input)
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
