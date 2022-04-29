package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	/*
	   				 10
	               /   \
	              8     20
	             / \      \
	            7   9      25
	           /          /  \
	          2          24  29
	*/

	bst := InitBST()
	var minV int
	bst.Insert(10)
	minV = min(bst.root)
	assert.Equal(t, 10, minV)

	bst.Insert(8)
	minV = min(bst.root)
	assert.Equal(t, 8, minV)

	bst.Insert(7)
	minV = min(bst.root)
	assert.Equal(t, 7, minV)

	bst.Insert(9)
	minV = min(bst.root)
	assert.Equal(t, 7, minV)

	bst.Insert(20)
	minV = min(bst.root)
	assert.Equal(t, 7, minV)

	bst.Insert(25)
	minV = min(bst.root)
	assert.Equal(t, 7, minV)

	bst.Insert(24)
	minV = min(bst.root)
	assert.Equal(t, 7, minV)

	bst.Insert(29)
	minV = min(bst.root)
	assert.Equal(t, 7, minV)

	bst.Insert(2)
	minV = min(bst.root)
	assert.Equal(t, 2, minV)
}
