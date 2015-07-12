package haar

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"syscall"
	"time"
	"unsafe"
)

const SIZEOF_INT32 = 4
const SIZEOF_INT64 = 8

const debug = false

func ProcessFile(in, out *os.File) error {
	info, err := in.Stat()
	if err != nil {
		return fmt.Errorf("Error stat input file '%s': %v", in.Name, err)
	}

	if info.Size() < SIZEOF_INT64 {
		return fmt.Errorf("input file to small to contain size metadata")
	}

	inBuf, err := syscall.Mmap(int(in.Fd()), 0, int(info.Size()), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_PRIVATE)
	if err != nil {
		return fmt.Errorf("Failed to mmap input file: %v", err)
	}
	defer syscall.Munmap(inBuf)

	dimension := getDimension(inBuf)
	data := castSlice(inBuf[SIZEOF_INT64:])
	if debug {
		log.Printf("dimension %d", dimension)
	}
	image := Image{data, dimension}
	if debug {
		log.Printf("data size %d", len(data))
	}

	start := time.Now()
	image.Transform()
	elapsed := time.Since(start)
	log.Printf("real time took %dms", elapsed.Nanoseconds()/1e6)

	if _, err := syscall.Write(int(out.Fd()), inBuf); err != nil {
		return fmt.Errorf("Failed to copy input file '%s' to  '%s': %v", in.Name(), out.Name(), err)
	}

	return nil
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
	if debug {
		fmt.Printf("origin: %d\n", i.Dimension)
		i.print()
	}

	for s := i.Dimension; s > 1; s /= 2 {
		mid := s / 2
		// row-transformation
		upper := mid * i.Dimension
		for row := uint64(0); row < upper; row += i.Dimension {
			upperInner := mid + row
			for pos := uint64(row); pos < upperInner; pos++ {
				pixel1 := int64(i.Pixels[pos])
				pixel2 := int64(i.Pixels[pos+mid])
				a := float64(pixel1+pixel2) / sqrt_2
				d := float64(pixel1-pixel2) / sqrt_2
				i.Pixels[pos] = int32(a)
				i.Pixels[pos+mid] = int32(d)
			}
		}
		if debug {
			fmt.Printf("after row-transformation: %d\n", mid)
			i.print()
		}
		// column-transformation
		midOffset := mid * i.Dimension
		upper2 := mid * i.Dimension
		for row := uint64(0); row < upper2; row += i.Dimension {
			upperInner := mid + row
			for pos := uint64(row); pos < upperInner; pos++ {
				pixel1 := int64(i.Pixels[pos])
				pixel2 := int64(i.Pixels[pos+midOffset])
				a := float64(pixel1+pixel2) / sqrt_2
				d := float64(pixel1-pixel2) / sqrt_2
				i.Pixels[pos] = int32(a)
				i.Pixels[pos+midOffset] = int32(d)
			}
		}
		if debug {
			fmt.Printf("after column-transformation: %d\n", mid)
			i.print()
		}
	}
}
