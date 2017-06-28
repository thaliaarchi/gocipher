package gocipher

import (
	"fmt"
	"testing"
)

type Language struct {
	name  string
	files []int
}

var languages = []Language{
	{"danish", []int{1, 2, 3, 4}},
	{"english", []int{1, 2, 3, 4, 5}},
	{"finnish", []int{1, 2, 3, 4}},
	{"french", []int{1, 2, 3, 4}},
	{"german", []int{1, 2, 3, 4}},
	{"icelandic", []int{1, 2, 3, 4}},
	{"polish", []int{1, 2, 3, 4}},
	{"russian", []int{1, 2, 3, 4}},
	{"spanish", []int{1, 2, 3, 4}},
	{"swedish", []int{1, 2, 3, 4}}}

func TestReadNgrams(t *testing.T) {
	for _, lang := range languages {
		for _, file := range lang.files {
			set := ReadNgramFile("ngrams/" + lang.name + "_" + string('0'+file) + "-grams.txt")
			fmt.Println(set.fileName)
		}
	}
}
