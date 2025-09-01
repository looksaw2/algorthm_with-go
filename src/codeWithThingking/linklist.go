package codewiththingking

// 定义链表的数据结构
type ListNode struct {
	Val  int
	Next *ListNode
}

// 删除==val的2节点
func RemoveElements(head *ListNode, val int) *ListNode {
	//加上虚拟的头节点
	dummy := &ListNode{}
	dummy.Next = head
	head = dummy
	//开始遍历
	prev := head
	curr := head.Next
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
			curr = curr.Next
			continue
		}
		curr = curr.Next
		prev = prev.Next
	}
	//
	return head.Next
}

// 反转链表
func ReverseList(head *ListNode) *ListNode {

	//初始为
	var prev *ListNode = nil
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// 辅助函数，添加虚拟头节点
func addDummyHead(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	return dummyHead
}
