package lettcode100test

import (
	"go_alg/src/leetcode100"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindContentChildren1(t *testing.T) {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	expected := 1
	actual := leetcode100.FindContentChildren(g, s)
	require.Equal(t, expected, actual)
}

func TestFindContentChildren2(t *testing.T) {
	g := []int{1, 2}
	s := []int{1, 2, 3}
	expected := 2
	actual := leetcode100.FindContentChildren(g, s)
	require.Equal(t, expected, actual)
}

func TestNumIslands1(t *testing.T) {
	input := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	expected := 1
	actual := leetcode100.NumIslands(input)
	require.Equal(t, expected, actual)
}

func TestNumIslands2(t *testing.T) {
	input := [][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}
	expected := 1
	actual := leetcode100.NumIslands(input)
	require.Equal(t, expected, actual)
}

func TestMinSubArrayLen1(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	expected := 2
	actual := leetcode100.MinSubArrayLen(target, nums)
	require.Equal(t, expected, actual)
}

func TestGenerateMatrix1(t *testing.T) {
	input := 3
	expected := [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
	}
	actual := leetcode100.GenerateMatrix(input)
	require.Equal(t, expected, actual)
}

func TestCombineWithNewSlice(t *testing.T) {
	n := 4
	k := 2
	expected := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{3, 4},
	}
	actual := leetcode100.CombineWithNewSlice(n, k)
	require.Equal(t, expected, actual)
}

func TestGetDecimalValue(t *testing.T) {
	head := leetcode100.ListNode{Val: 1}
	node1 := leetcode100.ListNode{Val: 0}
	node2 := leetcode100.ListNode{Val: 1}
	head.Next = &node1
	node1.Next = &node2
	actual := leetcode100.GetDecimalValue(&head)
	expected := 5
	require.Equal(t, expected, actual)
}

func arryToLinkList(t *testing.T, nums []int) *leetcode100.ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &leetcode100.ListNode{
		Val: nums[0],
	}
	prev := head
	for i := 1; i < len(nums); i++ {
		node := leetcode100.ListNode{
			Val: nums[i],
		}
		prev.Next = &node
		prev = prev.Next
	}
	return head
}

func TestNodesBetweenCriticalPoints(t *testing.T) {
	input := []int{5, 3, 1, 2, 5, 1, 2}
	head := arryToLinkList(t, input)
	expected := []int{1, 3}
	actual := leetcode100.NodesBetweenCriticalPoints(head)
	require.Equal(t, expected, actual)
}

func TestSumOfLeftLeaves(t *testing.T) {
	root := leetcode100.TreeNode{Val: 3}
	node1 := leetcode100.TreeNode{Val: 9}
	node2 := leetcode100.TreeNode{Val: 20}
	node21 := leetcode100.TreeNode{Val: 15}
	node22 := leetcode100.TreeNode{Val: 7}
	root.Left = &node1
	root.Right = &node2
	node2.Left = &node21
	node2.Right = &node22
	expected := 24
	actual := leetcode100.SumOfLeftLeaves(&root)
	require.Equal(t, expected, actual)

}
