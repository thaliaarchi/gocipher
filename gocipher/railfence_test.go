package gocipher

import "testing"

type railfenceTest struct {
	key      int
	expected string
}

func TestRailfenceEncipher(t *testing.T) {
	input := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tests := []railfenceTest{
		{3, "aeimquyCGKOSWbdfhjlnprtvxzBDFHJLNPRTVXZcgkoswAEIMQUY"},
		{6, "akuEOYbjltvDFNPXZcimswCGMQWdhnrxBHLRVegoqyAIKSUfpzJT"},
		{7, "amyKWblnxzJLVXckowAIMUYdjpvBHNTZeiquCGOSfhrtDFPRgsEQ"},
		{8, "aoCQbnpBDPRcmqAEOSdlrzFNTeksyGMUfjtxHLVZgiuwIKWYhvJX"}}
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
		{3, "anNobpOqcrPsdtQuevRwfxSygzTAhBUCiDVEjFWGkHXIlJYKmLZM"},
		{6, "agrBLVMCshbitDNWOEujckvFPXQGwldmxHRYSIyneozJTZUKApfq"},
		{7, "afoxGOWPHypgbhqzIQXRJAricjsBKSYTLCtkdluDMUZVNEvmenwF"},
		{8, "aelszGOWPHAtmfbgnuBIQXRJCvohcipwDKSYTLExqjdkryFMUZVN"}}
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
