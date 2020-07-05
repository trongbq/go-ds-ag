package tree

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
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
	root, err := buildTestingTree(randomArray(10))
	if err != nil {
		t.Errorf("Inserting Error: %v", err)
		return
	}

	if node := Search(root, 5); node == nil {
		t.Error("Can not find existing value")
	}
	if node := Search(root, 999); node != nil {
		t.Error("Found node not in tree")
	}
}

func TestFindMinimum(t *testing.T) {
	randArr := randomArray(10)
	root, err := buildTestingTree(randArr)
	if err != nil {
		t.Errorf("Inserting Error: %v", err)
		return
	}

	sort.Ints(randArr)
	mNode := FindMinimum(root)
	if ma := randArr[0]; mNode != nil && ma != mNode.Value {
		t.Errorf("Failed to find minimum error")
	}

	// Min value is root
	root, err = buildTestingTree(randArr)
	if err != nil {
		t.Errorf("Inserting Error: %v", err)
		return
	}
	mNode = FindMinimum(root)
	if mNode != nil && root.Value != mNode.Value {
		t.Errorf("Failed to find minimum error")
	}
}

func TestDelete(t *testing.T) {
	root, err := buildTestingTree([]int{5, 6, 2, 3, 9, 8, 1, 7, 4, 13, 11, 14, 12})
	if err != nil {
		t.Errorf("Inserting Error: %v", err)
		return
	}

	// Original: [[[<empty> 1 <empty>] 2 [<empty> 3 [<empty> 4 <empty>]]] 5 [<empty> 6 [[[<empty> 7 <empty>] 8 <empty>] 9 [[<empty> 11 [<empty> 12 <empty>]] 13 [<empty> 14 <empty>]]]]]

	Delete(root, 1)
	as := ToString(root)
	if as != "[[<empty> 2 [<empty> 3 [<empty> 4 <empty>]]] 5 [<empty> 6 [[[<empty> 7 <empty>] 8 <empty>] 9 [[<empty> 11 [<empty> 12 <empty>]] 13 [<empty> 14 <empty>]]]]]" {
		t.Error("Delete failed no child node")
	}

	Delete(root, 8)
	as = ToString(root)
	if as != "[[<empty> 2 [<empty> 3 [<empty> 4 <empty>]]] 5 [<empty> 6 [[<empty> 7 <empty>] 9 [[<empty> 11 [<empty> 12 <empty>]] 13 [<empty> 14 <empty>]]]]]" {
		t.Error("Delete failed left node")
	}

	Delete(root, 3)
	as = ToString(root)
	if as != "[[<empty> 2 [<empty> 4 <empty>]] 5 [<empty> 6 [[<empty> 7 <empty>] 9 [[<empty> 11 [<empty> 12 <empty>]] 13 [<empty> 14 <empty>]]]]]" {
		t.Error("Delete failed right child node")
	}

	Delete(root, 9)
	as = ToString(root)
	if as != "[[<empty> 2 [<empty> 4 <empty>]] 5 [<empty> 6 [[<empty> 7 <empty>] 11 [[<empty> 12 <empty>] 13 [<empty> 14 <empty>]]]]]" {
		t.Error("Delete failed both child node")
	}
}

func TestTraverseInOrder(t *testing.T) {
	randArr := randomArray(10)
	root, err := buildTestingTree(randArr)
	if err != nil {
		t.Errorf("Inserting Error: %v", err)
		return
	}

	o := TraverseInOrder(root)
	fmt.Println(o)
	for i, v := range o {
		if i+1 != v {
			t.Error("Error traverse in order")
		}
	}
}

func buildTestingTree(arr []int) (*BinaryTreeNode, error) {
	root := BinaryTreeNode{Value: arr[0]}
	for i := 1; i < len(arr); i++ {
		err := Insert(&root, arr[i])
		if err != nil {
			return nil, err
		}
	}
	return &root, nil
}

func randomArray(n int) []int {
	arr := make([]int, n, n)
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr
}
