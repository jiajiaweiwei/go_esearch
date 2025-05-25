package util_types

type BitMap struct {
	Table uint64
}

// NewBitMap 注意这里 uint64上某位表示的第多少个位置的索引存在，减去min值用来降低内存消耗，避免荣誉
func NewBitMap(min int, arr []int) *BitMap {
	bitMap := new(BitMap)
	for _, value := range arr {
		index := value - min
		bitMap.Table = SetBit1(bitMap.Table, index)
	}
	return bitMap
}

// SetBit1  将uint64位的某位置1
func SetBit1(num uint64, index int) uint64 {
	if index < 1 || index > 64 {
		panic("index can not smaller than 0 or bigger than 64.")
	}
	// 位移运算的优先级高于或运算
	return num | 1<<(index-1)
}

// IntersectionOfBitMap 位图求交集，注意返回数据要加会偏移量
func IntersectionOfBitMap(a, b *BitMap, min int) []int {
	// 返回值预分配容量
	intersection := make([]int, 0, 64)
	// 与运算求交集
	intersectionBit := a.Table & b.Table
	// 根据交集位图求返回切片
	for i := 1; i < 64; i++ {
		if IfBit1(intersectionBit, i) {
			intersection = append(intersection, i+min)
		}
	}
	return intersection
}

// IfBit1 判断某位是否为1
func IfBit1(num uint64, index int) bool {
	if index < 1 || index > 64 {
		panic("index can not smaller than 0 or bigger than 64.")
	}
	return (num & (1 << (index - 1))) != 0
}
