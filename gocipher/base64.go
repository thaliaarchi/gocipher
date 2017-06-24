package gocipher

import "encoding/base64"

/*
 * Base64
 */

// Base64Encode encodes a string using Base64 according to encoding.
func Base64Encode(text string, encoding *base64.Encoding) string {
	return encoding.EncodeToString([]byte(text))
}

// Base64Decode dencodes a string using Base64 according to encoding.
func Base64Decode(text string, encoding *base64.Encoding) (string, error) {
	res, err := encoding.DecodeString(text)
	if err != nil {
		return "", nil
	}
	return string(res), err
}

// Base64EncodeStd encodes a string using the standard Base64 encoding defined in RFC 4648.
// Uses `+/` and padding is with `=`.
func Base64EncodeStd(text string) string {
	return Base64Encode(text, base64.StdEncoding)
}

// Base64DecodeStd decodes a string using the standard Base64 encoding defined in RFC 4648.
// Uses `+/` and padding is with `=`.
func Base64DecodeStd(text string) (string, error) {
	return Base64Decode(text, base64.StdEncoding)
}

// Base64EncodeURL encodes a string using the alternate Base64 encoding defined in RFC 4648.
// Typically used in URLs and file names.
// Uses `-_` and padding is with `=`.
func Base64EncodeURL(text string) string {
	return Base64Encode(text, base64.URLEncoding)
}

// Base64DecodeURL decodes a string using the alternate Base64 encoding defined in RFC 4648.
// Typically used in URLs and file names.
// Uses `-_` and padding is with `=`.
func Base64DecodeURL(text string) (string, error) {
	return Base64Decode(text, base64.URLEncoding)
}

// Base64EncodeRawStd encodes a string using the standard raw, unpadded Base64 encoding defined in RFC 4648 section 3.2.
// This is the same as Base64EncodeStd but omits padding characters.
func Base64EncodeRawStd(text string) string {
	return Base64Encode(text, base64.RawStdEncoding)
}

// Base64DecodeRawStd decodes a string using the standard raw, unpadded Base64 encoding defined in RFC 4648 section 3.2.
// This is the same as Base64EncodeStd but omits padding characters.
func Base64DecodeRawStd(text string) (string, error) {
	return Base64Decode(text, base64.RawStdEncoding)
}

// Base64EncodeRawURL encodes a string using the unpadded alternate Base64 encoding defined in RFC 4648.
// Typically used in URLs and file names.
// This is the same as Base64EncodeURL but omits padding characters.
func Base64EncodeRawURL(text string) string {
	return Base64Encode(text, base64.RawURLEncoding)
}

// Base64DecodeRawURL decodes a string using the unpadded alternate Base64 encoding defined in RFC 4648.
// Typically used in URLs and file names.
// This is the same as Base64DecodeURL but omits padding characters.
func Base64DecodeRawURL(text string) (string, error) {
	return Base64Decode(text, base64.RawURLEncoding)
}
