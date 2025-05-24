// Package stack provides a generic stack data structure implemented using Go
// generics. It allows storing and manipulating elements of any type (T)in a
// last-in, first-out (LIFO) manner.
//
// This package is useful in a wide range of applications such as expression
// evaluation, backtracking algorithms, depth-first search, undo mechanisms, and
// function call simulations.
//
// Included features:
//   - Push elements onto the stack.
//   - Pop elements from the top of the stack.
//   - Peek at the top element without removing it.
//   - Check if the stack is empty.
//   - Get the number of elements in the stack.
//   - Clear all elements from the stack.
//   - Get a string representation of the stack contents.
//
// Attempting to pop or peek from an empty stack will return an error.
package stack

import (
	"errors"
	"fmt"
)

// Stack[T any] represents a generic stack data structure that can store any type
// of data (T). The stack is implemented internally as a slice of type T.
type Stack[T any] struct {
	data []T
}

// NewStack[T any]() creates and returns a new, empty stack. The type of elements
// stored in the stack is generic (T).
//
// Returns:
//   - A pointer to a new empty stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: make([]T, 0)}
}

// Push() adds an element to the top of the stack. The element is appended to the
// end of the internal slice.
//
// Parameters:
//   - data: The element to be added to the stack.
func (s *Stack[T]) Push(data T) {
	s.data = append(s.data, data)
}

// Pop() removes and returns the element at the top of the stack. If the stack is
// empty, it returns an error and the zero value for the type T.
//
// Returns:
//   - The element of type T at the top of the stack.
//   - An error if the stack is empty.
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack empty")
	}
	index := len(s.data) - 1
	value := s.data[index]
	s.data = s.data[:index]
	return value, nil
}

// Top() returns the element at the top of the stack without removing it. If the
// stack is empty, it returns an error and the zero value for the type T.
//
// Returns:
//   - The element of type T at the top of the stack.
//   - An error if the stack is empty.
func (s *Stack[T]) Top() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack empty")
	}
	return s.data[len(s.data)-1], nil
}

// IsEmpty() checks if the stack is empty.
//
// Returns:
//   - true if the stack is empty.
//   - false if the stack contains at least one element.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Size() returns the number of elements currently in the stack.
//
// Returns:
//   - The number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(s.data)
}

// Clear() removes all elements from the stack and frees the memory previously
// occupied.
func (s *Stack[T]) Clear() {
	s.data = make([]T, 0)
}

// String() returns a string representation of the stack, which is useful for
// debugging purposes.
//
// Returns:
//   - A string representing the current elements in the stack.
func (s *Stack[T]) String() string {
	return fmt.Sprintf("Stack: %v", s.data)
}
