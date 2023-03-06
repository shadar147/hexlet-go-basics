package ex27

import "fmt"

type ListNode struct {
	Next *ListNode
	Val  int
}

func (head *ListNode) Reverse() *ListNode {
	if head == nil {
		return nil
	}

	var prevNode *ListNode
	h := &ListNode{
		Next: head.Next,
		Val:  head.Val,
	}

	for h.Next != nil {
		nextNode := &ListNode{
			Next: h.Next.Next,
			Val:  h.Next.Val,
		}

		h.Next = prevNode

		prevNode = h
		h = nextNode
	}

	h.Next = prevNode

	return h
}

func (head *ListNode) Display() {
	h := head
	for h != nil {
		fmt.Printf("%v ->", h.Val)
		h = h.Next
	}
	fmt.Println()
}
