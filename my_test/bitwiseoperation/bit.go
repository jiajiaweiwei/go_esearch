package bitwiseoperation

// 位运算常见应用场景
// 判断某位的值
// 修改特定位的值
/*
练习题
	1.判断一个uint第四位是否为0
	2.把一个uint的第四位置为零
	3.求一个uint的二进制里包含几个1
*/

const (
	FirstBitCode = 1
	ForthBitCode = 1 << 3
)

// ForthBitIfZero 判断一个uint第四位是否为0
func ForthBitIfZero(num uint) bool {
	return num&ForthBitCode == 0
}

// ForthBitSetZero 把一个uint的第四位置为零
func ForthBitSetZero(num uint) uint {
	// &^ 是 Go 特有的按位清除（bit clear / AND NOT） 运算符
	return num &^ ForthBitCode
}

// CountOneBitNum 求一个uint的二进制里包含几个1
func CountOneBitNum(num uint) uint {
	var count uint
	for num != 0 {
		if num&1 != 0 {
			count++
		}
		num = num >> 1
	}
	return count
}

// 判断第i位是否为1，i从1开始
func IsBit1(n uint64, i int) bool {
	if i > 64 {
		panic("i can not more than 64")
	}
	constNum := uint64(1 << (i - 1))
	return n&constNum == constNum
}

// 把第i位置为1 ，i 从1开始
func SetBit1(n uint64, i int) uint64 {
	if i > 64 {
		panic("i can not more than 64")
	}
	constNum := uint64(1 << (i - 1))
	return n | constNum
}

// Candidate (翻译：候选人) 位运算实现搜索过滤
type Candidate struct {
	Id     int
	Gender string
	Vip    bool
	Active int    // 几天内活跃
	Bits   uint64 // 用于位运算筛选
}

const (
	MALE        = 1 << iota // 是否男性
	VIP                     // 是否VIP
	WEEK_ACTIVE             // 是否周活
)

func (c *Candidate) SetMale() {
	c.Gender = "male"
	c.Bits |= MALE
}
