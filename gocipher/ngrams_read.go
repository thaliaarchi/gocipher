package gocipher

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Language struct {
	name  string
	files []string
}

var languages = []Language{
	{"danish", []string{"monograms", "bigrams", "trigrams", "quadgrams"}},
	{"english", []string{"monograms", "bigrams", "trigrams", "quadgrams", "quintgrams", "words"}},
	{"finnish", []string{"monograms", "bigrams", "trigrams", "quadgrams"}},
	{"french", []string{"monograms", "bigrams", "trigrams", "quadgrams"}},
	{"german", []string{"monograms", "bigrams", "trigrams", "quadgrams"}},
	{"icelandic", []string{"monograms", "bigrams", "trigrams", "quadgrams"}},
	{"polish", []string{"monograms", "bigrams", "trigrams", "quadgrams"}},
	{"russian", []string{"monograms", "bigrams", "trigrams", "quadgrams"}},
	{"spanish", []string{"monograms", "bigrams", "trigrams", "quadgrams"}},
	{"swedish", []string{"monograms", "bigrams", "trigrams", "quadgrams", "words"}}}

/*type NgramSet struct {
	fileName string
	ngrams   []Ngram
	ngramMap map[string]Ngram
	n        int
}*/

type Ngram struct {
	chars string
	count int
	freq  float64
}

// readNgramFile reads and parses all n-grams from a file.
// See: https://stackoverflow.com/a/23667119/3238709.
func readNgramFile(fileName string) ([]*Ngram, map[string]*Ngram, int) {
	if fileName == "" {
		fmt.Println("Filename must not be empty")
		os.Exit(2)
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("error opening file %q: %v", fileName, err))
	}
	var ngrams []*Ngram
	var ngramMap = make(map[string]*Ngram)
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
		entry := &Ngram{chars: text[:split], count: int(count)}
		ngrams = append(ngrams, entry)
		ngramMap[entry.chars] = entry
		countTotal += int(count)
	}
	for i := range ngrams {
		ngrams[i].freq = float64(ngrams[i].count) / float64(countTotal)
	}
	return ngrams, ngramMap, countTotal
}
