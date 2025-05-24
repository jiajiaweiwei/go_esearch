package main

import (
	"fmt"
	"github.com/huandu/skiplist"
)

func main() {
	// 测试跳表的合并
	list := skiplist.New(skiplist.Int32)
	list.Set(24, 24)
	list.Set(13, 13)
	PrintSkipList(list)
	list.Set(24, 25)
	PrintSkipList(list)
	list.Set(24, 26)
	PrintSkipList(list)
	fmt.Println("----------------")

}

func PrintSkipList(list *skiplist.SkipList) {
	println("print begin...")
	node := list.Front()
	for node != nil {
		fmt.Println(node.Key(), ":", node.Value)
		node = node.Next()
	}
	println("print end...")
}
