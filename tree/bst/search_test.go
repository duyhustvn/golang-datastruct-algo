package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchBST(t *testing.T) {
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
	var realNode *TreeNode
	realNode = searchBST(bst.Root, 10)
	assert.Equal(t, 10, realNode.Val)

	bst.Insert(8)
	realNode = searchBST(bst.Root, 10)
	assert.Equal(t, 10, realNode.Val)
	realNode = searchBST(bst.Root, 8)
	assert.Equal(t, 8, realNode.Val)
}
