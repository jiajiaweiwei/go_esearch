package inverted

// 有序链表求交集 这里直接用切片代表有序链表

// IntersectionOfOrderList 有序链表求交集
func IntersectionOfOrderList(a, b []int) []int {
	resultList := make([]int, 0, 10)
	var indexA, indexB int
	lengthA, lengthB := len(a), len(b)
	for indexA < lengthA && indexB < lengthB {
		if a[indexA] < b[indexB] {
			indexA++
		} else if a[indexA] > b[indexB] {
			indexB++
		} else {
			resultList = append(resultList, a[indexA])
			indexA++
			indexB++
		}
	}
	return resultList
}

// IntersectionOfOrderLists 求多个有序链表的交集
func IntersectionOfOrderLists(lists ...[]int) []int {
	// 处理边界情况：如果没有链表或有空链表，直接返回空结果
	if len(lists) == 0 {
		return []int{}
	}

	// 为每个链表初始化一个指针，初始位置都在链表头部
	pointers := make([]int, len(lists))

	// 存储交集结果
	var result []int

	// 所有指针都有效时继续循环
	for allPointersValid(pointers, lists) {
		// 获取所有当前指针指向的元素值
		currentValues := getCurrentValues(pointers, lists)

		// 计算当前的最大值
		maxVal := MaxFromLists(currentValues...)

		// 检查是否所有链表的当前元素都等于最大值
		if allEqual(currentValues, maxVal) {
			// 所有链表的当前元素都等于最大值，将该值加入交集
			result = append(result, maxVal)

			// 所有链表指针向后移动一位
			for i := range pointers {
				pointers[i]++
			}
		} else {
			// 否则，将所有小于最大值的链表指针向后移动一位
			for i, ptr := range pointers {
				if lists[i][ptr] < maxVal {
					pointers[i]++
				}
			}
		}
	}

	return result
}

// allPointersValid 检查所有指针是否都在有效范围内
func allPointersValid(pointers []int, lists [][]int) bool {
	for i, ptr := range pointers {
		if ptr >= len(lists[i]) {
			return false
		}
	}
	return true
}

// getCurrentValues 获取所有指针当前指向的元素值
func getCurrentValues(pointers []int, lists [][]int) []int {
	values := make([]int, len(pointers))
	for i, ptr := range pointers {
		values[i] = lists[i][ptr]
	}
	return values
}

// allEqual 检查切片中的所有元素是否都等于给定值
func allEqual(nums []int, target int) bool {
	for _, num := range nums {
		if num != target {
			return false
		}
	}
	return true
}

func MaxFromLists(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	maxNum := nums[0]              // 初始化最大值为第一个元素
	for _, num := range nums[1:] { // 从第二个元素开始遍历
		maxNum = Max(maxNum, num) // 比较当前最大值和当前元素
	}
	return maxNum
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
