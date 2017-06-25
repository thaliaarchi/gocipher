package gocipher

import "testing"

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
		key, err := NewRailfenceKey(test.key)
		if err != nil {
			t.Error("Unexpected error:", err)
		}
		actual, err := RailfenceEncipher(text, key)
		if err != nil {
			t.Error("Unexpected error:", err)
		}
		if test.expected != actual {
			t.Errorf("Expected %q, but got %q (text: %q, key %d)", test.expected, actual, text, test.key)
		}
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
		key, err := NewRailfenceKey(test.key)
		if err != nil {
			t.Error("Unexpected error:", err)
		}
		actual, err := RailfenceDecipher(text, key)
		if err != nil {
			t.Error("Unexpected error:", err)
		}
		if test.expected != actual {
			t.Errorf("Expected %q, but got %q (text: %q, key %d)", test.expected, actual, text, test.key)
		}
	}
}
