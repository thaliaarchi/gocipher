package gocipher

import "testing"

type railfenceTest struct {
	expected string
	key      int
}

func TestRailfenceEncipher(t *testing.T) {
	input := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tests := []railfenceTest{
		{"aeimquyCGKOSWbdfhjlnprtvxzBDFHJLNPRTVXZcgkoswAEIMQUY", 3},
		{"akuEOYbjltvDFNPXZcimswCGMQWdhnrxBHLRVegoqyAIKSUfpzJT", 6},
		{"amyKWblnxzJLVXckowAIMUYdjpvBHNTZeiquCGOSfhrtDFPRgsEQ", 7},
		{"aoCQbnpBDPRcmqAEOSdlrzFNTeksyGMUfjtxHLVZgiuwIKWYhvJX", 8}}
	for _, test := range tests {
		key, err1 := NewRailfenceKey(test.key)
		if err1 != nil {
			t.Error("Unexpected error:", err1)
		}
		actual, err2 := RailfenceEncipher(input, key)
		if err2 != nil {
			t.Error("Unexpected error:", err2)
		}
		if test.expected != actual {
			t.Errorf("Expected %q, but got %q (key %d)", test.expected, actual, test.key)
		}
	}
}

func TestRailfenceDecipher(t *testing.T) {
	input := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tests := []railfenceTest{
		{"anNobpOqcrPsdtQuevRwfxSygzTAhBUCiDVEjFWGkHXIlJYKmLZM", 3},
		{"agrBLVMCshbitDNWOEujckvFPXQGwldmxHRYSIyneozJTZUKApfq", 6},
		{"afoxGOWPHypgbhqzIQXRJAricjsBKSYTLCtkdluDMUZVNEvmenwF", 7},
		{"aelszGOWPHAtmfbgnuBIQXRJCvohcipwDKSYTLExqjdkryFMUZVN", 8}}
	for _, test := range tests {
		key, err1 := NewRailfenceKey(test.key)
		if err1 != nil {
			t.Error("Unexpected error:", err1)
		}
		actual, err2 := RailfenceDecipher(input, key)
		if err2 != nil {
			t.Error("Unexpected error:", err2)
		}
		if test.expected != actual {
			t.Errorf("Expected %q, but got %q (key %d)", test.expected, actual, test.key)
		}
	}
}
