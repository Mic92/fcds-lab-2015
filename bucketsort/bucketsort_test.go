package bucketsort

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"runtime"
	_ "sort"
	"testing"
)

func ok(t testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		t.Fatalf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
	}
}
func TestSort(t *testing.T) {
	f, err := os.Open("input/medium.in")
	ok(t, err)
	r := bufio.NewReader(f)
	s := make([]byte, 1)
	for {
		line, isPrefix, err := r.ReadLine()
		if isPrefix {
			t.Fatalf("line too long")
		}
		if err == io.EOF {
			break
		}
		ok(t, err)
		s = append(s, []byte(line)...)
	}
	//t.Logf("s=%s\n", Sort(s, 6))
	for _, idx := range Sort(s, 1) {
		t.Logf("%s", idx, s[idx:idx+1])
	}
	t.FailNow()

	//data := []byte("3xUF5PDDE02eXIXFHGNfdCHjV4ArNtTdTkGbTkFzxLTZdFKPEAzPi8jLnMhp7WP5HOAyHTQEofyNNrFGuDLcBsLMG83CgMHwAZFn")
	//s := make([]string, len(data))
	//s2 := make([]string, len(data))
	//for _, v := range data {
	//	s2 = append(s2, string(v))
	//}
	//for _, idx := range Sort(data, 1) {
	//	s = append(s, string(data[idx:idx+1]))
	//	t.Logf("data[%d]:%s", idx, data[idx:idx+1])
	//}
	//t.Logf("returns=%v\n", Sort(data, 2))
	//sort.Strings(s)
	//sort.Strings(s2)
	//t.Logf("s=%s\n", s)
	//t.Logf("s2=%s\n", s2)
	//t.FailNow()
}
