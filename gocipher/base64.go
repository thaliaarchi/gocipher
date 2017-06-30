package gocipher

import "encoding/base64"

/*
 * Base64
 */

type Base64 struct {
	encoding *base64.Encoding
}

func NewBase64(encoding *base64.Encoding) *Base64 {
	return &Base64{encoding}
}

// NewBase64Std creates a Base64 struct using the standard Base64 encoding defined in RFC 4648.
// Uses `+/` and padding is with `=`.
func NewBase64Std() *Base64 {
	return NewBase64(base64.StdEncoding)
}

// NewBase64URL creates a Base64 struct using the alternate Base64 encoding defined in RFC 4648.
// Typically used in URLs and file names.
// Uses `-_` and padding is with `=`.
func NewBase64URL() *Base64 {
	return NewBase64(base64.URLEncoding)
}

// NewBase64RawStd creates a Base64 struct using the standard raw, unpadded Base64 encoding defined in RFC 4648 section 3.2.
// This is the same as Base64EncodeStd but omits padding characters.
func NewBase64RawStd() *Base64 {
	return NewBase64(base64.RawStdEncoding)
}

// NewBase64RawURL creates a Base64 struct using the unpadded alternate Base64 encoding defined in RFC 4648.
// Typically used in URLs and file names.
// This is the same as Base64EncodeURL but omits padding characters.
func NewBase64RawURL() *Base64 {
	return NewBase64(base64.RawURLEncoding)
}

// Encode encodes a string using Base64 according to encoding.
func (b *Base64) Encode(text string) string {
	return b.encoding.EncodeToString([]byte(text))
}

// Decode dencodes a string using Base64 according to encoding.
func (b *Base64) Decode(text string) (string, error) {
	res, err := b.encoding.DecodeString(text)
	if err != nil {
		return "", nil
	}
	return string(res), err
}
