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
	var text = "nopqrstuvwxyzabcdefghijklmNOPQRSTUVWXYZABCDEFGHIJKLM"
	var expected = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	actual := Rot13Decipher(text)
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestRot18Encipher(t *testing.T) {
	var text = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var expected = "stuvwxyz0123456789abcdefghijklmnopqrSTUVWXYZ0123456789ABCDEFGHijklmnopqr"
	actual := Rot18Encipher(text)
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestRot18Decipher(t *testing.T) {
	var text = "stuvwxyz0123456789abcdefghijklmnopqrSTUVWXYZ0123456789ABCDEFGHijklmnopqr"
	var expected = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHijklmnopqrSTUVWXYZ0123456789" // Loses capitalization on numbers
	actual := Rot18Decipher(text)
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
