package haar

import (
	"fmt"
	"os"
	"syscall"
)

func ProcessFile(in, out *os.File) error {
	info, err := in.Stat()
	if err != nil {
		return fmt.Errorf("Error stat input file: %v", err)
	}
	// linux specific, but fast
	syscall.Splice(int(in.Fd()), nil, int(out.Fd()), nil, int(info.Size()), 0)
	return nil
}

//	SQRT_2 = sqrt(2);
//	for (s = size; s > 1; s /= 2) {
//		mid = s / 2;
//		// row-transformation
//		for (y = 0; y < mid; y++) {
//			for (x = 0; x < mid; x++) {
//				a = pixel(x,y);
//				a = (a+pixel(mid+x,y))/SQRT_2;
//				d = pixel(x,y);
//				d = (d-pixel(mid+x,y))/SQRT_2;
//				pixel(x,y) = a;
//				pixel(mid+x,y) = d;
//			}
//		}
//
//#ifdef DEBUG
//		printf("after row-transformation: %lld\n", mid);
//		print(pixels, size);
//#endif
//		// column-transformation
//		for (y = 0; y < mid; y++) {
//			for (x = 0; x < mid; x++) {
//				a = pixel(x,y);
//				a = (a+pixel(x,mid+y))/SQRT_2;
//				d = pixel(x,y);
//				d = (d-pixel(x,mid+y))/SQRT_2;
//				pixel(x,y) = a;
//				pixel(x,mid+y) = d;
//			}
//		}
//
//#ifdef DEBUG
//		printf("after column-transformation: %lld\n", mid);
//		print(pixels, size);
//#endif
//	}
