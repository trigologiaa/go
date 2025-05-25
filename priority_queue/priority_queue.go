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

import "github.com/trigologiaa/go/heap"

type prioritized[T any] struct {
	value    T
	priority int
}

// PriorityQueue[T any] represents a generic priority queue of elements of type T.
// The queue can be configured as a min-priority or max-priority queue.
type PriorityQueue[T any] struct {
	heap *heap.Heap[prioritized[T]]
}

// NewMinPriorityQueue() creates a new priority queue where elements with the
// lowest priority value are dequeued first.
//
// Returns:
//   - A pointer to an empty min-priority PriorityQueue.
func NewMinPriorityQueue[T any]() *PriorityQueue[T] {
	compare := func(a, b prioritized[T]) int {
		return a.priority - b.priority
	}
	return &PriorityQueue[T]{heap: heap.NewGenericHeap(compare)}
}

// NewMaxPriorityQueue() creates a new priority queue where elements with the
// highest priority value are dequeued first.
//
// Returns:
//   - A pointer to an empty max-priority PriorityQueue.
func NewMaxPriorityQueue[T any]() *PriorityQueue[T] {
	compare := func(a, b prioritized[T]) int {
		return b.priority - a.priority
	}
	return &PriorityQueue[T]{heap: heap.NewGenericHeap(compare)}
}

// Enqueue() inserts a new element with the specified priority into the priority
// queue.
//
// Parameters:
//   - value: The element to insert.
//   - priority: The priority associated with the element.
func (pq *PriorityQueue[T]) Enqueue(value T, priority int) {
	pq.heap.Insert(prioritized[T]{value: value, priority: priority})
}

// Dequeue() removes and returns the element with the highest priority (lowest for
// min-queue, highest for max-queue).
//
// Returns:
//   - The element with the highest priority.
//   - An error if the queue is empty.
func (pq *PriorityQueue[T]) Dequeue() (T, error) {
	item, err := pq.heap.Remove()
	if err != nil {
		var zero T
		return zero, err
	}
	return item.value, nil
}

// Peek() returns the element with the highest priority without removing it.
//
// Returns:
//   - The element with the highest priority.
//   - An error if the queue is empty.
func (pq *PriorityQueue[T]) Peek() (T, error) {
	item, err := pq.heap.Peek()
	if err != nil {
		var zero T
		return zero, err
	}
	return item.value, nil
}

// IsEmpty() returns true if the priority queue has no elements.
//
// Returns:
//   - true if the queue is empty.
//   - false if the queue is not empty.
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.Size() == 0
}

// Size() returns the number of elements currently in the priority queue.
//
// Returns:
//   - An integer representing the number of elements.
func (pq *PriorityQueue[T]) Size() int {
	return pq.heap.Size()
}

// Clear() removes all elements from the priority queue, resetting it to empty.
func (pq *PriorityQueue[T]) Clear() {
	pq.heap = heap.NewGenericHeap(pq.heap.Comparator())
}
