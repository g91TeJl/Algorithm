package main

import (
	"fmt"
)

type ListNode struct {
	val  int
	Next *ListNode
}

func initList() *ListNode {
	return &ListNode{}
}

func (l *ListNode) Add(num int) {
	if l == nil {
		l = &ListNode{val: num}
	} else {
		cur := l
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = &ListNode{val: num}
	}
	return
}

func (l *ListNode) Print() {
	current := l
	for current != nil {
		fmt.Println(current.val)
		current = current.Next
	}

}
func Println(link *ListNode) {
	current := link
	for current.Next != nil {
		fmt.Println(current.val)
		current = current.Next
	}
}

func pairSum(head *ListNode) int {
	var result_slice []int
	var sum int
	current := head
	cur := head
	var first_val int
	for cur != nil {
		result_slice = append(result_slice, current.val)
		current = current.Next
		cur = cur.Next.Next
	}
	fmt.Println(result_slice)
	for current != nil {
		first_val, result_slice = result_slice[len(result_slice)-1], result_slice[:len(result_slice)-1]
		if current.val+first_val > sum {
			sum = current.val + first_val
		}
		current = current.Next
	}
	return sum
}

func main() {
	a := initList()
	a.val = 5
	a.Add(4)
	a.Add(2)
	a.Add(1)
	a.Add(2)
	a.Add(10)
	a.Add(2)
	a.Add(3)
	a.Print()
	fmt.Println(pairSum(a))
}
