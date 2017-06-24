package gocipher

import "encoding/base64"

func Base64Encode(text string, encoding *base64.Encoding) string {
	return encoding.EncodeToString([]byte(text))
}

func Base64Decode(text string, encoding *base64.Encoding) (string, error) {
	res, err := encoding.DecodeString(text)
	if err != nil {
		return "", nil
	}
	return string(res), err
}

func Base64EncodeStd(text string) string {
	return Base64Encode(text, base64.StdEncoding)
}

func Base64DecodeStd(text string) (string, error) {
	return Base64Decode(text, base64.StdEncoding)
}

func Base64EncodeURL(text string) string {
	return Base64Encode(text, base64.URLEncoding)
}

func Base64DecodeURL(text string) (string, error) {
	return Base64Decode(text, base64.URLEncoding)
}

func Base64EncodeRawStd(text string) string {
	return Base64Encode(text, base64.RawStdEncoding)
}

func Base64DecodeRawStd(text string) (string, error) {
	return Base64Decode(text, base64.RawStdEncoding)
}

func Base64EncodeRawURL(text string) string {
	return Base64Encode(text, base64.RawURLEncoding)
}

func Base64DecodeRawURL(text string) (string, error) {
	return Base64Decode(text, base64.RawURLEncoding)
}
