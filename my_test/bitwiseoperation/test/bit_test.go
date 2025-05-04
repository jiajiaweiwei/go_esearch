package test

import (
	"fmt"
	"my_test/bitwiseoperation"
	"testing"
)

func Test_bit(t *testing.T) {
	println(bitwiseoperation.ForthBitIfZero(6))
	println("============================")
	var num uint = 9
	println(num)
	num = bitwiseoperation.ForthBitSetZero(num)
	println(num)
	println("============================")
	num = 9
	println(num)
	println(bitwiseoperation.CountOneBitNum(num))
	fmt.Printf("%064b", num)

}
