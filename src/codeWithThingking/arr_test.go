package codewiththingking_test

import (
	codewiththingking "go_alg/src/codeWithThingking"
	"testing"
)

func TestSearch(t *testing.T) {
	input_arr := []int{1, 2, 3, 4, 5, 7}
	target := 7
	expected := 5
	actual := codewiththingking.Search(input_arr, target)
	if actual != expected {
		t.Errorf("7 case is Error ,expected is %d ,actual is %d", expected, actual)
	}

	target = 1
	expected = 0
	actual = codewiththingking.Search(input_arr, target)
	if actual != expected {
		t.Errorf("1 case is Error ,expected is %d ,actual is %d", expected, actual)
	}

	target = 3
	expected = 2
	actual = codewiththingking.Search(input_arr, target)
	if actual != expected {
		t.Errorf("3 case is Error ,expected is %d ,actual is %d", expected, actual)
	}

}

// 测试自己的sqrt
func TestMySqrt(t *testing.T) {
	input := 7
	expected := 2
	actual := codewiththingking.MySqrt(input)
	if expected != actual {
		t.Errorf("4 case is Error ,expected is %d ,actual is %d", expected, actual)
	}
}

// 测试自己的IsPerfectSquare
func TestIsPerfectSquare(t *testing.T) {
	input := 16
	expected := true
	actual := codewiththingking.IsPerfectSquare(input)
	if expected != actual {
		t.Errorf("16 case is Error ,expected is %v ,actual is %v", expected, actual)
	}
}

// 测试自己的RemoveDuplicates
func TestRemoveDuplicates(t *testing.T) {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expected := 5
	actual := codewiththingking.RemoveDuplicates(input)
	if expected != actual {
		t.Errorf("expected is %d,actual is %d", expected, actual)
	}
	if input[0] != 0 {
		t.Errorf("input[0] expected is 0 , actual is %d ", input[0])
	}
	if input[1] != 1 {
		t.Errorf("input[1] expected is 1 , actual is %d ", input[1])
	}

	if input[2] != 2 {
		t.Errorf("input[2] expected is 2 , actual is %d ", input[2])
	}
	if input[3] != 3 {
		t.Errorf("input[3] expected is 3 , actual is %d ", input[3])
	}
	if input[4] != 4 {
		t.Errorf("input[4] expected is 4 , actual is %d ", input[4])
	}
}
