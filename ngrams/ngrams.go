package ngrams

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode/utf8"
)

//go:generate go run generate_tables.go

// Alphabets for various languages.
var (
	AlphaDA = "ABCDEFGHIJKLMNOPQRSTUVWXYZÆØÅ"
	AlphaDE = "ABCDEFGHIJKLMNOPQRSTUVWXYZÄÖÜß"
	AlphaEN = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AlphaES = "ABCDEFGHIJKLMNOPQRSTUVWXYZÑ"
	AlphaFI = "ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ"
	AlphaFR = "AÀÂÆBCÇDEÉÈÊËFGHIÎÏJKLMNOÔŒPQRSTUÙÛÜVWXYŸZ"
	AlphaIS = "AÁBDÐEÉFGHIÍJKLMNOÓPRSTUÚVXYÝÞÆÖ"
	AlphaPL = "AĄBCĆDEĘFGHIJKLŁMNŃOÓPRSŚTUWYZŹŻ"
	AlphaRU = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
	AlphaSV = "ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ"
)

type NgramSet struct {
	freqs    []float64
	alpha    string
	alphaLen uint
	lookup   map[rune]uint
	n        int
}

func NewNgramSet(alpha string, n int) *NgramSet {
	alphaLen := utf8.RuneCountInString(alpha)
	lookup := make(map[rune]uint, alphaLen)
	for i, ch := range alpha {
		lookup[ch] = uint(i)
	}
	ngramsLen := 1
	for i := 0; i < n; i++ {
		ngramsLen *= alphaLen
	}
	return &NgramSet{
		freqs:    make([]float64, ngramsLen),
		alpha:    alpha,
		alphaLen: uint(alphaLen),
		lookup:   lookup,
		n:        n,
	}
}

func (set *NgramSet) Add(ngram string, freq float64) error {
	hash, err := set.hash(ngram)
	if err != nil {
		return err
	}
	set.freqs[hash] = freq
	return nil
}

func (set *NgramSet) Get(ngram string) (float64, error) {
	hash, err := set.hash(ngram)
	if err != nil {
		return 0, err
	}
	return set.freqs[hash], nil
}

type binNgrams struct {
	alpha string
	n     int
	freqs []float64
}

func ReadNgramSet(r io.Reader) (*NgramSet, error) {
	var data binNgrams
	if err := binary.Read(r, binary.LittleEndian, &data); err != nil {
		return nil, err
	}
	alphaLen := utf8.RuneCountInString(data.alpha)
	lookup := make(map[rune]uint, alphaLen)
	for i, ch := range data.alpha {
		lookup[ch] = uint(i)
	}
	ngramsLen := 1
	for i := 0; i < data.n; i++ {
		ngramsLen *= alphaLen
	}
	if len(data.freqs) != ngramsLen {
		return nil, fmt.Errorf("read %d frequencies, want %d", len(data.freqs), ngramsLen)
	}
	return &NgramSet{
		freqs:    data.freqs,
		alpha:    data.alpha,
		alphaLen: uint(alphaLen),
		lookup:   lookup,
		n:        data.n,
	}, nil
}

func (set *NgramSet) WriteBinary(w io.Writer) error {
	data := binNgrams{set.alpha, set.n, set.freqs}
	return binary.Write(w, binary.LittleEndian, &data)
}

func (set *NgramSet) hash(ngram string) (uint, error) {
	var hash uint
	if utf8.RuneCountInString(ngram) != set.n {
		return 0, fmt.Errorf("ngram not length %d: %s", set.n, ngram)
	}
	for _, ch := range ngram {
		if _, ok := set.lookup[ch]; !ok {
			return 0, fmt.Errorf("rune %c not in alphabet: %s", ch, ngram)
		}
		hash = hash*set.alphaLen + set.lookup[ch]
	}
	return hash, nil
}

// func makeNgramSet(ngrams []ngramEntry, totalCount, n int, alpha map[rune]uint) (*NgramSet, error) {
// 	set := &NgramSet{Alpha: al}
// 	for _, ngram := range ngrams {
// 		if c := utf8.RuneCountInString(ngram.Ngram); c != n {
// 			return nil, fmt.Errorf("ngram not length %d: %s", n, ngram.Ngram)
// 		}

// 	}
// }

// GetEntropy gets the entropy of a string according to n-gram
// frequencies.
// func (set NgramSet) GetEntropy(text string) float64 {
// 	lo, hi := 0, 0
// 	for i := 0; i < set.N; i++ {
// 		if hi >= len(text) {
// 			return 0
// 		}
// 		_, size := utf8.DecodeRuneInString(text[hi:])
// 		hi += size
// 	}
// 	var sum float64
// 	count := 0
// 	for {
// 		freq, ok := set.Set[text[lo:hi]]
// 		if ok && freq != 0 {
// 			sum += math.Log(freq)
// 			count++
// 		}
// 		if hi >= len(text) {
// 			break
// 		}
// 		_, loSize := utf8.DecodeRuneInString(text[lo:])
// 		_, hiSize := utf8.DecodeRuneInString(text[hi:])
// 		lo += loSize
// 		hi += hiSize
// 	}
// 	return -sum / (math.Log(2) * float64(count))
// }

// LoadNgramsFile reads and parses all n-grams from a file.
func LoadNgramsFile(language, alpha string, n int) (*NgramSet, error) {
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
	return LoadNgrams(alpha, n, r)
}

// LoadNgrams reads and parses all n-grams from a reader.
func LoadNgrams(alpha string, n int, r io.Reader) (*NgramSet, error) {
	type ngramEntry struct {
		Ngram string
		Count int
	}
	var ngramCounts []ngramEntry
	total := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		split := bytes.IndexByte(line, ' ')
		count, err := strconv.Atoi(string(line[split+1:]))
		if err != nil {
			return nil, err
		}
		ngram := string(line[:split])
		if utf8.RuneCountInString(ngram) != n {
			return nil, fmt.Errorf("ngram is not %d runes: %s", n, ngram)
		}
		ngramCounts = append(ngramCounts, ngramEntry{ngram, count})
		total += count
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	ngrams := NewNgramSet(alpha, n)
	for _, entry := range ngramCounts {
		freq := float64(entry.Count) / float64(total)
		if err := ngrams.Add(entry.Ngram, freq); err != nil {
			return nil, err
		}
	}
	return ngrams, nil
}
