// Package heap provides a generic implementation of a binary heap data structure,
// supporting both min-heap and max-heap configurations.
//
// A heap is a complete binary tree where the value of each node is ordered with
// respect to its children according to a comparator function. This package allows
// storing elements of any type and defines custom behavior via a comparator
// function.
//
// Included features:
//   - Create a generic heap using a custom comparator.
//   - Create a min-heap or max-heap.
//   - Insert elements into the heap.
//   - Remove and return the root element (minimun or maximum depending oh the
//     heap).
//   - Retrieve the current size of the heap.
//   - Access the internal slice of elements for inspection or testing purposes.
//
// The implementation ensures the heap property is maintained on insertions and
// removals using up-heap and down-heap operations.
package heap

import "errors"

// Heap[T any] represents a generic binary heap that stores elements of type T. The
// ordering of elements is determined by the provided compare function.
type Heap[T any] struct {
	elements []T
	compare  func(a T, b T) int
}

// NewGenericHeap() creates and returns a new generic heap using the provided
// comparator function.
//
// Parameters:
//   - compare: A function that compares two elements. It should return:
//   - A negative value if a < b
//   - Zero if a == b
//   - A positive value if a > b
//
// Returns:
//   - A pointer to a new Heap instance.
func NewGenericHeap[T any](compare func(a T, b T) int) *Heap[T] {
	return &Heap[T]{compare: compare, elements: make([]T, 0)}
}

// NewMinHeap creates a new min-heap where the smallest element is at the root.
//
// Parameters:
//   - compare: A function that compares two elements. It should return:
//   - A negative value if a < b
//   - Zero if a == b
//   - A positive value if a > b
//
// Returns:
//   - A pointer to a new min-heap.
func NewMinHeap[T any](compare func(a T, b T) int) *Heap[T] {
	return &Heap[T]{compare: compare, elements: make([]T, 0)}
}

// NewMaxHeap creates a new max-heap where the largest element is at the root.
//
// Parameters:
//   - compare: A function that compares two elements. It should return:
//   - A negative value if a < b
//   - Zero if a == b
//   - A positive value if a > b
//
// Returns:
//   - A pointer to a new max-heap.
func NewMaxHeap[T any](compare func(a T, b T) int) *Heap[T] {
	comp := func(a T, b T) int {
		return compare(b, a)
	}
	return &Heap[T]{compare: comp, elements: make([]T, 0)}
}

// Size() returns the number of elements in the heap.
//
// Returns:
//   - An integer representing the number of elements.
func (h *Heap[T]) Size() int {
	return len(h.elements)
}

// Insert() adds a new element to the heap and restores the heap property.
//
// Parameters:
//   - element: The value to insert into the heap.
func (h *Heap[T]) Insert(element T) {
	h.elements = append(h.elements, element)
	h.upHeap(len(h.elements) - 1)
}

// Remove() removes and returns the root element (smallest or largest depending on
// heap type). It restores the heap property after removal.
//
// Returns:
//   - The removed element.
//   - An error if the heap is empty.
func (h *Heap[T]) Remove() (T, error) {
	var element T
	if h.Size() == 0 {
		return element, errors.New("empty heap")
	}
	element = h.elements[0]
	h.elements[0] = h.elements[h.Size()-1]
	h.elements = h.elements[:h.Size()-1]
	h.downHeap(0)
	return element, nil
}

// Elements() returns a slice containing all elements in the heap.
//
// Returns:
//   - A slice of elements currently in the heap.
func (h *Heap[T]) Elements() []T {
	return h.elements
}

// Peek() returns the root element of the heap without removing it.
//
// Returns:
//   - The element at the root of the heap (minimum or maximum depending on the
//     heap type).
//   - An error if the heap is empty.
func (h *Heap[T]) Peek() (T, error) {
	if h.Size() == 0 {
		var zero T
		return zero, errors.New("empty heap")
	}
	return h.elements[0], nil
}

// downHeap() restores the heap property by shifting an element down the tree from
// the given index.
//
// Parameters:
//   - i: The index of the element to sift down.
func (h *Heap[T]) downHeap(i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i
		if left < h.Size() && h.compare(h.elements[left], h.elements[smallest]) < 0 {
			smallest = left
		}
		if right < h.Size() && h.compare(h.elements[right], h.elements[smallest]) < 0 {
			smallest = right
		}
		if smallest == i {
			break
		}
		h.elements[i], h.elements[smallest] = h.elements[smallest], h.elements[i]
		i = smallest
	}
}

// upHeap() restores the heap property by shifting an element up the tree from the
// given index.
//
// Parameters:
//   - i: The index of the element to sift up.
func (h *Heap[T]) upHeap(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.compare(h.elements[i], h.elements[parent]) > 0 {
			break
		}
		h.elements[i], h.elements[parent] = h.elements[parent], h.elements[i]
		i = parent
	}
}

// Comparator() returns the comparison function used by the heap.
//
// Returns:
//   - A function that compares two elements and returns:
//   - A negative value if the first element is less than the second.
//   - Zero if both elements are equal.
//   - A positive value if the first element is greater than the second.
func (h *Heap[T]) Comparator() func(a, b T) int {
	return h.compare
}
