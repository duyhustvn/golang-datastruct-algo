package priorityqueue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnqueue(t *testing.T) {

	pq := NewPriorityQueue()

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

	t.Log(pq.ToString())
}
