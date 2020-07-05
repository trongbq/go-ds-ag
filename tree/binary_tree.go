package tree

import (
	"errors"
)

// BinaryTreeNode represents a node in Binary Tree
type BinaryTreeNode struct {
	Value  int
	Parent *BinaryTreeNode
	Left   *BinaryTreeNode
	Right  *BinaryTreeNode
}

// Insert value into tree
// Runs in O(h) time where h = logn in balanced tree
func Insert(root *BinaryTreeNode, v int) error {
	if root == nil {
		root = &BinaryTreeNode{Value: v}
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
	// Skip if v == root.Value
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
