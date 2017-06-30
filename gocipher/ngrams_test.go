package gocipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUnigramEntropy(t *testing.T) {
	text := "HELLOWORLD"
	expected := 4.303018577099257
	actual := englishUnigrams.GetEntropy(text)
	assert.Equal(t, expected, actual)
}

func TestGetBigramEntropy(t *testing.T) {
	text := "HELLOWORLD"
	expected := 7.86033946544012
	actual := englishBigrams.GetEntropy(text)
	assert.Equal(t, expected, actual)
}

type language struct {
	name  string
	files []int
}

var languages = []language{
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

func TestLoadNgrams(t *testing.T) {
	LoadNgrams("english", 1)
	/*for _, lang := range languages {
		for _, file := range lang.files {
			set := LoadNgrams(lang.name, file)
			fmt.Println(set.fileName)
		}
	}*/
}
