package gocipher

import (
	"reflect"
	"testing"
)

func TestSplitChunks(t *testing.T) {
	text := "ABCDEFGHIJKL"
	n := 3
	expected := []string{"ADGJ", "BEHK", "CFIL"}
	actual := splitChunks(text, n)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v (text: %q, n: %d)", expected, actual, text, n)
	}

	text = "ABCD"
	n = 2
	expected = []string{"AC", "BD"}
	actual = splitChunks(text, n)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v (text: %q, n: %d)", expected, actual, text, n)
	}
}

func TestJoinChunks(t *testing.T) {
	chunks := []string{"ADGJ", "BEHK", "CFIL"}
	expected := "ABCDEFGHIJKL"
	actual := joinChunks(chunks)
	if expected != actual {
		t.Errorf("Expected %q, but got %q (chunks: %v)", expected, actual, chunks)
	}

	chunks = []string{"AC", "BD"}
	expected = "ABCD"
	actual = joinChunks(chunks)
	if expected != actual {
		t.Errorf("Expected %q, but got %q (chunks: %v)", expected, actual, chunks)
	}
}

func TestCartesianProduct(t *testing.T) {
	input := [][]string{
		[]string{"ABC", "abc", "123"},
		[]string{"DEF", "def", "456"},
		[]string{"GHI", "ghi", "789"}}
	expected := [][]string{
		[]string{"ABC", "DEF", "GHI"}, []string{"ABC", "DEF", "ghi"}, []string{"ABC", "DEF", "789"},
		[]string{"ABC", "def", "GHI"}, []string{"ABC", "def", "ghi"}, []string{"ABC", "def", "789"},
		[]string{"ABC", "456", "GHI"}, []string{"ABC", "456", "ghi"}, []string{"ABC", "456", "789"},
		[]string{"abc", "DEF", "GHI"}, []string{"abc", "DEF", "ghi"}, []string{"abc", "DEF", "789"},
		[]string{"abc", "def", "GHI"}, []string{"abc", "def", "ghi"}, []string{"abc", "def", "789"},
		[]string{"abc", "456", "GHI"}, []string{"abc", "456", "ghi"}, []string{"abc", "456", "789"},
		[]string{"123", "DEF", "GHI"}, []string{"123", "DEF", "ghi"}, []string{"123", "DEF", "789"},
		[]string{"123", "def", "GHI"}, []string{"123", "def", "ghi"}, []string{"123", "def", "789"},
		[]string{"123", "456", "GHI"}, []string{"123", "456", "ghi"}, []string{"123", "456", "789"}}
	actual := cartesianProduct(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v (input: %v)", expected, actual, input)
	}
}

func TestVigenerePossibilities(t *testing.T) {
	text := "ABCD"
	keyLength := 2
	expected := []string{
		"ABCD", "ACCE", "ADCF", "AECG", "AFCH", "AGCI", "AHCJ", "AICK", "AJCL", "AKCM", "ALCN", "AMCO", "ANCP",
		"AOCQ", "APCR", "AQCS", "ARCT", "ASCU", "ATCV", "AUCW", "AVCX", "AWCY", "AXCZ", "AYCA", "AZCB", "AACC",
		"BBDD", "BCDE", "BDDF", "BEDG", "BFDH", "BGDI", "BHDJ", "BIDK", "BJDL", "BKDM", "BLDN", "BMDO", "BNDP",
		"BODQ", "BPDR", "BQDS", "BRDT", "BSDU", "BTDV", "BUDW", "BVDX", "BWDY", "BXDZ", "BYDA", "BZDB", "BADC"} // ... 13 times more
	expectedLen := 676 // 26^26
	actual := VigenerePossibilities(text, keyLength)
	hasError := false
	if expectedLen != len(actual) {
		t.Errorf("Expected length %d, but got %d", expectedLen, len(actual))
		hasError = true
	}
	for i, value := range expected {
		if value != actual[i] {
			t.Errorf("Expected %q at index %d, but got %q", value, i, actual[i])
			hasError = true
		}
	}
	if hasError {
		t.Errorf("Expected %v, but got %v (text: %q, keyLength: %d)", expected, actual, text, keyLength)
	}
}

func TestVigenereCrack(t *testing.T) {
	text := "WOWEHAXBXXAKTXMNXZKGZKWEHLWGKZAVEGZAXOLZAKPOLK"
	keyLength := 2
	expected := []possibility{
		possibility{"DIDYOUEVERHEARTHETRAGEDYOFDARTHPLAGUEISTHEWISE", 8.172407984971093},
		possibility{"EIEYPUFVFRIEBRUHFTSAHEEYPFEASTIPMAHUFITTIEXITE", 10.090462966192014},
		possibility{"AVALLHBIBEERXEQUBGONDRALLSANOGECINDHBVPGERTVPR", 10.142719564434199},
		possibility{"AIAYLUBVBREEXRQHBTOADEAYLFAAOTEPIADUBIPTEETIPE", 10.445133187741906},
		possibility{"OIOYZUPVPRSELREHPTCAREOYZFOACTSPWARUPIDTSEHIDE", 10.453868589666566},
		possibility{"SISYDUTVTRWEPRIHTTGAVESYDFSAGTWPAAVUTIHTWELIHE", 10.622255029567489},
		possibility{"DSDIOEEFEBHOABTREDRKGODIOPDKRDHZLKGEESSDHOWSSO", 10.622900757893847},
		possibility{"ZIZYKUAVARDEWRPHATNACEZYKFZANTDPHACUAIOTDESIOE", 10.667060528380993},
		possibility{"HIHYSUIVIRLEERXHITVAKEHYSFHAVTLPPAKUIIWTLEAIWE", 10.709924833516112},
		possibility{"ESEIPEFFFBIOBBURFDSKHOEIPPEKSDIZMKHEFSTDIOXSTO", 10.78191849947525}}
	expectedLen := 676 // 26^26
	poss := VigenerePossibilities(text, keyLength)
	actual := SortPossibilities(poss)
	if expectedLen != len(actual) {
		t.Errorf("Expected length %d, but got %d", expectedLen, len(actual))
	}
	if !reflect.DeepEqual(expected, actual[:len(expected)]) {
		t.Errorf("Expected %v, but got %v (text: %q, keyLength: %d)", expected, actual, text, keyLength)
	}
}
