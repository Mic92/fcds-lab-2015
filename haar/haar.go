package haar

import (
	"fmt"
	"math"
	"os"
	"syscall"
)

func ProcessFile(in, out *os.File) error {
	info, err := in.Stat()
	if err != nil {
		return fmt.Errorf("Error stat input file '%s': %v", in.Name, err)
	}
	// check if out is a file
	if _, err := out.Stat(); err != nil {
		return fmt.Errorf("Error stat output file '%s': %v", out.Name, err)
	}

	// void *mmap(void *addr, size_t len, int prot, int flags, int fildes, off_t off);

	if err := out.Truncate(info.Size()); err != nil {
		return fmt.Errorf("Failed to resize output file '%s' to %d bytes: %v", out.Name, info.Size(), err)
	}

	// linux specific, but fast
	if _, err := syscall.Splice(int(in.Fd()), nil, int(out.Fd()), nil, int(info.Size()), 0); err != nil {
		return fmt.Errorf("Failed to copy input to output file: %v", err)
	}

	//buf, err := syscall.Mmap(int(out.Fd()), 0, int(info.Size()), syscall.PROT_WRITE|syscall.PROT_READ, syscall.MAP_SHARED)
	//if err != nil {
	//	return fmt.Errorf("Failed to mmap output file: %v", err)
	//}

	return nil
}

var sqrt_2 = math.Sqrt(2)

type Image struct {
	Pixels    []int
	Dimension int64
}

// will it be inlined?
func (i Image) at(x, y int64) int64 {
	return int64(i.Pixels[y*i.Dimension+x])
}

// will it be inlined?
func (i Image) to(x, y int64, value int) {
	i.Pixels[y*i.Dimension+x] = value
}

func (i Image) print() {
	for y := int64(0); y < i.Dimension; y++ {
		for x := int64(0); x < i.Dimension; x++ {
			fmt.Printf("%10d ", i.Pixels[y*i.Dimension+x])
		}
		fmt.Print("\n")
	}
}

const debug = true

func (i Image) Transform() {
	if debug {
		fmt.Printf("origin: %d\n", i.Dimension)
		i.print()
	}

	for s := i.Dimension; s > 1; s /= 2 {
		mid := s / 2
		// row-transformation
		for y := int64(0); y < mid; y++ {
			for x := int64(0); x < mid; x++ {
				pixel1 := int64(i.Pixels[y*i.Dimension+x])
				pixel2 := int64(i.Pixels[y*i.Dimension+x+mid])
				a := float64(pixel1+pixel2) / sqrt_2
				d := float64(pixel1-pixel2) / sqrt_2
				i.Pixels[y*i.Dimension+x] = int(a)
				i.Pixels[y*i.Dimension+x+mid] = int(d)
			}
		}
		if debug {
			fmt.Printf("after row-transformation: %d\n", mid)
			i.print()
		}
		// column-transformation
		for y := int64(0); y < mid; y++ {
			for x := int64(0); x < mid; x++ {
				pixel1 := int64(i.Pixels[y*i.Dimension+x])
				pixel2 := int64(i.Pixels[(y+mid)*i.Dimension+x])
				a := float64(pixel1+pixel2) / sqrt_2
				d := float64(pixel1-pixel2) / sqrt_2
				i.Pixels[y*i.Dimension+x] = int(a)
				i.Pixels[(y+mid)*i.Dimension+x] = int(d)
			}
		}
		if debug {
			fmt.Printf("after column-transformation: %d\n", mid)
			i.print()
		}
	}
}
