package haar

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"reflect"
	"sync"
	"time"
	"unsafe"
)

const SIZEOF_INT32 = 4
const SIZEOF_INT64 = 8
const debug = false

func ProcessFile(in, out *os.File) (time.Duration, error) {
	temp := make([]byte, SIZEOF_INT64)
	_, err := in.Read(temp)
	if err != nil {
		return 0, err
	}
	dimension := getDimension(temp)
	inBuf := make([]byte, dimension*dimension*SIZEOF_INT32+SIZEOF_INT64)
	if _, err := in.ReadAt(inBuf, 0); err != nil {
		return 0, err
	}

	data := castSlice(inBuf[SIZEOF_INT64:])
	image := Image{data, dimension}
	start := time.Now()
	image.Transform()
	elapsed := time.Since(start)

	if _, err := out.Write(inBuf); err != nil {
		return elapsed, fmt.Errorf("Failed to write to '%s': %v", out.Name(), err)
	}

	return elapsed, nil
}

func castSlice(data []byte) []int32 {
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	// The length and capacity of the slice are different.
	header.Len /= SIZEOF_INT32
	header.Cap /= SIZEOF_INT32

	// Convert slice header to an []int32
	return *(*[]int32)(unsafe.Pointer(&header))
}

func getDimension(data []byte) uint64 {
	if isLittleEndian() {
		return binary.LittleEndian.Uint64(data)
	} else {
		return binary.BigEndian.Uint64(data)
	}
}

func isLittleEndian() bool {
	var i int32 = 0x01020304
	u := unsafe.Pointer(&i)
	pb := (*byte)(u)
	b := *pb
	return (b == 0x04)
}

var sqrt_2 = math.Sqrt(2)

type Image struct {
	Pixels    []int32
	Dimension uint64
}

func (i Image) print() {
	for y := uint64(0); y < i.Dimension; y++ {
		for x := uint64(0); x < i.Dimension; x++ {
			fmt.Printf("%10d ", i.Pixels[y*i.Dimension+x])
		}
		fmt.Print("\n")
	}
}

func (i Image) Transform() {
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&i.Pixels))
	pixels := header.Data
	dimension := uintptr(i.Dimension) * SIZEOF_INT32

	var waitGroup sync.WaitGroup
	for s := i.Dimension; s > 1; s /= 2 {
		mid := uintptr(s/2) * SIZEOF_INT32
		upper := mid*uintptr(i.Dimension) + pixels

		// row-transformation

		waitGroup.Add(int(s / 2))
		for row := pixels; row < upper; row += dimension {
			go func(row uintptr) {
				upperInner := mid + row
				for p := row; p < upperInner; p += SIZEOF_INT32 {
					pixel1 := (*int32)(unsafe.Pointer(p))
					pixel2 := (*int32)(unsafe.Pointer(p + mid))

					a := float64(*pixel1+*pixel2) / sqrt_2
					d := float64(*pixel1-*pixel2) / sqrt_2
					*pixel1 = int32(a)
					*pixel2 = int32(d)
				}
				waitGroup.Done()
			}(row)
		}
		waitGroup.Wait()
		if debug {
			fmt.Printf("after row-transformation: %d\n", s/2)
			i.print()
		}

		// column-transformation

		waitGroup.Add(int(s / 2))
		midOffset2 := mid * uintptr(i.Dimension)
		for row := pixels; row < upper; row += dimension {
			go func(row uintptr) {
				upperInner := mid + row
				for p := row; p < upperInner; p += SIZEOF_INT32 {
					pixel1 := (*int32)(unsafe.Pointer(p))
					pixel2 := (*int32)(unsafe.Pointer(p + midOffset2))

					a := float64(*pixel1+*pixel2) / sqrt_2
					d := float64(*pixel1-*pixel2) / sqrt_2
					*pixel1 = int32(a)
					*pixel2 = int32(d)
				}
				waitGroup.Done()
			}(row)
		}
		waitGroup.Wait()
		if debug {
			fmt.Printf("after column-transformation: %d\n", s/2)
			i.print()
		}
	}
}
