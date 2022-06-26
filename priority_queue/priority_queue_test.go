package priorityqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	// 6 -> 5 -> 4 -> 2 -> 1 -> 0
	pq := NewPriorityQueue(false)

	pq.Enqueue(5, 5)
	assert.Equal(t, 5, pq.Front.Value)
	assert.Equal(t, 5, pq.Rear.Value)

	pq.Enqueue(1, 1)
	assert.Equal(t, 5, pq.Front.Value)
	assert.Equal(t, 1, pq.Rear.Value)

	pq.Enqueue(2, 2)
	assert.Equal(t, 5, pq.Front.Value)
	assert.Equal(t, 1, pq.Rear.Value)

	pq.Enqueue(6, 6)
	assert.Equal(t, 6, pq.Front.Value)
	assert.Equal(t, 1, pq.Rear.Value)

	pq.Enqueue(0, 0)
	assert.Equal(t, 6, pq.Front.Value)
	assert.Equal(t, 0, pq.Rear.Value)

	pq.Enqueue(4, 4)
	assert.Equal(t, 6, pq.Front.Value)
	assert.Equal(t, 0, pq.Rear.Value)

	// t.Log(pq.ToString())
}

func TestDeEnqueue(t *testing.T) {
	// 6 -> 5 -> 4 -> 2 -> 1 -> 0
	pq := NewPriorityQueue(false)

	var node *Node
	var ok bool

	node, ok = pq.Dequeue()
	assert.Nil(t, node)
	assert.False(t, ok)

	pq.Enqueue(5, 5)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 5)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Nil(t, node)
	assert.False(t, ok)

	pq.Enqueue(5, 5)
	pq.Enqueue(1, 1)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 5)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 1)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Nil(t, node)
	assert.False(t, ok)

	pq.Enqueue(5, 5)
	pq.Enqueue(1, 1)
	pq.Enqueue(2, 2)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 5)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 2)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 1)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Nil(t, node)
	assert.False(t, ok)

	pq.Enqueue(5, 5)
	pq.Enqueue(1, 1)
	pq.Enqueue(2, 2)
	pq.Enqueue(6, 6)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 6)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 5)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 2)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 1)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Nil(t, node)
	assert.False(t, ok)

	pq.Enqueue(5, 5)
	pq.Enqueue(1, 1)
	pq.Enqueue(2, 2)
	pq.Enqueue(6, 6)
	pq.Enqueue(0, 0)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 6)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 5)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 2)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 1)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 0)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Nil(t, node)
	assert.False(t, ok)

	pq.Enqueue(5, 5)
	pq.Enqueue(1, 1)
	pq.Enqueue(2, 2)
	pq.Enqueue(6, 6)
	pq.Enqueue(0, 0)
	pq.Enqueue(4, 4)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 6)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 5)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 4)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 2)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 1)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Equal(t, node.Value, 0)
	assert.True(t, ok)

	node, ok = pq.Dequeue()
	assert.Nil(t, node)
	assert.False(t, ok)

	// t.Log(pq.ToString())

}
