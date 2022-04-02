package single

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBetween(t *testing.T) {
	var ll LinkedList
	ll = LinkedList{}
	ll.reverseBetween(1, 1)
	assert.Nil(t, ll.head)

	ll = LinkedList{}
	ll.Insert(1)
	ll.reverseBetween(1, 1)
	assert.Equal(t, ll.head.value, 1)

	ll = LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.reverseBetween(1, 2)
	assert.Equal(t, ll.head.value, 2)
	assert.Equal(t, ll.head.next.value, 1)
	assert.Nil(t, ll.head.next.next)

	ll = LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)
	ll.Insert(5)
	ll.reverseBetween(2, 4)
	assert.Equal(t, ll.head.value, 1)
	assert.Equal(t, ll.head.next.value, 4)
	assert.Equal(t, ll.head.next.next.value, 3)
	assert.Equal(t, ll.head.next.next.next.value, 2)
	assert.Equal(t, ll.head.next.next.next.next.value, 5)

	ll = LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)
	ll.Insert(5)
	ll.reverseBetween(2, 5)
	assert.Equal(t, ll.head.value, 1)
	assert.Equal(t, ll.head.next.value, 5)
	assert.Equal(t, ll.head.next.next.value, 4)
	assert.Equal(t, ll.head.next.next.next.value, 3)
	assert.Equal(t, ll.head.next.next.next.next.value, 2)

	ll = LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)
	ll.Insert(5)
	ll.reverseBetween(1, 5)
	assert.Equal(t, ll.head.value, 5)
	assert.Equal(t, ll.head.next.value, 4)
	assert.Equal(t, ll.head.next.next.value, 3)
	assert.Equal(t, ll.head.next.next.next.value, 2)
	assert.Equal(t, ll.head.next.next.next.next.value, 1)
}
