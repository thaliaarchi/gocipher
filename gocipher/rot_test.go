package gocipher

import "testing"
import "strings"

func TestRotEncipher(t *testing.T) {
	var text = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var expected = "STUVWXYZ0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789ABCDEFGHIJKLMNOPQR"
	actual := RotEncipher(strings.ToUpper(text), 18, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestRotDecipher(t *testing.T) {
	var text = "stuvwxyz0123456789abcdefghijklmnopqrSTUVWXYZ0123456789ABCDEFGHijklmnopqr"
	var expected = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // Loses capitalization on numbers
	actual := RotDecipher(strings.ToUpper(text), 18, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestRotEncipherCaps(t *testing.T) {
	var text = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var expected = "stuvwxyz0123456789abcdefghijklmnopqrSTUVWXYZ0123456789ABCDEFGHijklmnopqr"
	actual := RotEncipherCaps(text, 18, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestRotDecipherCaps(t *testing.T) {
	var text = "stuvwxyz0123456789abcdefghijklmnopqrSTUVWXYZ0123456789ABCDEFGHijklmnopqr"
	var expected = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHijklmnopqrSTUVWXYZ0123456789" // Loses capitalization on numbers
	actual := RotDecipherCaps(text, 18, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestRotEncipherRange(t *testing.T) {
	var text = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var expected = "abcdefghijklmnopqrstuvwxyz0123456789NOPQRSTUVWXYZABCDEFGHIJKLM0123456789"
	actual := RotEncipherRange(text, 13, 'A', 'Z') // Only changes A-Z, not lowercase
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestRotDecipherRange(t *testing.T) {
	var text = "abcdefghijklmnopqrstuvwxyz0123456789NOPQRSTUVWXYZABCDEFGHIJKLM0123456789"
	var expected = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	actual := RotDecipherRange(text, 13, 'A', 'Z') // Only changes A-Z, not lowercase
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

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
	var expected = "nopqrstuvwxyzabcdefghijklm5678901234NOPQRSTUVWXYZABCDEFGHIJKLM5678901234"
	actual := Rot18Encipher(text)
	if expected != actual {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestRot18Decipher(t *testing.T) {
	var text = "nopqrstuvwxyzabcdefghijklm5678901234NOPQRSTUVWXYZABCDEFGHIJKLM5678901234"
	var expected = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
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
