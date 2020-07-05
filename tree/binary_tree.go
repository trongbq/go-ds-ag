package tree

import (
	"errors"
	"fmt"
)

// BinaryTreeNode represents a node in Binary Tree
type BinaryTreeNode struct {
	Value  int
	Parent *BinaryTreeNode
	Left   *BinaryTreeNode
	Right  *BinaryTreeNode
}

func (n BinaryTreeNode) String() string {
	return fmt.Sprintf("Value: %v; Left: %p; Right: %p", n.Value, n.Left, n.Right)
}

// Insert value into tree
// Runs in O(h) time where h = logn in balanced tree
func Insert(root *BinaryTreeNode, v int) error {
	if root == nil {
		return errors.New("Root is nil")
	}
	// Find position for this value
	var parent *BinaryTreeNode
	node := root
	for node != nil {
		switch {
		case v > node.Value:
			parent = node
			node = node.Right
		case v < node.Value:
			parent = node
			node = node.Left
		default:
			return errors.New("Dupplicate value")
		}
	}
	// Found, create new node
	newNode := BinaryTreeNode{Value: v, Parent: parent}
	if v > parent.Value {
		parent.Right = &newNode
	} else {
		parent.Left = &newNode
	}
	return nil
}

// Search searches through tree to find node contains value v
// Runs in O(h) time where h = logn in balanced tree
func Search(root *BinaryTreeNode, v int) *BinaryTreeNode {
	node := root
	for node != nil && node.Value != v {
		switch {
		case v > node.Value:
			node = node.Right
		case v < node.Value:
			node = node.Left
		}
	}
	return node
}

func FindMinimum(root *BinaryTreeNode) *BinaryTreeNode {
	node := root
	for node != nil {
		if node.Left != nil {
			node = node.Left
		} else {
			return node
		}
	}
	return node
}

func TraverseInOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	TraverseInOrder(root.Left)
	fmt.Printf("%v ", root.Value)
	TraverseInOrder(root.Right)
}

// Display shows root-left-right from top to bottom
func Display(root *BinaryTreeNode) {
	queue := make([]*BinaryTreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue[0] = nil
		queue = queue[1:]
		fmt.Printf("%v ", node.Value)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	fmt.Println("")
}
