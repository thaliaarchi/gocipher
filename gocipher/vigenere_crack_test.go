package gocipher

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitChunks(t *testing.T) {
	text := "ABCDEFGHIJKLMN"
	n := 3
	expected := []string{"ADGJM", "BEHKN", "CFIL"}
	actual := splitChunks(text, n)
	assert.Equal(t, expected, actual)

	text = "ABCDE"
	n = 2
	expected = []string{"ACE", "BD"}
	actual = splitChunks(text, n)
	assert.Equal(t, expected, actual)
}

func TestJoinChunks(t *testing.T) {
	chunks := []string{"ADGJM", "BEHKN", "CFIL"}
	expected := "ABCDEFGHIJKLMN"
	actual := joinChunks(chunks)
	assert.Equal(t, expected, actual)

	chunks = []string{"ACE", "BD"}
	expected = "ABCDE"
	actual = joinChunks(chunks)
	assert.Equal(t, expected, actual)
}

func TestCartesianProduct(t *testing.T) {
	input := [][]string{
		{"ABC", "abc", "123"},
		{"DEF", "def", "456"},
		{"GHI", "ghi", "789"}}
	expected := [][]string{
		{"ABC", "DEF", "GHI"}, {"ABC", "DEF", "ghi"}, {"ABC", "DEF", "789"},
		{"ABC", "def", "GHI"}, {"ABC", "def", "ghi"}, {"ABC", "def", "789"},
		{"ABC", "456", "GHI"}, {"ABC", "456", "ghi"}, {"ABC", "456", "789"},
		{"abc", "DEF", "GHI"}, {"abc", "DEF", "ghi"}, {"abc", "DEF", "789"},
		{"abc", "def", "GHI"}, {"abc", "def", "ghi"}, {"abc", "def", "789"},
		{"abc", "456", "GHI"}, {"abc", "456", "ghi"}, {"abc", "456", "789"},
		{"123", "DEF", "GHI"}, {"123", "DEF", "ghi"}, {"123", "DEF", "789"},
		{"123", "def", "GHI"}, {"123", "def", "ghi"}, {"123", "def", "789"},
		{"123", "456", "GHI"}, {"123", "456", "ghi"}, {"123", "456", "789"}}
	actual := cartesianProduct(input)
	assert.Equal(t, expected, actual)
}

func TestVigenerePossibilities(t *testing.T) {
	text := "ABCD"
	keyLength := 2
	expected := []string{
		"ABCD", "ACCE", "ADCF", "AECG", "AFCH", "AGCI", "AHCJ", "AICK", "AJCL", "AKCM", "ALCN", "AMCO", "ANCP",
		"AOCQ", "APCR", "AQCS", "ARCT", "ASCU", "ATCV", "AUCW", "AVCX", "AWCY", "AXCZ", "AYCA", "AZCB", "AACC",
		"BBDD", "BCDE", "BDDF", "BEDG", "BFDH", "BGDI", "BHDJ", "BIDK", "BJDL", "BKDM", "BLDN", "BMDO", "BNDP",
		"BODQ", "BPDR", "BQDS", "BRDT", "BSDU", "BTDV", "BUDW", "BVDX", "BWDY", "BXDZ", "BYDA", "BZDB", "BADC"} // ... 13 times more
	expectedLen := 26 * 26
	actual := VigenerePossibilities(text, keyLength)
	assert.Equal(t, expectedLen, len(actual))
	assert.Equal(t, expected, actual[:len(expected)])
}

func TestVigenereCrack(t *testing.T) {
	text := "TKSYMWRJGHKBPTEIKCYRWXIELQUPSUTLLGYFIKYIAVFNRLQFKVVSMBMJOCZGILSEAPZRGCVVHTVQYKXJSHARVIPCOGHXGZCGLQNEEXLPDQVXWBLVKCTRSVXYWUORPNEJKVYBROGIQRABKZEGZAAJSMQRANLAGZCGLKVATZSUMEAFQICYSXLNPUSJLVORWIQVMULEMVXVJHHPIGIKGPLVWAITMTLJLQPVLJLBXPIIHGYZMBWVSXLFHZSGHKUTEKSDHCYVWWRTZCYGQICJMINRWBXYSVAJSXVFYTHZWPEMWUPZMTEIXGHGYZIJSNAUSCKYGPLUEAKRHKUTWMGLJKALLWPVKYOVPMXYWQAUIZHFWUUGEVIOHGYVIVGVVEYLTBSXJCWUIZGRFVLYPBLVVKMSIZIEUGZBGIRRLJPRJ"
	keyLength := 8
	poss := VigenerePossibilities(text, keyLength)
	sorted := SortPossibilities(poss)
	fmt.Println(sorted)
}
