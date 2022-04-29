package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
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
	var expectedMaxDepth int
	var realMaxLength int
	bst.Insert(10)
	expectedMaxDepth = 1
	realMaxLength = maxDepth(bst.Root)
	assert.Equal(t, expectedMaxDepth, realMaxLength)

	bst.Insert(8)
	expectedMaxDepth = 2
	realMaxLength = maxDepth(bst.Root)
	assert.Equal(t, expectedMaxDepth, realMaxLength)

	bst.Insert(7)
	expectedMaxDepth = 3
	realMaxLength = maxDepth(bst.Root)
	assert.Equal(t, expectedMaxDepth, realMaxLength)

	bst.Insert(9)
	expectedMaxDepth = 3
	realMaxLength = maxDepth(bst.Root)
	assert.Equal(t, expectedMaxDepth, realMaxLength)

	bst.Insert(20)
	expectedMaxDepth = 3
	realMaxLength = maxDepth(bst.Root)
	assert.Equal(t, expectedMaxDepth, realMaxLength)

	bst.Insert(25)
	expectedMaxDepth = 3
	realMaxLength = maxDepth(bst.Root)
	assert.Equal(t, expectedMaxDepth, realMaxLength)

	bst.Insert(24)
	expectedMaxDepth = 4
	realMaxLength = maxDepth(bst.Root)
	assert.Equal(t, expectedMaxDepth, realMaxLength)
	bst.Insert(29)

	expectedMaxDepth = 4
	realMaxLength = maxDepth(bst.Root)
	assert.Equal(t, expectedMaxDepth, realMaxLength)
}
