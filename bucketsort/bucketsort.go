package bucketsort

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"sync"
)

const (
	NUMBER_OF_BUCKETS = 94
	WORD_LENGTH       = 7
)

type Uint64Slice []uint64

func (p Uint64Slice) Len() int           { return len(p) }
func (p Uint64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Sort(data []byte) []uint64 {
	var buckets [NUMBER_OF_BUCKETS][]uint64
	size := len(data) / WORD_LENGTH

	returns := make([]uint64, size)
	step := size / NUMBER_OF_BUCKETS
	for i := range buckets {
		offset := i * step
		buckets[i] = returns[offset:offset : offset+step]
	}

	for i := 0; i < size; i++ {
		a := data[i*WORD_LENGTH : i*WORD_LENGTH+WORD_LENGTH]
		key := a[0] - 0x21
		val := uint64(a[6]) |
			uint64(a[5])<<8 |
			uint64(a[4])<<16 |
			uint64(a[3])<<24 |
			uint64(a[2])<<32 |
			uint64(a[1])<<40 |
			uint64(a[0])<<48
		buckets[key] = append(buckets[key], val)
	}

	var wg sync.WaitGroup
	for i, bucket := range buckets {
		wg.Add(1)
		go func(i int, b []uint64) {
			sort.Sort(Uint64Slice(b))
			wg.Done()
		}(i, bucket)
	}
	wg.Wait()

	return returns
}

func readInput(in *os.File) ([]byte, error) {
	r := bufio.NewReader(in)
	var lineCount int
	n, err := fmt.Fscanln(r, &lineCount)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return []byte{}, nil
	}

	data := make([]byte, 0, lineCount*WORD_LENGTH)
	for {
		line, isPrefix, err := r.ReadLine()
		if isPrefix {
			return nil, fmt.Errorf("line too long")
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error while reading file %v: %v", in, err)
		}
		data = append(data, []byte(line)...)
	}
	return data, nil
}

func SortFile(in *os.File, out *os.File) error {
	data, err := readInput(in)
	if err != nil {
		return err
	}

	buffedOut := bufio.NewWriter(out)
	defer buffedOut.Flush()
	for _, v := range Sort(data) {
		outData := []byte{
			byte(v >> 48),
			byte(v >> 40),
			byte(v >> 32),
			byte(v >> 24),
			byte(v >> 16),
			byte(v >> 8),
			byte(v),
			'\n',
		}
		_, err = buffedOut.Write(outData)
		if err != nil {
			return err
		}
	}
	return nil
}
