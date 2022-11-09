package main

import "fmt"

func main() {
	main := ListNode{
		0,
		&ListNode{},
	}
	head := &main
	for _, i := range []int{0, 1, 2, 2, 1, 0} {
		head = head.Next
		head.Val = i
		head.Next = &ListNode{}
	}
	fmt.Println(isPalindrome(&main))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var prev *ListNode
	for slow != nil {
		tmp := slow.Next
		slow.Next = prev
		prev = slow
		slow = tmp
	}
	left, right := head, prev
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}
