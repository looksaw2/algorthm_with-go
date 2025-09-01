package lettcode100test

import (
	"go_alg/src/leetcode100"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDailyTempratures(t *testing.T) {
	input := []int{73, 74, 75, 71, 69, 72, 76, 73}
	expected := []int{1, 1, 4, 2, 1, 1, 0, 0}
	actual := leetcode100.DailyTempratures(input)
	require.Equal(t, expected, actual)
}
func TestFinalPrice(t *testing.T) {
	input := []int{8, 4, 6, 2, 3}
	expected := []int{4, 2, 4, 2, 3}
	actual := leetcode100.FinalPrice(input)
	require.Equal(t, expected, actual)
}

func TestNextGreaterElements1(t *testing.T) {
	input := []int{1, 2, 1}
	expected := []int{2, -1, 2}
	actual := leetcode100.NextGreaterElements(input)
	require.Equal(t, expected, actual)
}

func TestNextGreaterElements2(t *testing.T) {
	input := []int{1, 2, 3, 4, 3}
	expected := []int{2, 3, 4, -1, 4}
	actual := leetcode100.NextGreaterElements(input)
	require.Equal(t, expected, actual)
}
