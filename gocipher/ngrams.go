package gocipher

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var englishUnigrams = LoadNgrams("english", 1)
var englishBigrams = LoadNgrams("english", 2)

type NgramSet struct {
	language   string
	n          int
	ngrams     []*Ngram
	ngramMap   map[string]*Ngram
	totalCount int
}

type Ngram struct {
	chars string
	count int
	freq  float64
}

func (set *NgramSet) GetNgram(ngram string) (*Ngram, bool) {
	n, ok := set.ngramMap[ngram]
	//log.Println(ngram, n, ok)
	return n, ok
}

// GetEntropy gets the entropy of a string according to English n-gram frequencies.
func (set *NgramSet) GetEntropy(text string) float64 {
	n := set.n
	runes := []rune(text)
	ngrams := make([]string, len(runes)-n+1)
	for i := 0; i <= len(runes)-n; i++ {
		ngrams[i] = string(runes[i : i+n])
	}
	var sum float64
	var ignored int
	for _, ngramName := range ngrams {
		ngram, hasFreq := set.GetNgram(ngramName)
		if hasFreq && ngram.freq != 0 {
			sum += math.Log(ngram.freq)
		} else {
			ignored++
		}
	}
	return -sum / math.Log(2) / float64(len(ngrams)-ignored)
}

// LoadNgrams reads and parses all n-grams from a file.
// See: https://stackoverflow.com/a/23667119/3238709.
func LoadNgrams(language string, n int) *NgramSet {
	fileName := "ngrams/" + language + "_" + strconv.Itoa(n) + "-grams.txt"
	if fileName == "" {
		fmt.Println("Filename must not be empty")
		os.Exit(2)
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("error opening file %q: %v", fileName, err))
	}
	var ngrams = []*Ngram{}
	var ngramMap = make(map[string]*Ngram)
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
		entry := &Ngram{chars: text[:split], count: int(count)}
		ngrams = append(ngrams, entry)
		ngramMap[entry.chars] = entry
		totalCount += int(count)
	}
	for i := range ngrams {
		ngrams[i].freq = float64(ngrams[i].count) / float64(totalCount)
	}
	return &NgramSet{language, n, ngrams, ngramMap, totalCount}
}
