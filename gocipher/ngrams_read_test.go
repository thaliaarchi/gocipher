package gocipher

import (
	"fmt"
	"testing"
)

func TestReadNgrams(t *testing.T) {
	type fileData struct {
		lang  string
		file  string
		count int
	}
	counts := []fileData{}
	for _, lang := range languages {
		for _, file := range lang.files {
			_, _, count := readNgramFile("ngrams/" + lang.name + "_" + file + ".txt")
			data := fileData{lang.name, file, count}
			counts = append(counts, data)
			fmt.Println(data)
		}
	}
	fmt.Println(counts)
}
