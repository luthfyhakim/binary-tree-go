package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t := &TreeNode{Val: 8}

	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(4)
	t.Insert(5)
	t.Insert(6)
	t.Insert(7)

	t.Find(11)

	t.Delete(5)
	t.Delete(7)

	t.PrintInorder()
	fmt.Println("")

	fmt.Println("min is", t.FindMin())
	fmt.Println("max is", t.FindMax())
}

// PrintInorder() prints the elements in order
func (t *TreeNode) PrintInorder() {
	if t == nil {
		return
	}

	t.Left.PrintInorder()
	fmt.Print(t.Val)
	t.Right.PrintInorder()
}

// Insert() inserts a new node into the binary tree while adhering to the rules of a perfect BST
func (t *TreeNode) Insert(value int) error {
	if t == nil {
		return errors.New("tree is nil")
	}

	if t.Val == value {
		return errors.New("this node value already exists")
	}

	if t.Val > value {
		if t.Left == nil {
			t.Left = &TreeNode{Val: value}
			return nil
		}

		return t.Left.Insert(value)
	}

	if t.Val < value {
		if t.Right == nil {
			t.Right = &TreeNode{Val: value}
			return nil
		}

		return t.Right.Insert(value)
	}

	return nil
}

// Find() finds the treenode for the given node val
func (t *TreeNode) Find(value int) (TreeNode, bool) {
	if t == nil {
		return TreeNode{}, false
	}

	switch {
	case value == t.Val:
		return *t, true
	case value < t.Val:
		return t.Left.Find(value)
	default:
		return t.Right.Find(value)
	}
}

// Delete() removes the item with value from the tree
func (t *TreeNode) Delete(value int) {
	t.remove(value)
}

func (t *TreeNode) remove(value int) *TreeNode {
	if t == nil {
		return nil
	}

	if value < t.Val {
		t.Left = t.Left.remove(value)
		return t
	}

	if value > t.Val {
		t.Right = t.Right.remove(value)
		return t
	}

	if t.Left == nil && t.Right == nil {
		t = nil
		return nil
	}

	if t.Left == nil {
		t = t.Right
		return t
	}

	if t.Right == nil {
		t = t.Left
		return t
	}

	smallestValOnRight := t.Right
	for {
		// find the smallest value on the right side
		if smallestValOnRight != nil && smallestValOnRight.Left != nil {
			smallestValOnRight = smallestValOnRight.Left
		} else {
			break
		}
	}

	t.Val = smallestValOnRight.Val
	t.Right = t.Right.remove(t.Val)
	return t
}

// FindMax() finds the max element in the given BST
func (t *TreeNode) FindMax() int {
	if t.Right == nil {
		return t.Val
	}

	return t.Right.FindMax()
}

// FindMin() finds the min element in the given BST
func (t *TreeNode) FindMin() int {
	if t.Left == nil {
		return t.Val
	}

	return t.Left.FindMin()
}
