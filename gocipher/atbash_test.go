package gocipher

import "testing"

func TestAtbashEncipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "zyxwvutsrqponmlkjihgfedcbaZYXWVUTSRQPONMLKJIHGFEDCBA"
	actual := AtbashEncipher(text)
	if expected != actual {
		t.Errorf("Expected %q, but got %q (text: %q)", expected, actual, text)
	}
}

func TestAtbashDecipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "zyxwvutsrqponmlkjihgfedcbaZYXWVUTSRQPONMLKJIHGFEDCBA"
	actual := AtbashDecipher(text)
	if expected != actual {
		t.Errorf("Expected %q, but got %q (text: %q)", expected, actual, text)
	}
}
