package gocipher

import "testing"

func TestRot5Encipher(t *testing.T) {
	text := "0123456789"
	expected := "5678901234"
	actual := Rot5Encipher(text)
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestRot5Decipher(t *testing.T) {
	text := "5678901234"
	expected := "0123456789"
	actual := Rot5Decipher(text)
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestRot13Encipher(t *testing.T) {
	var text = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var expected = "nopqrstuvwxyzabcdefghijklmNOPQRSTUVWXYZABCDEFGHIJKLM"
	actual := Rot13Encipher(text)
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestRot13Decipher(t *testing.T) {
	var text = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var expected = "nopqrstuvwxyzabcdefghijklmNOPQRSTUVWXYZABCDEFGHIJKLM"
	actual := Rot13Decipher(text)
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestRot47Encipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "23456789:;<=>?@ABCDEFGHIJKpqrstuvwxyz{|}~!\"#$%&'()*+"
	actual := Rot47Encipher(text)
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestRot47Decipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := "23456789:;<=>?@ABCDEFGHIJKpqrstuvwxyz{|}~!\"#$%&'()*+"
	actual := Rot47Decipher(text)
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}
