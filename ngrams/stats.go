package ngrams

import "math"

/*
 * Statistics routines for cryptanalysis
 */

// IC - Calculates index of coincidence for ciphertext
func IC(text string) float64 {
	counts := NgramCount(text, 1)
	icval := 0
	for _, count := range counts {
		icval += count * (count - 1)
	}
	return float64(icval) / float64(len(text)*(len(text)-1))
}

// NgramCount - Returns a map containing each ngram and how many times it occurred
// unigrams (letters), bigrams (letter pairs), trigrams, quadgrams, quintgrams, etc.
func NgramCount(text string, n int) map[string]int {
	counts := make(map[string]int)
	for i := 0; i < len(text)-n+1; i++ {
		ngram := text[i : i+n]
		counts[ngram]++
	}
	return counts
}

// NgramFreq - Returns the n-gram frequencies of all n-grams encountered in text.
// Standard probabilities.
// Only n-grams occurring in 'text' will have probabilities.
// For the probability of not-occurring n-grams, use freq["floor"].
// This is set to floor/len(text)
func NgramFreq(text string, n int, floor float64) map[string]float64 { // floor=0.01
	counts := NgramCount(text, n)
	freqs := make(map[string]float64)
	len := float64(len(text) - n + 1)
	for ngram, count := range counts {
		freqs[ngram] = float64(count) / len
	}
	freqs["floor"] = floor / len
	return freqs
}

// NgramFreqLog - Log probabilities
func NgramFreqLog(text string, n int, floor float64) map[string]float64 {
	freqs := NgramFreq(text, n, floor)
	for ngram, freq := range freqs {
		freqs[ngram] = math.Log10(freq)
	}
	return freqs
}
