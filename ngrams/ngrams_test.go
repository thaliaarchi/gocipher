package ngrams

import (
	"fmt"
	"testing"
)

func TestNgrams(t *testing.T) {
	en1, err := LoadNgramsFile("en", AlphaEN, 2)
	if err != nil {
		t.Errorf("en1: %v", err)
	}
	fmt.Println(len(en1.freqs), en1)
	t.Fail()
}
