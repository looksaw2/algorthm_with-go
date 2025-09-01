package leetcode100

type auxTem struct {
	Val   int
	Index int
}

func DailyTempratures(temperatures []int) []int {
	if len(temperatures) == 0 {
		return []int{}
	}
	//需要返回的答案
	ans := make([]int, len(temperatures))
	// 辅助栈
	aux := make([]auxTem, 0)
	for i, v := range temperatures {
		if len(aux) == 0 {
			aux = append(aux, auxTem{Val: v, Index: i})
		}
		for len(aux) != 0 && v > aux[len(aux)-1].Val {
			t := aux[len(aux)-1]
			aux = aux[:len(aux)-1]
			ans[t.Index] = i - t.Index
		}
		if len(aux) == 0 {
			aux = append(aux, auxTem{Val: v, Index: i})
			continue
		}
		aux = append(aux, auxTem{Val: v, Index: i})
	}
	return ans
}

func FinalPrice(prices []int) []int {
	if len(prices) == 0 {
		return []int{}
	}
	//需要返回的数值
	ans := make([]int, len(prices))
	//单调栈
	aux := make([]auxTem, 0)
	for i, v := range prices {
		if len(aux) == 0 {
			aux = append(aux, auxTem{Val: v, Index: i})
			continue
		}
		for len(aux) != 0 && v < aux[len(aux)-1].Val {
			t := aux[len(aux)-1]
			aux = aux[:len(aux)-1]
			ans[t.Index] = t.Val - v
		}

		if len(aux) == 0 {
			aux = append(aux, auxTem{Val: v, Index: i})
			continue
		}

		aux = append(aux, auxTem{Val: v, Index: i})
	}
	// 释放aux
	for _, v := range aux {
		ans[v.Index] = v.Val
	}
	return ans
}

func NextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	//ans数组
	ans := make([]int, len(nums))
	//aux数组
	aux := make([]auxTem, 0)
	//第一次过数组
	for i, v := range nums {
		if len(aux) == 0 {
			aux = append(aux, auxTem{Val: v, Index: i})
			continue
		}
		for len(aux) != 0 && v > aux[len(aux)-1].Val {
			t := aux[len(aux)-1]
			aux = aux[:len(aux)-1]
			ans[t.Index] = v
		}
		aux = append(aux, auxTem{Val: v, Index: i})
	}
	//第二次过数组
	for _, v := range nums {
		if len(aux) == 0 {
			continue
		}
		for len(aux) != 0 && v > aux[len(aux)-1].Val {
			t := aux[len(aux)-1]
			aux = aux[:len(aux)-1]
			ans[t.Index] = v
		}
	}
	//处理剩余的数值
	for _, v := range aux {
		ans[v.Index] = -1
	}
	return ans
}
