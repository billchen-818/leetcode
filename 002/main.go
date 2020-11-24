package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// l1 2-4-3
	// l2 5-6-4

	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l := addTwoNumbers(l1, l2)

	l.print()

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := new(ListNode)
	temp := l
	jinwei := 0

	for l1 != nil && l2 != nil {
		p := new(ListNode)
		temp.Next = p
		p.Val = (l1.Val + l2.Val + jinwei) % 10
		jinwei = (l1.Val + l2.Val + jinwei) / 10
		p.Next = nil
		temp = temp.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		p := new(ListNode)
		temp.Next = p
		p.Val = (l1.Val + jinwei) % 10
		jinwei = (l1.Val + jinwei) / 10
		p.Next = nil
		temp = temp.Next
		l1 = l1.Next
	}

	for l2 != nil {
		p := new(ListNode)
		temp.Next = p
		p.Val = (l2.Val + jinwei) % 10
		jinwei = (l2.Val + jinwei) / 10
		p.Next = nil
		temp = temp.Next
		l2 = l2.Next
	}

	if jinwei != 0 {
		p := new(ListNode)
		temp.Next = p
		p.Val = jinwei
		p.Next = nil
		temp = temp.Next
	}

	return l.Next
}

func (l *ListNode) print() {
	for l != nil {
		fmt.Printf("%d", l.Val)
		l = l.Next
	}

	fmt.Println()
}
