package bucketsort

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

const (
	NUMBER_OF_BUCKETS = 94
	WORD_LENGTH       = 7
)

func sort(data []byte, bucket []int, length int) {
	for j, key := range bucket {
		i := j - 1
		bIdx := key * length
		b := data[bIdx:(bIdx + length)]

		for i >= 0 {
			aIdx := bucket[i] * length
			a := data[aIdx:(aIdx + length)]
			if bytes.Compare(a, b) <= 0 {
				break
			}
			bucket[i+1] = bucket[i]
			i--
		}
		bucket[i+1] = key
	}
}

func Sort(data []byte, wordLength int) []int {
	var buckets [NUMBER_OF_BUCKETS][]int
	size := len(data) / wordLength

	returns := make([]int, size)
	step := size / NUMBER_OF_BUCKETS
	for i := range buckets {
		offset := i * step
		buckets[i] = returns[offset:offset : offset+step]
	}

	for i := 0; i < size; i++ {
		key := data[i*wordLength] - 0x21
		buckets[key] = append(buckets[key], i)
	}

	for _, bucket := range buckets {
		sort(data, bucket, wordLength)
	}

	return returns
}

func readInput(in *os.File, wordLength int) ([]byte, error) {
	r := bufio.NewReader(in)
	var lineCount int
	n, err := fmt.Fscanln(r, &lineCount)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return []byte{}, nil
	}

	data := make([]byte, 0, lineCount*wordLength)
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
	data, err := readInput(in, WORD_LENGTH)

	if err != nil {
		return err
	}
	for _, idx := range Sort(data, WORD_LENGTH) {
		offset := idx * WORD_LENGTH
		_, err = out.Write(data[offset : offset+WORD_LENGTH])
		if err != nil {
			return err
		}
		_, err = out.Write([]byte{'\n'})
		if err != nil {
			return err
		}
	}
	return nil
}
