package tree

import (
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
