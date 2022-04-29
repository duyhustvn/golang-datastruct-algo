package path_sum

import (
	"testing"

	. "ds/tree/bst"

	"github.com/stretchr/testify/assert"
)

func TestMaxPathSum(t *testing.T) {
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
	assert.Equal(t, 10, maxPathSum(bst.Root))
	bst.Insert(8)
	assert.Equal(t, 18, maxPathSum(bst.Root))
	bst.Insert(20)
	assert.Equal(t, 30, maxPathSum(bst.Root))
	bst.Insert(7)
	assert.Equal(t, 30, maxPathSum(bst.Root))
	bst.Insert(9)
	assert.Equal(t, 30, maxPathSum(bst.Root))
	bst.Insert(25)
	assert.Equal(t, 55, maxPathSum(bst.Root))
	bst.Insert(24)
	assert.Equal(t, 79, maxPathSum(bst.Root))
	bst.Insert(29)
	assert.Equal(t, 84, maxPathSum(bst.Root))
}
