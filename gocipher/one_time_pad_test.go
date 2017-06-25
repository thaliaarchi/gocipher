package gocipher

import "testing"

func TestOneTimePadEncrypt(t *testing.T) {
	text := "Hello"
	key := "XMCKL"
	expected := "Eqnvz"
	actual, err := OneTimePadEncrypt(text, key)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	if expected != actual {
		t.Errorf("Expected %q, but got %q (text: %q, key: %q)", expected, actual, text, key)
	}
}

func TestOneTimePadDecrypt(t *testing.T) {
	text := "Eqnvz"
	key := "XMCKL"
	expected := "Hello"
	actual, err := OneTimePadDecrypt(text, key)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	if expected != actual {
		t.Errorf("Expected %q, but got %q (text: %q, key: %q)", expected, actual, text, key)
	}
}
