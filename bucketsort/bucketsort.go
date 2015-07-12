package bucketsort

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

const (
	NUMBER_OF_BUCKETS = 94
	WORD_LENGTH       = 7
)

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
			qsort(b)
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

func SortFile(in *os.File, out *os.File) (time.Duration, error) {
	data, err := readInput(in)
	if err != nil {
		return 0, err
	}

	buffedOut := bufio.NewWriter(out)
	defer buffedOut.Flush()

	start := time.Now()
	res := Sort(data)
	elapsed := time.Since(start)

	for _, v := range res {
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
			return elapsed, err
		}
	}
	return elapsed, nil
}

func qsort(a []uint64) {
	var qsort_rec func(int, int)
	qsort_rec = func(lower, upper int) {
		for {
			switch upper - lower {
			case -1, 0:
				return
			case 1:
				if a[upper] < a[lower] {
					a[upper], a[lower] = a[lower], a[upper]
				}
				return
			}

			bx := (upper + lower) / 2
			b := a[bx]
			lp := lower
			up := upper
		outer:
			for {
				for lp < upper && !(b < a[lp]) {
					lp++
				}
				for {
					if lp > up {
						break outer
					}
					if a[up] < b {
						break
					}
					up--
				}
				a[lp], a[up] = a[up], a[lp]
				lp++
				up--
			}
			if bx < lp {
				if bx < lp-1 {
					a[bx], a[lp-1] = a[lp-1], b
				}
				up = lp - 2
			} else {
				if bx > lp {
					a[bx], a[lp] = a[lp], b
				}
				up = lp - 1
				lp++
			}
			if up-lower < upper-lp {
				qsort_rec(lower, up)
				lower = lp
			} else {
				qsort_rec(lp, upper)
				upper = up
			}
		}
	}
	qsort_rec(0, len(a)-1)
}
