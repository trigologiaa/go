// Package queue provides a generic queue data structure implemented using Go
// generics. It allows storing and manipulating elements of any type (T) in a
// first-in, first-out (FIFO) manner.
//
// This package is useful in a wide range of applications such as task scheduling,
// beadth-first search, buffering, and order processing systems.
//
// Included features:
//   - Enqueue elements to the back of the queue.
//   - Dequeue elements from the front of the queue.
//   - Peek at the front element without removing it.
//   - Check if the queue is empty.
//   - Get the number of elements in the queue.
//   - Clear all elements from the queue.
//   - Get a string representation of the queue contents.
//
// Attempting to dequeue or peek from an empty queue will return an error.
package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewQueue() verifies that a newly created queue is not nil and is empty.
func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	assert.NotNil(t, q)
	assert.True(t, q.IsEmpty())
}

// TestQueueEnqueue() checks that elements can be enqueued and the queue becomes
// non-empty.
func TestQueueEnqueue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	assert.False(t, q.IsEmpty())
}

// TestQueueDequeue() verifies that elements are dequeued in FIFO order and the
// queue becomes empty after all elements are removed.
func TestQueueDequeue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	v, _ := q.Dequeue()
	assert.Equal(t, 1, v)
	v, _ = q.Dequeue()
	assert.Equal(t, 2, v)
	assert.True(t, q.IsEmpty())
}

// TestQueueDequeueOnEmptyQueue() ensures that Dequeue() returns an error when
// called on an empty queue.
func TestQueueDequeueOnEmptyQueue(t *testing.T) {
	q := NewQueue[int]()
	_, err := q.Dequeue()
	assert.EqualError(t, err, "empty queue")
}

// TestQueueFrontOnEmptyQueue() ensures that Front() returns an error when called
// on an empty queue.

func TestQueueFrontOnEmptyQueue(t *testing.T) {
	q := NewQueue[int]()
	_, err := q.Front()
	assert.EqualError(t, err, "empty queue")
}

// TestQueueFront() verifies that Front() returns the front element without
// removing it and updates correctly after dequeuing.
func TestQueueFront(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	v, _ := q.Front()
	assert.Equal(t, 1, v)
	q.Dequeue()
	v, _ = q.Front()
	assert.Equal(t, 2, v)
	q.Dequeue()
	v, _ = q.Front()
	assert.Equal(t, 3, v)
}

// TestQueueSize() checks that Size() reflects the correct number of elements after
// enqueuing and dequeuing operations.
func TestQueueSize(t *testing.T) {
	q := NewQueue[string]()
	assert.Equal(t, 0, q.Size())
	q.Enqueue("a")
	q.Enqueue("b")
	assert.Equal(t, 2, q.Size())
	q.Dequeue()
	assert.Equal(t, 1, q.Size())
}

// TestQueueClear() verifies that calling Clear() removes all elements and resets
// the queue state.
func TestQueueClear(t *testing.T) {
	q := NewQueue[float64]()
	q.Enqueue(1.1)
	q.Enqueue(2.2)
	assert.False(t, q.IsEmpty())
	q.Clear()
	assert.True(t, q.IsEmpty())
	assert.Equal(t, 0, q.Size())
}

// TestQueueString() ensures that the string representation includes all enqueued
// elements.
func TestQueueString(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(10)
	q.Enqueue(20)
	str := q.String()
	assert.Contains(t, str, "10")
	assert.Contains(t, str, "20")
}

// TestQueueOperationsCombined() performs a sequence of enqueue, dequeue, and front
// operations to validate queue behavior and integrity.
func TestQueueOperationsCombined(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	v, _ := q.Front()
	assert.Equal(t, 1, v)
	v, _ = q.Dequeue()
	assert.Equal(t, 1, v)
	v, _ = q.Dequeue()
	assert.Equal(t, 2, v)
	assert.Equal(t, 1, q.Size())
}
