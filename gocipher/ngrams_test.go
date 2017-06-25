package gocipher

import "testing"

func TestGetUnigramEntropy(t *testing.T) {
	text := "HELLOWORLD"
	expected := 4.318919301709485
	actual := GetUnigramEntropy(text)
	if expected != actual {
		t.Errorf("Expected %v, but got %v (text: %q)", expected, actual, text)
	}
}

func TestGetBigramEntropy(t *testing.T) {
	text := "HELLOWORLD"
	expected := 7.799115199460557
	actual := GetBigramEntropy(text)
	if expected != actual {
		t.Errorf("Expected %v, but got %v (text: %q)", expected, actual, text)
	}
}
