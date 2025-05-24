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
	"errors"
	"fmt"
)

// Queue[T any] represents a generic queue data structure that can store any type
// data (T). The queue is implemented internally as a slice of type T.
type Queue[T any] struct {
	data []T
}

// NewQueue[T any]() creates and returns a new empty queue. The type of elements
// stored in the queue is generic (T).
//
// Returns:
//   - A pointer to a new empty queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Enqueue() adds an element to the back of the queue. The element is appended to
// the end of the internal slice.
//
// Parameters:
//   - data: The element to be added to the queue.
func (q *Queue[T]) Enqueue(data T) {
	q.data = append(q.data, data)
}

// Dequeue() removes and returns the element at the front of the queue. If the
// queue is empty, it returns an error and the zero value for the type T.
//
// Returns:
//   - The element of type T at the front of the queue.
//   - An error if the queue is empty.
func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("empty queue")
	}
	head := q.data[0]
	q.data = q.data[1:]
	return head, nil
}

// Front() returns the element at the front of the queue without removing it. If
// the queue is empty, it returns an error and the zero value for the type T.
//
// Returns:
//   - The element of type T at the front of the queue.
//   - An error if the queue is empty.
func (q *Queue[T]) Front() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("empty queue")
	}
	head := q.data[0]
	return head, nil
}

// IsEmpty() checks if the queue is empty.
//
// Returns:
//   - true if the queue is empty.
//   - false if the queue contains at least one element.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

// Size() returns the number of elements currently in the queue.
//
// Returns:
//   - The number of elements in the queue.
func (q *Queue[T]) Size() int {
	return len(q.data)
}

// Clear() removes all elements from the queue and frees the memory previously
// occupied.
func (q *Queue[T]) Clear() {
	q.data = make([]T, 0)
}

// String() returns a string representation of the queue, which is useful for
// debugging purposes.
//
// Returns:
//   - A string representing the current elements in the queue.
func (q *Queue[T]) String() string {
	return fmt.Sprintf("Queue: %v", q.data)
}
