package gocipher

import (
	"math"
)

/*
 * Cracks the Vigenère cipher by testing every possible key of given length and ordering by ngrams
 */

// splitChunks splits a string into chunks.
func splitChunks(text string, size int) []string {
	runes := []rune(text)
	len := len(runes)
	n := int(math.Ceil(float64(len) / float64(size)))
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
	res := ""
	size := len(chunks[0])
	for i := 0; i < size; i++ {
		for _, chunk := range chunks {
			res += string(chunk[i])
		}
	}
	return res
}

// cartesian creates a slice of all combinations of 2D array items.
// See: https://stackoverflow.com/a/15310051/3238709
func cartesian(array [][]string) [][]string {
	res := [][]string{}
	max := len(array) - 1
	var f func([]string, int)
	f = func(arr []string, i int) {
		len := len(array[i])
		for j := 0; j < len; j++ {
			a := append(arr, array[i][j])
			if i == max {
				res = append(res, a)
			} else {
				f(a, i+1)
			}
		}
	}
	f([]string{}, 0)
	return res
}

// VigenerePossibilities returns all possible plaintexts for a Vigenère cipher for a given key length.
func VigenerePossibilities(text string, keyLength int) []string {
	chunks := splitChunks(text, keyLength)
	chunkShifts := make([][]string, len(chunks))
	for i, chunk := range chunks {
		shifts := make([]string, 26)
		for j := 0; j < 26; j++ {
			shifts[j] = CaesarEncipher(chunk, j)
		}
		chunkShifts[i] = shifts
	}
	possible := cartesian(chunkShifts)
	res := make([]string, len(possible))
	for i, chunks := range possible {
		res[i] = joinChunks(chunks)
	}
	return res
}
