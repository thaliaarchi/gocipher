package ngrams

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

//go:generate go run generate_monograms.go

var (
	EnglishUnigrams *NgramSet
	EnglishBigrams  *NgramSet
)

type NgramSet struct {
	Language   string
	N          int
	Ngrams     []*Ngram
	NgramMap   map[string]*Ngram
	TotalCount int
}

type Ngram struct {
	Chars string
	Count int
	Freq  float64
}

func (set *NgramSet) GetNgram(ngram string) (*Ngram, bool) {
	n, ok := set.NgramMap[ngram]
	return n, ok
}

// GetEntropy gets the entropy of a string according to English n-gram frequencies.
func (set *NgramSet) GetEntropy(text string) float64 {
	n := set.N
	runes := []rune(text)
	ngrams := make([]string, len(runes)-n+1)
	for i := 0; i <= len(runes)-n; i++ {
		ngrams[i] = string(runes[i : i+n])
	}
	var sum float64
	var ignored int
	for _, ngramName := range ngrams {
		ngram, hasFreq := set.GetNgram(ngramName)
		if hasFreq && ngram.Freq != 0 {
			sum += math.Log(ngram.Freq)
		} else {
			ignored++
		}
	}
	return -sum / math.Log(2) / float64(len(ngrams)-ignored)
}

// LoadNgramsFile reads and parses all n-grams from a file.
func LoadNgramsFile(language string, n int) (*NgramSet, error) {
	filename := fmt.Sprintf("testdata/%s_%d.txt", language, n)
	var r io.Reader
	if stat, err := os.Stat(filename + ".zip"); err == nil {
		filename += ".zip"
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		zr, err := zip.NewReader(file, stat.Size())
		if len(zr.File) != 1 {
			return nil, fmt.Errorf("zip should contain only 1 file: %s", filename)
		}
		fr, err := zr.File[0].Open()
		if err != nil {
			return nil, err
		}
		r = fr
	} else {
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		r = file
	}
	return LoadNgrams(language, n, r)
}

// LoadNgrams reads and parses all n-grams from a reader.
func LoadNgrams(language string, n int, r io.Reader) (*NgramSet, error) {
	var ngrams = []*Ngram{}
	var ngramMap = make(map[string]*Ngram)
	totalCount := 0
	scanner := bufio.NewScanner(r)
	for i := 0; scanner.Scan(); i++ {
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		text := scanner.Text()
		split := strings.IndexByte(text, ' ')
		count, err := strconv.ParseInt(text[split+1:], 10, 0)
		if err != nil {
			return nil, err
		}
		entry := &Ngram{Chars: text[:split], Count: int(count)}
		ngrams = append(ngrams, entry)
		ngramMap[entry.Chars] = entry
		totalCount += int(count)
	}
	for i := range ngrams {
		ngrams[i].Freq = float64(ngrams[i].Count) / float64(totalCount)
	}
	return &NgramSet{language, n, ngrams, ngramMap, totalCount}, nil
}

func init() {
	var err error
	EnglishUnigrams, err = LoadNgramsFile("en", 1)
	if err != nil {
		panic(err)
	}
	EnglishBigrams, err = LoadNgramsFile("en", 2)
	if err != nil {
		panic(err)
	}
}
