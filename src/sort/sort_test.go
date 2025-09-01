package sort_test

import (
	"go_alg/src/sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSelectionSort(t *testing.T) {
	input := []int{6, 5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5, 6}
	sort.SelectionSort(input)
	require.Equal(t, expected, input)
}

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 2, 1, 7}
	inorder := []int{9, 3, 1, 2, 7}
	actualRoot := sort.BuildTree(preorder, inorder)
	require.Equal(t, 3, actualRoot.Value)
	require.Equal(t, 9, actualRoot.Left.Value)
	require.Equal(t, 2, actualRoot.Right.Value)
	require.Equal(t, 1, actualRoot.Right.Left.Value)
	require.Equal(t, 7, actualRoot.Right.Right.Value)
}
