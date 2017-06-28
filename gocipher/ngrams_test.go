package gocipher

import "testing"

func TestGetUnigramEntropy(t *testing.T) {
	text := "HELLOWORLD"
	expected := 4.303018577099257
	actual := englishUnigrams.GetEntropy(text)
	if expected != actual {
		t.Errorf("Expected %v, but got %v (text: %q)", expected, actual, text)
	}
}

func TestGetBigramEntropy(t *testing.T) {
	text := "HELLOWORLD"
	expected := 7.86033946544012
	actual := englishBigrams.GetEntropy(text)
	if expected != actual {
		t.Errorf("Expected %v, but got %v (text: %q)", expected, actual, text)
	}
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
