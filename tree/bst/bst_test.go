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
	assert.Equal(t, bst.Root.Val, 10)
	assert.Nil(t, bst.Root.Left)
	assert.Nil(t, bst.Root.Right)

	bst.Insert(8)
	assert.Equal(t, bst.Root.Left.Val, 8)
	bst.Insert(7)
	assert.Equal(t, bst.Root.Left.Left.Val, 7)
	bst.Insert(9)
	assert.Equal(t, bst.Root.Left.Right.Val, 9)

	bst.Insert(20)
	assert.Equal(t, bst.Root.Right.Val, 20)
	bst.Insert(25)
	assert.Equal(t, bst.Root.Right.Right.Val, 25)
	bst.Insert(24)
	assert.Equal(t, bst.Root.Right.Right.Left.Val, 24)
	bst.Insert(29)
	assert.Equal(t, bst.Root.Right.Right.Right.Val, 29)

}

func createTree() *Bst {
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
	result := BFS(bst.Root)
	expectdOuput := []int{10, 8, 20, 7, 9, 25, 24, 29}
	assert.Equal(t, expectdOuput, result)
}

func TestDFS(t *testing.T) {
	bst := createTree()
	root := bst.Root
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
