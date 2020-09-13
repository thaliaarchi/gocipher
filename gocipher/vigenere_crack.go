package gocipher

import "sort"

/*
 * Cracks the Vigenère cipher by testing every possible key of given length and ordering by ngrams
 */

// splitChunks splits a string into chunks.
func splitChunks(text string, n int) []string {
	runes := []rune(text)
	len := len(runes)
	chunks := make([]string, n)
	for i := 0; i < n; i++ {
		chunk := ""
		for j := i; j < len; j += n {
			chunk += string(runes[j])
		}
		chunks[i] = chunk
	}
	return chunks
}

// joinChunks joins chunks that were split by splitChunks.
// Assumes first chunk has largest length.
func joinChunks(chunks []string) string {
	size := len(chunks[0])
	res := make([]rune, len(chunks)*size)
	pos := 0
	for i := 0; i < size; i++ {
		for _, chunk := range chunks {
			if i < len(chunk) {
				res[pos] = []rune(chunk)[i]
				pos++
			}
		}
	}
	return string(res[:pos])
}

// cartesianProduct returns the cartestian product of multiple slices.
// See: https://stackoverflow.com/a/15310051/3238709
func cartesianProduct(slices [][]string) [][]string {
	product := [][]string{}
	max := len(slices) - 1
	var cartesian func([]string, int)
	cartesian = func(slice []string, i int) {
		len := len(slices[i])
		for j := 0; j < len; j++ {
			s := append(slice, slices[i][j])
			if i == max {
				product = append(product, s)
			} else {
				cartesian(s, i+1)
			}
		}
	}
	cartesian([]string{}, 0)
	return product
}

// VigenerePossibilities returns all possible plaintexts for a Vigenère cipher for a given key length.
func VigenerePossibilities(text string, keyLength int) []string {
	chunks := splitChunks(text, keyLength)
	chunkShifts := make([][]string, len(chunks))
	for i, chunk := range chunks {
		shifts := make([]string, 26)
		for j := 0; j < 26; j++ {
			shifts[j] = caesarEncipher(chunk, j)
		}
		chunkShifts[i] = shifts
	}
	possible := cartesianProduct(chunkShifts)
	res := make([]string, len(possible))
	for i, chunks := range possible {
		res[i] = joinChunks(chunks)
	}
	return res
}

type possibility struct {
	text    string
	entropy float64
}

// SortPossibilities sorts strings by bigram entropy
func SortPossibilities(possible []string) []possibility {
	poss := make([]possibility, len(possible))
	for i, p := range possible {
		poss[i] = possibility{p, englishBigrams.GetEntropy(p)}
	}
	sort.Slice(poss, func(i, j int) bool {
		return poss[i].entropy <= poss[j].entropy
	})
	return poss
}

func caesarEncipher(text string, key int) string {
	if key == 0 {
		return text
	}
	return mapAlpha(text, func(i, char int) int {
		return char + key
	})
}
