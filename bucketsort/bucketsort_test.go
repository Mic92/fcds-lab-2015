package bucketsort

import (
	sorting "sort"
	"testing"
)

func TestSort(t *testing.T) {
	data := make([]byte, 0)
	for i := byte('~'); i >= byte('!'); i-- {
		data = append(data, i)
		for j := byte('a') + 6; j > byte('a'); j-- {
			data = append(data, j)
		}
	}
	sorted := make([]string, len(data))
	for _, idx := range Sort(data, 7) {
		sorted = append(sorted, string(data[idx:idx+7]))
	}
	if !sorting.IsSorted(sorting.StringSlice(sorted)) {
		t.Fatalf("expect data to be sorted, got: %v", sorted)
	}
	t.FailNow()
}
