package gocipher

import "testing"

func TestTrithemiusEncipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "acegikmoqsuwyacegikmoqsuwyACEGIKMOQSUWYACEGIKMOQSUWY"
	actual := TrithemiusEncipher(text)
	if expected != actual {
		t.Errorf("Expected %q, but got %q (text: %q)", expected, actual, text)
	}
}

func TestTrithemiusDecipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "aaaaaaaaaaaaaaaaaaaaaaaaaaAAAAAAAAAAAAAAAAAAAAAAAAAA"
	actual := TrithemiusDecipher(text)
	if expected != actual {
		t.Errorf("Expected %q, but got %q (text: %q)", expected, actual, text)
	}
}
