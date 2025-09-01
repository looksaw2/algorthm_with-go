package codewiththingking

import (
	"math"
	"sort"
)

// 设计的一个Mapper的泛型算法
func Map[T any, U any](slice []T, mapper func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = mapper(v)
	}
	return result
}

// 二分查找对应的>=的第一个数值
func LowerBoundIter(nums []int, target int) int {
	left, right := 0, len(nums)-1
	//默认插入的位置
	ans := len(nums)
	//开始二分查找
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

// 二分查的查找对应的 <=的第一个
func HigherBoundIter(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := len(nums)
	//开始查找
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

// 第一个二分查找算法
func Search(nums []int, target int) int {
	return search_recursive_left(nums, target, 0, len(nums)-1)
}

// 找出规定的范围
func SearchRange(nums []int, target int) []int {
	leftBound := search_recursive_left(nums, target, 0, len(nums)-1)
	rightBound := search_recursive_right(nums, target, 0, len(nums)-1)
	return []int{leftBound, rightBound}
}

// 找出平方根
func MySqrt(x int) int {
	if x < 0 {
		return -1
	} else if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	ans := -1
	//开始从1..n/2中寻找
	left := 1
	right := x / 2
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid > x {
			right = mid - 1
		} else if mid*mid < x {
			ans = mid
			left = mid + 1
		} else {
			return mid
		}
	}
	return ans
}

// 平方数组
func SortedSquares2(nums []int) []int {
	//定义平方函数
	var square func(i int) int
	square = func(i int) int {
		return i * i
	}
	//如果全是正数
	if nums[0] >= 0 {
		return Map(nums, square)
	}
	//如果全是负数
	if nums[len(nums)-1] <= 0 {
		result := make([]int, len(nums))
		for i := 0; i < len(nums); i++ {
			result[i] = nums[len(nums)-1-i] * nums[len(nums)-1-i]
		}
		return result
	}
	//有一部分正数，有一部分的负数
	leftZero := LowerBoundIter(nums, 0)
	rightZero := HigherBoundIter(nums, 0)
	result := make([]int, 0)
	//填充0
	for i := 0; i < (rightZero - leftZero); i++ {
		result = append(result, 0)
	}
	leftZero--
	//开始向周围扩展
	for leftZero >= 0 && rightZero < len(nums) {
		if -nums[leftZero] > nums[rightZero] {
			result = append(result, nums[rightZero]*nums[rightZero])
			rightZero++
		} else if -nums[leftZero] < nums[rightZero] {
			result = append(result, nums[leftZero]*nums[leftZero])
			leftZero--
		} else {
			result = append(result, nums[rightZero]*nums[rightZero])
			rightZero++
		}
	}
	//填充其他的数
	if leftZero < 0 {
		result = append(result, Map(nums[rightZero:], square)...)
	}
	if rightZero >= len(nums) {
		result = append(result, Map(nums[:leftZero+1], square)...)
	}
	return result
}

// 判断是不是完全平方数
func IsPerfectSquare(num int) bool {
	if num < 0 {
		return false
	} else if num == 0 {
		return true
	} else if num == 1 {
		return true
	}
	left := 1
	right := num / 2
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == num {
			return true
		} else if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// 快慢指针移除元素
func RemoveElement(nums []int, val int) int {
	//慢指针
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

// 移除重复的元素
func RemoveDuplicates(nums []int) int {
	//记录重复的数据
	duplicatesValue := make(map[int]struct{})
	//慢指针
	left := 0
	for _, v := range nums {
		//看在map中是否有
		_, ok := duplicatesValue[v]
		if !ok {
			nums[left] = v
			left++
			duplicatesValue[v] = struct{}{}
		}
	}
	return left
}

// 移动零
func MoveZeroes(nums []int) {
	//目标是0
	target := 0
	//慢指针
	left := 0
	for _, v := range nums {
		if v != target {
			nums[left] = v
			left++
		}
	}
	//移动剩下的零
	for i := left; i < len(nums); i++ {
		nums[i] = target
	}
}

// 后退的格子
func BackspaceCompare(s string, t string) bool {
	s1, l1 := removeBackspaceString(s)
	t1, l2 := removeBackspaceString(t)
	if l1 != l2 {
		return false
	}
	if s1 != t1 {
		return false
	}
	return true
}

// 移除后退格子的字符串
func removeBackspaceString(s string) (string, int) {
	//慢指针
	left := 0
	sS := []rune(s)
	for _, v := range sS {
		if v != '#' {
			sS[left] = v
			left++
			continue
		}
		//等于#(空格需要回退)
		if left != 0 {
			left--
		}
	}
	return string(sS)[:left], left
}

func search_recursive_left(nums []int, target int, left int, right int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if target > nums[mid] {
		return search_recursive_left(nums, target, mid+1, right)
	} else if target < nums[mid] {
		return search_recursive_left(nums, target, left, mid-1)
	} else {
		res := search_recursive_left(nums, target, left, mid-1)
		if res != -1 {
			return res
		}
		return mid
	}
}

func search_recursive_right(nums []int, target int, left int, right int) int {
	//越界情况
	if left > right {
		return -1
	}
	//开始搜索
	mid := left + (right-left)/2
	if target > nums[mid] {
		return search_recursive_right(nums, target, mid+1, right)
	} else if target < nums[mid] {
		return search_recursive_right(nums, target, left, mid-1)
	} else {
		res := search_recursive_right(nums, target, mid+1, right)
		if res != -1 {
			return res
		}
		return mid
	}
}

func SortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = v * v
	}
	sort.Ints(result)
	return result
}

func MinSubArrayLen(target int, nums []int) int {
	var min func(int, int) int
	min = func(i1, i2 int) int {
		if i1 > i2 {
			return i2
		}
		return i1
	}
	//开始快慢指针
	left := 0
	sum := 0
	ans := math.MaxInt
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		// 如果sum大于target开始缩小范围
		for sum > target {
			//如果i和left相等
			ans = min(ans, i-left+1)
			if i == left {
				return 1
			}
			//移除left
			sum -= nums[left]
			left++
		}
		// 如果sum == target
		if sum == target {
			ans = min(ans, i-left+1)
		}
	}
	if ans == math.MaxInt {
		return 0
	}
	return ans
}
