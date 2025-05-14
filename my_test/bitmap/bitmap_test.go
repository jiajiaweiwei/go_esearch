package bitmap

import (
	"fmt"
	"testing"
)

func Test_BitMap(t *testing.T) {
	min := 14
	bm1 := NewBitMap(min, []int{15, 30, 20, 50, 23})
	bm2 := NewBitMap(min, []int{30, 15, 50, 20, 23, 45})
	fmt.Println(IntersectionOfBitMap(bm1, bm2, min))
}
