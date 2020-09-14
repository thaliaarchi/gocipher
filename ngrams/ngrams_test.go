package ngrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUnigramEntropy(t *testing.T) {
	text := "HELLOWORLD"
	expected := 4.303018577099257
	actual := EnglishUnigrams.GetEntropy(text)
	assert.Equal(t, expected, actual)
}

func TestGetBigramEntropy(t *testing.T) {
	text := "HELLOWORLD"
	expected := 7.86033946544012
	actual := EnglishBigrams.GetEntropy(text)
	assert.Equal(t, expected, actual)
}

type language struct {
	name  string
	files []int
}

var languages = []language{
	{"da", []int{1, 2, 3, 4}},
	{"en", []int{1, 2, 3, 4, 5}},
	{"fi", []int{1, 2, 3, 4}},
	{"fr", []int{1, 2, 3, 4}},
	{"de", []int{1, 2, 3, 4}},
	{"is", []int{1, 2, 3, 4}},
	{"pl", []int{1, 2, 3, 4}},
	{"ru", []int{1, 2, 3, 4}},
	{"es", []int{1, 2, 3, 4}},
	{"sv", []int{1, 2, 3, 4}},
}

func TestLoadNgrams(t *testing.T) {
	t.SkipNow()
	for _, lang := range languages {
		for _, n := range lang.files {
			set, err := LoadNgramsFile(lang.name, n)
			if err != nil {
				t.Errorf("err for %s %d: %v", lang.name, n, err)
			}
			if set.language != lang.name || set.n != n {
				t.Errorf("want language %s %d, got language %s %d", lang.name, n, set.language, set.n)
			}
			if len(set.ngramMap) == 0 {
				t.Errorf("empty set: %s %d", lang.name, n)
			}
		}
	}
}
