package test

import (
	"fmt"
	"math/rand"
	"my_test/inverted"
	"strconv"
	"sync"
	"testing"
)

// 测试自定义倒排索引
func Test_D(t *testing.T) {
	docs := []*inverted.Doc{
		{
			Id:       1,
			Keywords: []string{"go语言", "进阶专题"},
		},
		{
			Id:       2,
			Keywords: []string{"go语言", "编程"},
		},
	}
	index := inverted.BuildInvertIndex(docs)
	for k, v := range index {
		fmt.Println(k, ":", v)
	}

}

// 基准测试 测试自定义线程安全的map
/*
go test -bench= -run=^$ -count=1 -benchmem -benchtime=3s
-bench=	运行所有 BenchmarkXxx 函数（等价于 -bench=.，但更安全简洁）
-run=^$	正则匹配，意思是“跳过所有 TestXxx 单元测试”，因为没有函数名能匹配空字符串
-count=1	执行一次基准测试（默认行为，显式写出来更可控）
-benchmem	输出每次操作的内存分配情况（B/op 和 allocs/op）
-benchtime=3s	每个基准测试至少运行 3 秒，提高结果稳定性和精度

*/
var myMap = inverted.NewConcurrentHashMap(8, 1000)
var syncMap = sync.Map{}

func BenchmarkMyMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		const P = 300
		var wg sync.WaitGroup
		// 并发读
		for i := 0; i < P; i++ {
			go func() {
				wg.Add(1)
				defer wg.Done()
				readMyMap()
			}()
		}
		// 并发写
		for i := 0; i < P; i++ {
			go func() {
				wg.Add(1)
				defer wg.Done()
				writeMyMap()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkSyncMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		const P = 300
		var wg sync.WaitGroup
		// 并发读
		for i := 0; i < P; i++ {
			go func() {
				wg.Add(1)
				defer wg.Done()
				readSyncMap()
			}()
		}
		// 并发写
		for i := 0; i < P; i++ {
			go func() {
				wg.Add(1)
				defer wg.Done()
				writeSyncMap()
			}()
		}
		wg.Wait()
	}
}

// 自定义线程安全的map
func readMyMap() {
	for i := 0; i < 10000; i++ {
		myMap.Load(strconv.Itoa(int(rand.Int63())))
	}
}

func writeMyMap() {
	for i := 0; i < 10000; i++ {
		myMap.Store(strconv.Itoa(int(rand.Int63())), "")
	}
}

// 原生sync.map
func readSyncMap() {
	for i := 0; i < 10000; i++ {
		syncMap.Load(strconv.Itoa(int(rand.Int63())))
	}
}

func writeSyncMap() {
	for i := 0; i < 10000; i++ {
		syncMap.Store(strconv.Itoa(int(rand.Int63())), "")
	}
}
