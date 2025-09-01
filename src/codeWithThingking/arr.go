package codewiththingking

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
