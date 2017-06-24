package gocipher

import "testing"

func TestKeyedAlphabet(t *testing.T) {
	key := "Hello, World!"
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "HELOWRDABCFGIJKMNPQSTUVXYZ"
	actual := KeyedAlphabet(key, alphabet)
	if expected != actual {
		t.Errorf("Expected %q, but got %q (key: %q, alphabet: %q)", expected, actual, key, alphabet)
	}
}

func TestKeyedAlphabetRange(t *testing.T) {
	key := "HELLO, WORLD!"
	min, max := 'A', 'Z'
	expected := "HELOWRDABCFGIJKMNPQSTUVXYZ"
	actual := KeyedAlphabetRange(key, min, max)
	if expected != actual {
		t.Errorf("Expected %q, but got %q (key: %q, min: %v, max %v)", expected, actual, key, min, max)
	}
}
