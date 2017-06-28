package gocipher

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type NgramSet struct {
	fileName   string
	ngrams     []*ngram
	ngramMap   map[string]*ngram
	n          int
	totalCount int
}

type ngram struct {
	chars string
	count int
}

// ReadNgramFile reads and parses all n-grams from a file.
// See: https://stackoverflow.com/a/23667119/3238709.
func ReadNgramFile(fileName string) *NgramSet {
	if fileName == "" {
		fmt.Println("Filename must not be empty")
		os.Exit(2)
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("error opening file %q: %v", fileName, err))
	}
	var ngrams []*ngram
	var ngramMap = make(map[string]*ngram)
	totalCount := 0
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
		entry := &ngram{chars: text[:split], count: int(count)}
		ngrams = append(ngrams, entry)
		ngramMap[entry.chars] = entry
		totalCount += int(count)
	}
	return &NgramSet{fileName, ngrams, ngramMap, 0, totalCount}
}
