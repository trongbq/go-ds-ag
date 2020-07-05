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
	return fmt.Sprintf("[%v;%v;%v]", n.Value, n.Left, n.Right)
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
			if v == 11 && node.Value == 10 {
				fmt.Println("Left")
			}
			parent = node
			node = node.Right
		case v < node.Value:
			if v == 11 && node.Value == 10 {
				fmt.Println("Right")
			}
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

func Delete(root *BinaryTreeNode, v int) {
	// Find node to be deleted
	node := root
	for node != nil && node.Value != v {
		if v > node.Value {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	if node == nil {
		return // Tree doesn't contains this value
	}
	updateParent := func(n *BinaryTreeNode, nn *BinaryTreeNode) {
		if n.Parent == nil {
			return
		}
		if n.Value > n.Parent.Value {
			n.Parent.Right = nn
		} else {
			n.Parent.Left = nn
		}
	}
	// No children
	if node.Left == nil && node.Right == nil {
		updateParent(node, nil)
		return
	}
	// If tree has one left child
	if node.Left != nil && node.Right == nil {
		node.Left.Parent = node.Parent
		updateParent(node, node.Left)
		return
	}
	// If tree has one right child
	if node.Right != nil && node.Left == nil {
		node.Right.Parent = node.Parent
		updateParent(node, node.Right)
		return
	}
	// Tree has two children
	// Find smallest value in right subtree in sorted order
	minNode := FindMinimum(node.Right)
	// minNode can only have right child, let's link right child to parent of minNode
	if minNode.Right != nil {
		minNode.Right.Parent = minNode.Parent
	}
	updateParent(minNode, minNode.Right)
	// Replace node with new node
	minNode.Parent = node.Parent
	updateParent(node, minNode)
	minNode.Left = node.Left
	minNode.Right = node.Right
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

func ToString(root *BinaryTreeNode) string {
	s := ""
	empty := "<empty>"
	if root == nil {
		return empty
	}
	s += "["
	s += ToString(root.Left)
	s += fmt.Sprintf(" %v ", root.Value)
	s += ToString(root.Right)
	s += "]"
	return s
}
