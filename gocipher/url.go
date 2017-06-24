package gocipher

import "net/url"

// URLEncode escapes a url string
func URLEncode(text string) string {
	return url.QueryEscape(text)
}

// URLDecode unescapes an escaped url string
// It returns an error if any % is not followed by two hexadecimal digits.
func URLDecode(text string) (string, error) {
	return url.QueryUnescape(text)
}
