package gocipher

import "testing"

func TestFracMorseEncipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "FJHDVGSLMSCCEQFDHQHOEHKTGCPFAPOJQEWCFJHDVGSLMSCCEQFDHQHOEHKTGCPFAPOJQEWC"
	key := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	f, err := NewFracMorse(key)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	actual, err := f.Encipher(text)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	if expected != actual {
		t.Errorf("Expected %q, but got %q (key: %q)", expected, actual, key)
	}
}

func TestFracMorseDecipher(t *testing.T) {
	text := "FJHDVGSLMSCCEQFDHQHOEHKTGCPFAPOJQEWCFJHDVGSLMSCCEQFDHQHOEHKTGCPFAPOJQEWC"
	expected := "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	key := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	f, err := NewFracMorse(key)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	actual, err := f.Decipher(text)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	if expected != actual {
		t.Errorf("Expected %q, but got %q (key: %q)", expected, actual, key)
	}
}
