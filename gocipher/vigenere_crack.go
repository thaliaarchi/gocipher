package gocipher

import (
	"math"
	"strings"
)

/*
 * Cracks the Vigenère cipher by testing every possible key of given length and ordering by ngrams
 */

// splitChunks splits a string into chunks.
func splitChunks(text string, size int) []string {
	len := len(text)
	n := int(math.Ceil(float64(len) / float64(size)))
	chunks := make([]string, n)
	for i := 0; i < n; i++ {
		chunk := ""
		for j := i; j < len; j += size - 1 {
			chunk += string(text[j])
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
func cartesian(array [][]string) []string {
	res := []string{}
	max := len(array) - 1
	var f func([]string, int)
	f = func(arr []string, i int) {
		len := len(array[i])
		for j := 0; j < len; j++ {
			a := append(arr, array[i][j])
			if i == max {
				res = append(res, strings.Join(a, ""))
			} else {
				f(a, i+1)
			}
		}
	}
	f([]string{}, 0)
	return res
}
