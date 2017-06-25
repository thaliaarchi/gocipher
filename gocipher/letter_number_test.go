package gocipher

import (
	"reflect"
	"testing"
)

func TestLetterNumberEncrypt(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}
	actual := LetterNumberEncrypt(text)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v (text: %q)", expected, actual, text)
	}
}

func TestLetterNumberDecrypt(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}
	expected := "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	actual := LetterNumberDecrypt(numbers)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %q, but got %q (text: %v)", expected, actual, numbers)
	}
}
