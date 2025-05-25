// Package priorityqueue provides a generic priority queue implementation using a
// binary heap.
//
// A priority queue stores elements along with associated priorities and allows
// efficient retrieval of the element with the highest or lowest priority,
// depending on configuration.
//
// Included features:
//   - Supports generic element types with priorities as integers.
//   - Allows creation of min-priority queues (lowest priority dequeued first).
//   - Allows creation of max-priority queues (highest priority dequeued first).
//   - Enqueue elements with a given priority.
//   - Dequeue elements with the current highest priority (min or max).
//   - Peek at the element with highest priority without removing it.
//   - Check if the queue is empty, get its size, or clear all elements.
//
// Internally, the priority queue uses a generic binary heap from the heap package,
// where elements are wrapped with their priorities for comparison.
package priorityqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPriorityQueueNewMinPriorityQueue() verifies that a new min priority queue is
// created empty.
func TestPriorityQueueNewMinPriorityQueue(t *testing.T) {
	pq := NewMinPriorityQueue[string]()
	assert.NotNil(t, pq)
	assert.True(t, pq.IsEmpty())
	assert.Equal(t, 0, pq.Size())
}

// TestPriorityQueueNewMaxPriorityQueue() verifies that a new max priority queue is
// created empty.
func TestPriorityQueueNewMaxPriorityQueue(t *testing.T) {
	pq := NewMaxPriorityQueue[string]()
	assert.NotNil(t, pq)
	assert.True(t, pq.IsEmpty())
	assert.Equal(t, 0, pq.Size())
}

// TestPriorityQueueEnqueueDequeueMin() verifies Enqueue and Dequeue operations on
// a min priority queue.
func TestPriorityQueueEnqueueDequeueMin(t *testing.T) {
	pq := NewMinPriorityQueue[string]()
	pq.Enqueue("low", 10)
	pq.Enqueue("medium", 5)
	pq.Enqueue("high", 1)
	assert.Equal(t, 3, pq.Size())
	assert.False(t, pq.IsEmpty())
	val, err := pq.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, "high", val)
	val, err = pq.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, "medium", val)
	val, err = pq.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, "low", val)
	assert.True(t, pq.IsEmpty())
	_, err = pq.Dequeue()
	assert.Error(t, err)
}

// TestPriorityQueueEnqueueDequeueMax() verifies Enqueue and Dequeue operations on
// a max priority queue.
func TestPriorityQueueEnqueueDequeueMax(t *testing.T) {
	pq := NewMaxPriorityQueue[string]()
	pq.Enqueue("low", 1)
	pq.Enqueue("medium", 5)
	pq.Enqueue("high", 10)
	assert.Equal(t, 3, pq.Size())
	assert.False(t, pq.IsEmpty())
	val, err := pq.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, "high", val)
	val, err = pq.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, "medium", val)
	val, err = pq.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, "low", val)
	assert.True(t, pq.IsEmpty())
	_, err = pq.Dequeue()
	assert.Error(t, err)
}

// TestPriorityQueuePeek() verifies the Peek method returns the element with
// highest priority without removing it.
func TestPriorityQueuePeek(t *testing.T) {
	pq := NewMinPriorityQueue[string]()
	pq.Enqueue("a", 2)
	pq.Enqueue("b", 1)
	val, err := pq.Peek()
	assert.NoError(t, err)
	assert.Equal(t, "b", val)
	sizeBefore := pq.Size()
	val2, err := pq.Peek()
	assert.NoError(t, err)
	assert.Equal(t, val, val2)
	assert.Equal(t, sizeBefore, pq.Size())
	_, err = pq.Dequeue()
	assert.NoError(t, err)
	_, err = pq.Dequeue()
	assert.NoError(t, err)
	_, err = pq.Peek()
	assert.Error(t, err)
}

// TestPriorityQueueClear() verifies that Clear empties the priority queue.
func TestPriorityQueueClear(t *testing.T) {
	pq := NewMaxPriorityQueue[string]()
	pq.Enqueue("x", 1)
	pq.Enqueue("y", 2)
	assert.Equal(t, 2, pq.Size())
	pq.Clear()
	assert.True(t, pq.IsEmpty())
	assert.Equal(t, 0, pq.Size())
	_, err := pq.Dequeue()
	assert.Error(t, err)
}
