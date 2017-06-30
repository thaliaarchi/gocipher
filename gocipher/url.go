package gocipher

import "net/url"

type URLEncode struct{}

func NewURLEncode() *URLEncode {
	return &URLEncode{}
}

// Encode escapes a url string.
func (u *URLEncode) Encode(text string) string {
	return url.QueryEscape(text)
}

// Decode unescapes an escaped url string.
// It returns an error if any % is not followed by two hexadecimal digits.
func (u *URLEncode) Decode(text string) (string, error) {
	return url.QueryUnescape(text)
}
