package gocipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type railfenceTest struct {
	expected string
	key      int
}

func TestRailfenceEncipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tests := []railfenceTest{
		{"aeimquyCGKOSWbdfhjlnprtvxzBDFHJLNPRTVXZcgkoswAEIMQUY", 3},
		{"akuEOYbjltvDFNPXZcimswCGMQWdhnrxBHLRVegoqyAIKSUfpzJT", 6},
		{"amyKWblnxzJLVXckowAIMUYdjpvBHNTZeiquCGOSfhrtDFPRgsEQ", 7},
		{"aoCQbnpBDPRcmqAEOSdlrzFNTeksyGMUfjtxHLVZgiuwIKWYhvJX", 8}}
	for _, test := range tests {
		railfence, err := NewRailfence(test.key)
		assert.Nil(t, err)
		actual, err := railfence.Encipher(text)
		assert.Nil(t, err)
		assert.Equal(t, test.expected, actual)
	}
}

func TestRailfenceDecipher(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tests := []railfenceTest{
		{"anNobpOqcrPsdtQuevRwfxSygzTAhBUCiDVEjFWGkHXIlJYKmLZM", 3},
		{"agrBLVMCshbitDNWOEujckvFPXQGwldmxHRYSIyneozJTZUKApfq", 6},
		{"afoxGOWPHypgbhqzIQXRJAricjsBKSYTLCtkdluDMUZVNEvmenwF", 7},
		{"aelszGOWPHAtmfbgnuBIQXRJCvohcipwDKSYTLExqjdkryFMUZVN", 8}}
	for _, test := range tests {
		railfence, err := NewRailfence(test.key)
		assert.Nil(t, err)
		actual, err := railfence.Decipher(text)
		assert.Nil(t, err)
		assert.Equal(t, test.expected, actual)
	}
}
