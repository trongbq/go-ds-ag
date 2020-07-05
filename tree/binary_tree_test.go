package tree

import (
	"testing"
)

func TestInsertingSuccess(t *testing.T) {
	root := BinaryTreeNode{Value: 9}
	err := Insert(&root, 3)
	if err != nil {
		t.Errorf("Inserting Error: %v", err)
	}

	err = Insert(nil, 3)
	if err == nil {
		t.Errorf("Can not detect insert to nil root")
	}

	err = Insert(&root, 3)
	if err == nil {
		t.Errorf("Can not detect inserting dupplicated value")
	}
}

func TestSearching(t *testing.T) {
	root := BinaryTreeNode{Value: 9}
	err := Insert(&root, 3)
	if err != nil {
		t.Errorf("Inserting Error: %v", err)
	}
	err = Insert(&root, 5)
	if err != nil {
		t.Errorf("Inserting Error: %v", err)
	}

	if node := Search(&root, 5); node == nil {
		t.Error("Can not find existing value")
	}
	if node := Search(&root, 1); node != nil {
		t.Error("Found node not in tree")
	}
}
