package gocipher

import (
	"reflect"
	"testing"
)

func TestSplitChunks(t *testing.T) {
	text := "ABCDEFGHIJKL"
	expected := []string{"ADGJ", "BEHK", "CFIL"}
	actual := splitChunks(text, 4)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v (text: %q)", expected, actual, text)
	}
}

func TestJoinChunks(t *testing.T) {
	chunks := []string{"ADGJ", "BEHK", "CFIL"}
	expected := "ABCDEFGHIJKL"
	actual := joinChunks(chunks)
	if expected != actual {
		t.Errorf("Expected %q, but got %q (chunks: %v)", expected, actual, chunks)
	}
}

func TestCartesian(t *testing.T) {
	input := [][]string{
		[]string{"ABC", "abc", "123"},
		[]string{"DEF", "def", "456"},
		[]string{"GHI", "ghi", "789"}}
	expected := []string{
		"ABCDEFGHI", "ABCDEFghi", "ABCDEF789",
		"ABCdefGHI", "ABCdefghi", "ABCdef789",
		"ABC456GHI", "ABC456ghi", "ABC456789",
		"abcDEFGHI", "abcDEFghi", "abcDEF789",
		"abcdefGHI", "abcdefghi", "abcdef789",
		"abc456GHI", "abc456ghi", "abc456789",
		"123DEFGHI", "123DEFghi", "123DEF789",
		"123defGHI", "123defghi", "123def789",
		"123456GHI", "123456ghi", "123456789"}
	actual := cartesian(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v (input: %v)", expected, actual, input)
	}
}
