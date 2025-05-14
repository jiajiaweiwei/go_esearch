package main

import (
	"fmt"
	"sync"
)

func main() {
	var test sync.Locker
	test = &sync.Mutex{}
	mutex, ok := test.(*sync.RWMutex)
	if !ok {
		panic("test")
	}
	mutex.Lock()
}

func main2() {
	defer func() {
		fmt.Println("First defer starts")
		if r := recover(); r != nil {
			fmt.Println("Recovered in first defer:", r)
		} else {
			fmt.Println("First defer: no panic to recover")
		}
		fmt.Println("First defer ends")
	}()

	defer func() {
		fmt.Println("Second defer starts")
		// 这里覆盖了第一个panic
		panic("new panic from second defer")
		fmt.Println("Second defer ends") // 不会执行
	}()

	panic("original panic")
	fmt.Println("After original panic") // 不会执行
}

/*
执行结果
Second defer starts
First defer starts
Recovered in first defer: new panic from second defer
First defer ends
panic : original panic

*/
