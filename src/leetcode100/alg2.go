package leetcode100

import (
	"math"
	"sort"
)

func FindContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ans := 0
	gIndex := 0
	for _, v := range s {
		if gIndex >= len(g) {
			break
		}
		g[gIndex] -= v
		if g[gIndex] <= 0 {
			gIndex++
			ans++
		}
	}
	return ans
}

func NumIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	directX := []int{1, 0, -1, 0}
	directY := []int{0, 1, 0, -1}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	ans := 0
	var dfs func(x int, y int)
	dfs = func(x int, y int) {
		for i := 0; i < 4; i++ {
			nextX := x + directX[i]
			nextY := y + directY[i]
			if nextX < 0 || nextX >= len(grid) || nextY < 0 || nextY >= len(grid[0]) {
				continue
			}
			if !visited[nextX][nextY] && grid[nextX][nextY] == '1' {
				visited[nextX][nextY] = true
				dfs(nextX, nextY)
			}
		}
	}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if grid[row][col] == '1' && !visited[row][col] {
				visited[row][col] = true
				ans++
				dfs(row, col)
			}
		}
	}
	return ans
}

func MinSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] >= target {
			return 1
		}
		return 0
	}
	if nums[0] > target {
		return 1
	}
	ans := math.MaxInt
	leftPoint := 0
	rightPoint := 1
	sum := nums[0]
	for rightPoint < len(nums) && leftPoint <= rightPoint {
		//待添加的元素
		toAdd := nums[rightPoint]
		rightPoint++
		sum += toAdd
		if sum < target {
			continue
		}
		for sum >= target && leftPoint <= rightPoint {
			//开始向左边减少元素
			sum -= nums[leftPoint]
			leftPoint++
		}
		if rightPoint-leftPoint < ans {
			ans = rightPoint - leftPoint
		}

	}
	if ans == math.MaxInt {
		return 0
	}
	return ans + 1
}

func GenerateMatrix(n int) [][]int {
	//预分配空间
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	top := 0
	bottom := n - 1
	left := 0
	right := n - 1
	num := 1

	for top <= bottom && left <= right {
		//从左向右填入数值
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++
		// 从上到下的填入数值
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		//从右到左的填入数值
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		//从下到上的填入数值
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}

func CombineWithNewSlice(n int, k int) [][]int {
	var result [][]int
	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		if len(path) == k {
			result = append(result, path)
			return
		}
		for i := start; i <= n; i++ {
			newPath := make([]int, len(path), len(path)+1)
			copy(newPath, path)
			newPath = append(newPath, i)
			backtrack(i+1, newPath)
		}
	}
	backtrack(1, []int{})
	return result
}

func CombineWithStatus(n int, k int) [][]int {
	var result [][]int
	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			backtrack(i+1, path)
			path = path[:len(path)-1]
		}
	}
	backtrack(1, []int{})
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}
	root := dfsBuildTree(preorder, inorderMap, 0, 0, len(inorder)-1)
	return root
}

func dfsBuildTree(preorder []int, inorderMap map[int]int, i int, l int, r int) *TreeNode {
	if l > r {
		return nil
	}
	root := newTreeNode(preorder[i])
	m := inorderMap[preorder[i]]
	root.Left = dfsBuildTree(preorder, inorderMap, i+1, l, m-1)
	root.Right = dfsBuildTree(preorder, inorderMap, i+1+m-l, m+1, r)
	return root
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getLinkLen(head *ListNode) int {
	num := 0
	for node := head; node != nil; node = node.Next {
		num++
	}
	return num
}

func getWeight(num int) int {
	i := 0
	for num/10 != 0 {
		num /= 10
		i++
	}
	return i
}
func GetDecimalValue(head *ListNode) int {
	sum := 0
	n := getLinkLen(head)
	for i := n - 1; i >= 0; i-- {
		sum += head.Val * int(math.Pow(2, float64(i)))
		head = head.Next
	}
	return sum
}

func NodesBetweenCriticalPoints(head *ListNode) []int {
	n := getLinkLen(head)
	if n < 5 {
		return []int{-1, -1}
	}
	re := make([]int, 0)
	prev := head
	index := 0
	for node := head.Next; node.Next != nil; node = node.Next {
		nt := node.Next
		index++
		if (node.Val > prev.Val && node.Val > nt.Val) || (node.Val < prev.Val && node.Val < nt.Val) {
			re = append(re, index)
		}
		prev = node
	}
	if len(re) < 2 {
		return []int{-1, -1}
	}
	min := re[len(re)-1] - re[0]
	max := re[len(re)-1] - re[0]
	for i := 1; i < len(re); i++ {
		n := re[i] - re[i-1]
		if n < min {
			min = n
		}
	}
	return []int{
		min, max,
	}
}

func isLeaf(node *TreeNode) bool {
	if node.Left == nil && node.Right == nil {
		return true
	}
	return false
}

func getLeafArray(root *TreeNode) []int {
	ans := make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if isLeaf(node) {
			ans = append(ans, node.Val)
		}
		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
	}
	dfs(root)
	return ans
}

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	arr1 := getLeafArray(root1)
	arr2 := getLeafArray(root2)
	if len(arr1) != len(arr2) {
		return false
	}
	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}
	return true
}

func NunColor(root *TreeNode) int {
	colorMap := make(map[int]bool)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		colorMap[node.Val] = true
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return len(colorMap)
}

func SumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if isLeaf(node) {
			return
		}
		if node.Left != nil {
			if isLeaf(node.Left) {
				sum += node.Left.Val
			} else {
				dfs(node.Left)
			}
		}
		if node.Right != nil && !isLeaf(node.Right) {
			dfs(node.Right)
		}
	}
	dfs(root)
	return sum
}

func FindSecondMinmumValue(root *TreeNode) int {
	min := root.Val
	secondMin := math.MaxInt
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val > min {
			if node.Val < secondMin {
				secondMin = node.Val
			}
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	if secondMin == math.MaxInt {
		return -1
	}
	return secondMin
}
