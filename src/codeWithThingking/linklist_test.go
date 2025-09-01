package codewiththingking_test

import (
	codewiththingking "go_alg/src/codeWithThingking"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	// 开始测试
	head := codewiththingking.ListNode{
		Val: 1,
	}
	node1 := codewiththingking.ListNode{
		Val: 2,
	}
	node2 := codewiththingking.ListNode{
		Val: 6,
	}
	node3 := codewiththingking.ListNode{
		Val: 3,
	}
	node4 := codewiththingking.ListNode{
		Val: 4,
	}
	node5 := codewiththingking.ListNode{
		Val: 5,
	}
	node6 := codewiththingking.ListNode{
		Val: 6,
	}
	head.Next = &node1
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6
	actual := codewiththingking.RemoveElements(&head, 6)
	n := actual
	if n.Val != 1 {
		t.Errorf("[1] is not 1")
	}
	n = n.Next
	if n.Val != 2 {
		t.Errorf("[2] is not 2")
	}
	n = n.Next
	if n.Val != 3 {
		t.Errorf("[3] is not 3")
	}
	n = n.Next
	if n.Val != 4 {
		t.Errorf("[4] is not 4")
	}
	n = n.Next
	if n.Val != 5 {
		t.Errorf("[5] is not 5")
	}
}
