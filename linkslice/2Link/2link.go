package main

import "fmt"

type ele struct {
	next *ele
	prev *ele
	val  int
}

type Link2 struct {
	next *Link2
	prev *Link2
	root *ele
	len  int
}

func initList() *Link2 {
	return &Link2{}
}

func (l *Link2) PushFront(v int) {
	ele := &ele{
		val: v,
	}
	if l.len == 0 {
		l.root = ele
	} else {
		current := l.root
		for current.prev != nil {
			current = current.prev
		}
		current.prev = ele
		current.next = current
	}
	l.len++
	return
}



func (l *Link2) Print() error {
	if l.len == 0 {
		fmt.Errorf("")
	}
	current := l.root
	for current.prev != nil {
		fmt.Println("current -> ", current.next.val, "next -> ", current.prev.val)
		current = current.prev
	}
	return nil

}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	link2 := initList()

	for _, x := range arr {
		link2.PushFront(x)
	}
	link2.Print()
}
