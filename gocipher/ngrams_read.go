package gocipher

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type NgramEntry struct {
	ngram string
	count int
	freq  float64
}

// readNgramFile reads and parses all n-grams from a file.
// See: https://stackoverflow.com/a/23667119/3238709.
func readNgramFile(fileName string) ([]*NgramEntry, map[string]*NgramEntry) {
	if fileName == "" {
		fmt.Println("Filename must not be empty")
		os.Exit(2)
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("error opening file %q: %v", fileName, err))
	}
	var ngrams []*NgramEntry
	var ngramMap = make(map[string]*NgramEntry)
	countTotal := 0
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading from file:", err)
			os.Exit(3)
		}
		text := scanner.Text()
		split := strings.IndexByte(text, ' ')
		count, err := strconv.ParseInt(text[split+1:], 10, 0)
		if err != nil {
			fmt.Println("N-gram count cannot be parsed to integer:", text, err)
		}
		entry := &NgramEntry{ngram: text[:split], count: int(count)}
		ngrams = append(ngrams, entry)
		ngramMap[entry.ngram] = entry
		countTotal += int(count)
		fmt.Println(text, entry)
	}
	for i, entry := range ngrams {
		ngrams[i].freq = float64(entry.count) / float64(countTotal)
		fmt.Println(entry)
	}
	fmt.Println(ngrams[0], ngramMap["A"])
	return ngrams, ngramMap
}
