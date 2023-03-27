package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func TreeInit() *TreeNode {
	return &TreeNode{}
}

func (t *TreeNode) Insert(value int) error {
	if t == nil {
		return errors.New("Tree is nil")
	}
	if t.value == value {
		return errors.New("This node value alredy exists")
	}
	if t.value > value {
		if t.left == nil {
			t.left = &TreeNode{value: value}
			return nil
		}
		return t.left.Insert(value)
	}
	if t.value < value {
		if t.right == nil {
			t.right = &TreeNode{value: value}
			return nil
		}
		return t.right.Insert(value)
	}
	return nil
}

func (t *TreeNode) FindMin() int {
	if t.left == nil {
		return t.value
	}
	return t.left.FindMin()
}

func (t *TreeNode) FindMax() int {
	if t.right == nil {
		return t.value
	}
	return t.right.FindMax()
}

func (t *TreeNode) PrintInorder() {
	if t == nil {
		return
	}
	t.left.PrintInorder()
	fmt.Println(t.value)
	t.right.PrintInorder()
}

func main() {
	tree := TreeNode{}
	tree.value = -1
	fmt.Println(tree.Insert(1))
	fmt.Println(tree.Insert(2))
	tree.PrintInorder()
}
