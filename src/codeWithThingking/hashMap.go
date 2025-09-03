package codewiththingking

func IsAnagram(s string, r string) bool {
	sMap := make(map[rune]int)
	//遍历s
	for _, sS := range s {
		sMap[sS]++
	}
	for _, v := range r {
		_, ok := sMap[v]
		if !ok {
			return false
		}
		sMap[v]--
	}
	//验证次数
	for _, v := range sMap {
		if v != 0 {
			return false
		}
	}
	return true
}

// 判断异
func CanConstruct(ransomNote string, magazine string) bool {
	mMap := make(map[rune]int)
	for _, v := range magazine {
		mMap[v] += 1
	}
	for _, v := range ransomNote {
		_, ok := mMap[v]
		if !ok {
			return false
		}
		mMap[v]--
	}
	for _, v := range mMap {
		if v < 0 {
			return false
		}
	}
	return true
}

type MyQueue[T any] struct {
	StackIn  []T //第一个进入的栈
	StackOut []T //
}

func Construct[T any]() MyQueue[T] {
	return MyQueue[T]{
		StackIn:  make([]T, 0),
		StackOut: make([]T, 0),
	}
}

// 加入数组
func (this *MyQueue[T]) Push(x T) {
	this.StackIn = append(this.StackIn, x)
}

// 移除元素
func (this *MyQueue[T]) Pop() T {
	//如果入队也为空
	if len(this.StackOut) == 0 {
		if len(this.StackIn) == 0 {
			panic("The Queue is empty.......")
		}
		//反转
		for i := len(this.StackIn) - 1; i >= 0; i-- {
			x := this.StackIn[i]
			this.StackOut = append(this.StackOut, x)
		}
		//清空stdIn
		this.StackIn = []T{}
	}
	//共同逻辑
	q := this.StackOut[len(this.StackOut)-1]
	this.StackOut = this.StackOut[:len(this.StackOut)-1]
	return q
}

// Peek
func (this *MyQueue[T]) Peek() T {
	//如果入队也为空
	if len(this.StackOut) == 0 {
		if len(this.StackIn) == 0 {
			panic("The Queue is empty.......")
		}
		//反转
		for i := len(this.StackIn) - 1; i >= 0; i-- {
			x := this.StackIn[i]
			this.StackOut = append(this.StackOut, x)
		}
	}
	//共同逻辑
	q := this.StackOut[len(this.StackOut)-1]
	return q
}

//

func (this *MyQueue[T]) Empty() bool {
	return len(this.StackIn) == 0 && len(this.StackOut) == 0
}

// 用Queue来实现
type MyStack struct {
	Arr []int
}

func NewMyStack() MyStack {
	return MyStack{
		Arr: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.Arr = append(this.Arr, x)
}

func (this *MyStack) Pop() int {
	q := this.Arr[len(this.Arr)-1]
	this.Arr = this.Arr[:len(this.Arr)-1]
	return q
}

func (this *MyStack) Top() int {
	return this.Arr[len(this.Arr)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.Arr) == 0
}

// 判断括号是否有效
func IsValid(s string) bool {
	stack := make([]rune, 0)
	for _, v := range s {
		switch v {
		case '(':
			stack = append(stack, v)
		case '{':
			stack = append(stack, v)
		case '[':
			stack = append(stack, v)
		case ')':
			if len(stack) == 0 {
				return false
			}
			q := stack[len(stack)-1]
			if q != '(' {
				return false
			}
			if q == '(' {
				stack = stack[:len(stack)-1]
			}
		case ']':
			if len(stack) == 0 {
				return false
			}
			q := stack[len(stack)-1]
			if q != '[' {
				return false
			}
			if q == '[' {
				stack = stack[:len(stack)-1]
			}
		case '}':
			if len(stack) == 0 {
				return false
			}
			q := stack[len(stack)-1]
			if q != '{' {
				return false
			}
			if q == '{' {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}

// 用栈移除
func IRemoveDuplicates(s string) string {
	stack := make([]rune, 0)
	if len(s) == 0 {
		return ""
	}
	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}
		if v == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return string(stack)
}
