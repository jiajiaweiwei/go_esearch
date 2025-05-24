package linklist

import (
	"fmt"
	"testing"
)

func Test_list(t *testing.T) {
	// 测试有序链表求交集
	fmt.Println(IntersectionOfOrderList([]int{2, 4, 6, 8}, []int{4, 8}))
	fmt.Println(IntersectionOfOrderList([]int{15, 20, 23, 30, 50}, []int{12, 15, 23, 30, 45, 50}))
}

func Test_lists(t *testing.T) {
	// 测试读个有序链表求交集
	fmt.Println(IntersectionOfOrderLists([]int{2, 4, 6, 8}, []int{4, 8}, []int{4, 15, 20, 23, 30, 50}, []int{4, 12, 15, 23, 30, 45, 50}))
	fmt.Println(IntersectionOfOrderLists([]int{15, 20, 23, 30, 50}, []int{12, 15, 23, 30, 45, 50}))
}
