package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	/*
	   				 10
	               /   \
	              8     20
	             / \      \
	            7   9      25
	                      /  \
	                     24  29
	*/

	bst := InitBST()
	bst.Insert(10)
	assert.Equal(t, bst.root.value, 10)
	assert.Nil(t, bst.root.left)
	assert.Nil(t, bst.root.right)

	bst.Insert(8)
	assert.Equal(t, bst.root.left.value, 8)
	bst.Insert(7)
	assert.Equal(t, bst.root.left.left.value, 7)
	bst.Insert(9)
	assert.Equal(t, bst.root.left.right.value, 9)

	bst.Insert(20)
	assert.Equal(t, bst.root.right.value, 20)
	bst.Insert(25)
	assert.Equal(t, bst.root.right.right.value, 25)
	bst.Insert(24)
	assert.Equal(t, bst.root.right.right.left.value, 24)
	bst.Insert(29)
	assert.Equal(t, bst.root.right.right.right.value, 29)

}

func createTree() *bst {
	/*
	   				 10
	               /   \
	              8     20
	             / \      \
	            7   9      25
	                      /  \
	                     24  29
	*/

	bst := InitBST()
	bst.Insert(10)

	bst.Insert(8)
	bst.Insert(7)
	bst.Insert(9)

	bst.Insert(20)
	bst.Insert(25)
	bst.Insert(24)
	bst.Insert(29)
	return bst
}

func TestBFS(t *testing.T) {
	bst := createTree()
	result := bst.BFS()
	expectdOuput := []int{10, 8, 20, 7, 9, 25, 24, 29}
	assert.Equal(t, expectdOuput, result)
}

func TestDFS(t *testing.T) {
	bst := createTree()
	root := bst.root
	var dFSPreOderResult []int
	DFSPreOrder(root, &dFSPreOderResult)
	expectdDFSPreOrderResult := []int{10, 8, 7, 9, 20, 25, 24, 29}
	assert.Equal(t, expectdDFSPreOrderResult, dFSPreOderResult)

	var dFSInOderResult []int
	expectedDFSInOrderResult := []int{7, 8, 9, 10, 20, 24, 25, 29}
	DFSInOrder(root, &dFSInOderResult)
	assert.Equal(t, expectedDFSInOrderResult, dFSInOderResult)

	var dFSPostOderResult []int
	expectedDFSPostOrderResult := []int{7, 9, 8, 24, 29, 25, 20, 10}
	DFSPostOrder(root, &dFSPostOderResult)
	assert.Equal(t, expectedDFSPostOrderResult, dFSPostOderResult)

}
