package sort

import "container/list"

// 选择排序
func SelectionSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		k := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[k] {
				k = j
			}
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
}

// 冒泡排序
func BubbleSort(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// 插入排序
func InsertionSort(nums []int) {
	n := len(nums)
	for i := 1; i < n; i++ {
		base := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > base {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = base
	}
}

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

// 构建二叉树
func dfsBuildTree(preorder []int, inorderMap map[int]int, i int, l int, r int) *TreeNode {
	//如果r < l时终止
	if r < l {
		return nil
	}
	//根节点
	root := NewTreeNode(preorder[i])
	//根据上面得到的root划分中序遍历
	m := inorderMap[preorder[i]]
	root.Left = dfsBuildTree(preorder, inorderMap, i+1, l, m-1)
	root.Right = dfsBuildTree(preorder, inorderMap, i+1+m-l, m+1, r)
	return root
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int, len(inorder))
	for i, v := range inorder {
		inorderMap[v] = i
	}
	root := dfsBuildTree(preorder, inorderMap, 0, 0, len(inorder)-1)
	return root
}

// 汉诺塔
func move(src *list.List, tar *list.List) {
	pan := src.Back()
	tar.PushBack(pan)
	src.Remove(pan)
}

func dfsHanota(i int, src *list.List, buf *list.List, tar *list.List) {
	if i == 1 {
		move(src, tar)
		return
	}
	dfsHanota(i-1, src, tar, buf)
	move(src, tar)
	dfsHanota(i-1, buf, src, tar)
}

func SolveHanota(A *list.List, B *list.List, C *list.List) {
	n := A.Len()
	if n == 0 {
		return
	}
	dfsHanota(n, A, B, C)
}

func backtrackI(state *[]int, choices *[]int, selected *[]bool, res *[][]int) {
	//choices的长度和state相等
	if len(*state) == len(*choices) {
		newState := append([]int{}, *state...)
		*res = append(*res, newState)
	}
	for i := 0; i < len(*choices); i++ {
		choice := (*choices)[i]
		if !(*selected)[i] {
			(*selected)[i] = true
			*state = append(*state, choice)
			backtrackI(state, choices, selected, res)
			//开始回溯
			(*selected)[i] = false
			*state = (*state)[:len(*state)-1]
		}
	}
}

func PermutationsI(nums []int) [][]int {
	res := make([][]int, 0)
	state := make([]int, 0)
	selected := make([]bool, len(nums))
	backtrackI(&state, &nums, &selected, &res)
	return res
}

func backtrackII(state *[]int, choices *[]int, selected *[]bool, res *[][]int) {
	if len(*state) == len(*choices) {
		newState := append([]int{}, *state...)
		*res = append(*res, newState)
	}
	duplicated := make(map[int]struct{}, 0)
	for i := 0; i < len(*choices); i++ {
		choice := (*choices)[i]
		if _, ok := duplicated[choice]; !ok && !(*selected)[i] {
			duplicated[choice] = struct{}{}
			(*selected)[i] = true
			*state = append(*state, choice)
			backtrackII(state, choices, selected, res)
			//开始回溯
			(*selected)[i] = false
			*state = (*state)[:len(*state)-1]
		}
	}
}

func PermutationsII(nums []int) [][]int {
	res := make([][]int, 0)
	state := make([]int, 0)
	selected := make([]bool, len(nums))
	backtrackII(&state, &nums, &selected, &res)
	return res
}

func backtrackSubsetSumINaive(total int, target int, state *[]int, choices *[]int, res *[][]int) {
	if total == target {
		newState := append([]int{}, *state...)
		*res = append(*res, newState)
		return
	}
	for i := 0; i < len(*choices); i++ {
		if total+(*choices)[i] > target {
			return
		}
		*state = append(*state, (*choices)[i])
		backtrackSubsetSumINaive(total+(*choices)[i], target, state, choices, res)
		*state = (*state)[:len(*state)-1]
	}
}

func SubsetSumINative(nums []int, target int) [][]int {
	state := make([]int, 0)
	total := 0
	res := make([][]int, 0)
	backtrackSubsetSumINaive(total, target, &state, &nums, &res)
	return res
}
