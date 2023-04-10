package main

import (
	"fmt"
)

type ele struct {
	next  *ele
	value int
}

type singleList struct {
	len  int
	root *ele
	next *singleList
}

func initList() *singleList {
	return &singleList{}
}

func (s *singleList) AddFront(val int) {
	ele := &ele{
		value: val,
	}
	if s.root == nil {
		s.root = ele
	} else {
		ele.next = s.root
		s.root = ele
	}
	s.len++
	return
}

func (s *singleList) AddBack(val int) {
	ele := &ele{
		value: val,
	}
	if s.root == nil {
		s.root = ele
	} else {
		current := s.root
		for current.next != nil {
			current = current.next
		}
		current.next = ele
	}
	s.len++
	return
}

func (s *singleList) Reverse() *singleList {
	var cur = s
	var prev *singleList
	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	return prev
}

func (s *singleList) Print() error {
	if s.root == nil {
		return fmt.Errorf("asdsda")
	}
	current := s.root
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
	return nil
}

func PrintAllListElements(s *singleList) {
	currentList := s
	for currentList != nil {
		fmt.Println("-------")
		for currentList.root != nil {
			fmt.Println(currentList.root.value)
			currentList.root = currentList.root.next
		}
		currentList = currentList.next
	}
}

func main() {
	singleList := initList()
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	for _, x := range arr {
		singleList.AddBack(x)
	}
	singleList.Print()
	fmt.Println()
	singleList2 := initList()
	arr1 := []int{7, 6, 5, 4, 3, 2, 1}
	for _, x := range arr1 {
		singleList2.AddBack(x)
	}
	singleList3 := initList()
	arr3 := []int{2, 3, 1, 4, 2}

	for _, x := range arr3 {
		singleList3.AddBack(x)
	}
	singleList.next = singleList2
	singleList2.next = singleList3
	singleList.Print()
	PrintAllListElements(singleList)

}
