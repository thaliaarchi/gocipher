package gocipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLEncode(t *testing.T) {
	text := "https://golang.org/pkg/net/url/#QueryEscape?q=你好"
	expected := "https%3A%2F%2Fgolang.org%2Fpkg%2Fnet%2Furl%2F%23QueryEscape%3Fq%3D%E4%BD%A0%E5%A5%BD"
	actual := NewURLEncode().Encode(text)
	assert.Equal(t, expected, actual)
}

func TestURLDecode(t *testing.T) {
	text := "https%3A%2F%2Fgolang.org%2Fpkg%2Fnet%2Furl%2F%23QueryEscape%3Fq%3D%E4%BD%A0%E5%A5%BD"
	expected := "https://golang.org/pkg/net/url/#QueryEscape?q=你好"
	actual, err := NewURLEncode().Decode(text)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	assert.Equal(t, expected, actual)
}
