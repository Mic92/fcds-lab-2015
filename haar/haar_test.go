package haar_test

import (
	"github.com/Mic92/fcds-lab-2015/haar"
	"testing"
)

var pixels_4x4_in = haar.Image{[]int{
	55731,
	27808,
	30516,
	51498,
	56973,
	9353,
	41497,
	24883,
	62000,
	40078,
	65001,
	48498,
	28652,
	7372,
	9317,
	41237,
},
	4,
}
var pixels_4x4_out = haar.Image{[]int{
	126616,
	13415,
	17829,
	-16751,
	28337,
	22330,
	10943,
	-10981,
	-717,
	11313,
	65001,
	48498,
	28974,
	11904,
	9317,
	41237,
},
	4,
}

func TestTransform(t *testing.T) {
	pixels_4x4_in.Transform()
	sucess := true
	for i := range pixels_4x4_in.Pixels {
		if pixels_4x4_in.Pixels[i] == pixels_4x4_out.Pixels[i] {
			t.Logf("[%d]=%d", i, pixels_4x4_in.Pixels[i])
		} else {
			sucess = false
			t.Logf("[%d] %d != %d", i, pixels_4x4_in.Pixels[i], pixels_4x4_out.Pixels[i])
		}
	}
	if !sucess {
		t.Fail()
	}
}
