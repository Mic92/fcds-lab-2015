package bucketsort

import (
	"math/rand"
	sorting "sort"
	"testing"
)

func TestSort(t *testing.T) {
	rand.Seed(1)
	data := make([]byte, 0)
	for i := byte('~'); i >= byte('!'); i-- {
		data = append(data, i)
		for j := 0; j < 6; j++ {
			data = append(data, byte(rand.Int()%256))
		}
	}
	sorted := make([]string, len(data)/7)
	for _, v := range Sort(data) {
		word := []byte{
			byte(v >> 48),
			byte(v >> 40),
			byte(v >> 32),
			byte(v >> 24),
			byte(v >> 16),
			byte(v >> 8),
			byte(v),
			'\n',
		}
		sorted = append(sorted, string(word))
	}
	if !sorting.IsSorted(sorting.StringSlice(sorted)) {
		t.Fatalf("expect data to be sorted")
	}
}
