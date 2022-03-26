package single

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	ll := LinkedList{}
	assert.Nil(t, ll.head)
	// Insert to empty list

	ll.Insert(1)
	assert.Equal(t, ll.head.value, 1)
	assert.Nil(t, ll.head.next)

	// insert new
	ll.Insert(2)
	assert.Equal(t, ll.head.next.value, 2)
	assert.Nil(t, ll.head.next.next)

	ll.Insert(3)
	assert.Equal(t, ll.head.next.next.value, 3)
	assert.Nil(t, ll.head.next.next.next)
}

func TestInsertBefore(t *testing.T) {
	ll := LinkedList{}
	assert.Nil(t, ll.head)
	// insert to empty list
	ll.InsertBefore(1, nil)
	assert.Equal(t, ll.head.value, 1)
	assert.Nil(t, ll.head.next)

	// Insert before head. 2 --> 1
	ll.InsertBefore(2, ll.head)
	assert.Equal(t, ll.head.value, 2)
	assert.Equal(t, ll.head.next.value, 1)
	assert.Nil(t, ll.head.next.next)

	// Insert between 2 and 1. 2 --> 3 --> 1
	ll.InsertBefore(3, ll.head.next)
	assert.Equal(t, ll.head.value, 2)
	assert.Equal(t, ll.head.next.value, 3)
	assert.Equal(t, ll.head.next.next.value, 1)
	assert.Nil(t, ll.head.next.next.next)
}

//
func TestInsertAfter(t *testing.T) {
	ll := LinkedList{}
	assert.Nil(t, ll.head)
	// insert to empty list
	ll.InsertAfter(1, nil)
	assert.Equal(t, ll.head.value, 1)
	assert.Nil(t, ll.head.next)

	// Insert before head. 1 --> 2
	ll.InsertAfter(2, ll.head)
	assert.Equal(t, ll.head.value, 1)
	assert.Equal(t, ll.head.next.value, 2)
	assert.Nil(t, ll.head.next.next)

	// Insert after 2. 1 --> 2 --> 3
	ll.InsertAfter(3, ll.head.next)
	assert.Equal(t, ll.head.value, 1)
	assert.Equal(t, ll.head.next.value, 2)
	assert.Equal(t, ll.head.next.next.value, 3)
	assert.Nil(t, ll.head.next.next.next)

	// Insert after 2. 1 --> 2 --> 4 --> 3
	ll.InsertAfter(4, ll.head.next)
	assert.Equal(t, ll.head.value, 1)
	assert.Equal(t, ll.head.next.value, 2)
	assert.Equal(t, ll.head.next.next.value, 4)
	assert.Equal(t, ll.head.next.next.next.value, 3)
	assert.Nil(t, ll.head.next.next.next.next)

}

func TestDelete(t *testing.T) {
	ll1 := LinkedList{}
	assert.Nil(t, ll1.head)
	// delete head node
	ll1.Insert(1)
	ll1.Delete(ll1.head)
	assert.Nil(t, ll1.head)

	// delete node after head
	ll2 := LinkedList{}
	ll2.Insert(1)
	ll2.Insert(2)
	ll2.Delete(ll2.head)
	assert.Equal(t, ll2.head.value, 2)
	assert.Nil(t, ll2.head.next)

	ll3 := LinkedList{}
	ll3.Insert(1)
	ll3.Insert(2)
	ll3.Insert(3)
	ll3.Delete(ll3.head.next)
	assert.Equal(t, ll3.head.value, 1)
	assert.Equal(t, ll3.head.next.value, 3)
	assert.Nil(t, ll3.head.next.next)

	ll4 := LinkedList{}
	ll4.Insert(1)
	ll4.Insert(2)
	ll4.Insert(3)
	ll4.Insert(4)
	ll4.Insert(5)
	ll4.Delete(ll4.head.next)
	assert.Equal(t, ll4.head.value, 1)
	assert.Equal(t, ll4.head.next.value, 3)
	assert.Equal(t, ll4.head.next.next.value, 4)
	assert.Equal(t, ll4.head.next.next.next.value, 5)
	assert.Nil(t, ll4.head.next.next.next.next)
}

func TestRevert(t *testing.T) {
	var ll LinkedList
	ll = LinkedList{}
	ll.Revert()
	assert.Nil(t, ll.head)
	// delete head node

	ll = LinkedList{}
	ll.Insert(1)
	ll.Revert()
	assert.Equal(t, ll.head.value, 1)
	assert.Nil(t, ll.head.next)

	ll = LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Revert()
	assert.Equal(t, ll.head.value, 2)
	assert.Equal(t, ll.head.next.value, 1)
	assert.Nil(t, ll.head.next.next)

	ll = LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)
	ll.Revert()
	assert.Equal(t, ll.head.value, 4)
	assert.Equal(t, ll.head.next.value, 3)
	assert.Equal(t, ll.head.next.next.value, 2)
	assert.Equal(t, ll.head.next.next.next.value, 1)
	assert.Nil(t, ll.head.next.next.next.next)
}

func TestFindMiddleNode(t *testing.T) {
	var ll LinkedList
	var foundNode *Node

	ll = LinkedList{}
	foundNode = ll.FindMiddleNode()
	assert.Nil(t, foundNode)
	ll.Insert(1)
	foundNode = ll.FindMiddleNode()
	assert.Equal(t, foundNode.value, 1)
	ll.Insert(2)
	foundNode = ll.FindMiddleNode()
	assert.Equal(t, foundNode.value, 2)
	ll.Insert(3)
	foundNode = ll.FindMiddleNode()
	assert.Equal(t, foundNode.value, 2)
	ll.Insert(4)
	foundNode = ll.FindMiddleNode()
	assert.Equal(t, foundNode.value, 3)
	ll.Insert(5)
	foundNode = ll.FindMiddleNode()
	assert.Equal(t, foundNode.value, 3)
	ll.Insert(6)
	foundNode = ll.FindMiddleNode()
	assert.Equal(t, foundNode.value, 4)
	ll.Insert(7)
	foundNode = ll.FindMiddleNode()
	assert.Equal(t, foundNode.value, 4)
	ll.Insert(8)
	foundNode = ll.FindMiddleNode()
	assert.Equal(t, foundNode.value, 5)
	ll.Insert(9)
	foundNode = ll.FindMiddleNode()
	assert.Equal(t, foundNode.value, 5)
}
