package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumBST(t *testing.T) {
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
	assert.Equal(t, 10, sumBST(bst.Root))
	bst.Insert(8)
	assert.Equal(t, 18, sumBST(bst.Root))
	bst.Insert(20)
	assert.Equal(t, 38, sumBST(bst.Root))
	bst.Insert(7)
	assert.Equal(t, 45, sumBST(bst.Root))
	bst.Insert(9)
	assert.Equal(t, 54, sumBST(bst.Root))
	bst.Insert(25)
	assert.Equal(t, 79, sumBST(bst.Root))
	bst.Insert(24)
	assert.Equal(t, 103, sumBST(bst.Root))
	bst.Insert(29)
	assert.Equal(t, 132, sumBST(bst.Root))
}
