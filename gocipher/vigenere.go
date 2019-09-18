package gocipher

import (
	"fmt"
	"math"
)

var englishFreq = [26]float64{
	0.080406052, 0.014846488, 0.033437737, 0.038169583, 0.124920625,
	0.024031234, 0.018693758, 0.050533014, 0.075692775, 0.001587737,
	0.005405135, 0.04068986, 0.025117606, 0.072336292, 0.076406929,
	0.02135891, 0.001204689, 0.062794207, 0.065127666, 0.092755648,
	0.027297018, 0.010532516, 0.016756642, 0.002348569, 0.016649801, 0.000899507}

func VigenereDecrypt(text string, key []int) string {
	plain := make([]rune, len(text))
	for i, c := range text {
		switch {
		case 'a' <= c && c <= 'z':
			plain[i] = (c-'a'+rune(key[i%len(key)]))%26 + 'a'
		case 'A' <= c && c <= 'Z':
			plain[i] = (c-'A'+rune(key[i%len(key)]))%26 + 'A'
		default:
			plain[i] = c
		}
	}
	return string(plain)
}

func VigenereKey(text string, keyLen int) ([]int, error) {
	freqs := make([][26]float64, keyLen)
	for i, c := range text {
		switch {
		case 'a' <= c && c <= 'z':
			freqs[i%keyLen][c-'a']++
		case 'A' <= c && c <= 'Z':
			freqs[i%keyLen][c-'A']++
		default:
			return nil, fmt.Errorf("illegal character at index %d: %c", i, c)
		}
	}

	for k := 0; k < keyLen; k++ {
		n := float64((len(text) + keyLen - k) / keyLen)
		for i := 0; i < 26; i++ {
			freqs[k][i] /= n
		}
	}

	key := make([]int, keyLen)
	for k := 0; k < keyLen; k++ {
		best := math.MaxFloat64
		for i := 0; i < 26; i++ {
			chi := chiSquared(&freqs[k], &englishFreq, i)
			if chi < best {
				best = chi
				key[k] = i
			}
		}
	}

	return key, nil
}

func chiSquared(observeFreq, expectFreq *[26]float64, shift int) float64 {
	var sum float64
	for i := 0; i < 26; i++ {
		f := observeFreq[(i+shift)%26] - expectFreq[i]
		sum += f * f / expectFreq[i]
	}
	return sum
}
