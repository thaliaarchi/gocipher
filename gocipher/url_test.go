package gocipher

import "testing"

func TestURLEncode(t *testing.T) {
	text := "https://golang.org/pkg/net/url/#QueryEscape?q=你好"
	expected := "https%3A%2F%2Fgolang.org%2Fpkg%2Fnet%2Furl%2F%23QueryEscape%3Fq%3D%E4%BD%A0%E5%A5%BD"
	actual := URLEncode(text)
	if expected != actual {
		t.Errorf("Expected %q, but got %q (text: %q)", expected, actual, text)
	}
}

func TestURLDecode(t *testing.T) {
	text := "https%3A%2F%2Fgolang.org%2Fpkg%2Fnet%2Furl%2F%23QueryEscape%3Fq%3D%E4%BD%A0%E5%A5%BD"
	expected := "https://golang.org/pkg/net/url/#QueryEscape?q=你好"
	actual, err := URLDecode(text)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	if expected != actual {
		t.Errorf("Expected %q, but got %q (text: %q)", expected, actual, text)
	}
}
