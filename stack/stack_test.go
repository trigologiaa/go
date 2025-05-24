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
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStackNewStack() verifies that a newly created stack is not nil, is empty,
// and has a size zero.
func TestStackNewStack(t *testing.T) {
	s := NewStack[int]()
	assert.NotNil(t, s)
	assert.True(t, s.IsEmpty())
	assert.Equal(t, 0, s.Size())
}

// TestStackTop() verifies that Top() returns the most recently pushed element
// without removit it, and returns an error if the stack is empty.
func TestStackTop(t *testing.T) {
	s1 := NewStack[int]()
	s1.Push(100)
	v1, err := s1.Top()
	assert.Equal(t, 100, v1)
	assert.NoError(t, err)
	assert.Equal(t, 1, s1.Size())
	s2 := NewStack[float64]()
	s2.Push(1.1)
	s2.Push(2.2)
	v2, err := s2.Top()
	assert.Equal(t, 2.2, v2)
	assert.NoError(t, err)
	s3 := NewStack[string]()
	_, err = s3.Top()
	assert.Error(t, err)
}

// TestStackPushAndSize() checks that elements are correctly pushed onto the stack
// and that the size reflects the number of elements.
func TestStackPushAndSize(t *testing.T) {
	s := NewStack[string]()
	s.Push("a")
	s.Push("b")
	s.Push("c")
	assert.False(t, s.IsEmpty())
	assert.Equal(t, 3, s.Size())
}

// TestStackClear() verifies that calling Clear() removes all elements from the
// stack.
func TestStackClear(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Clear()
	assert.True(t, s.IsEmpty())
	assert.Equal(t, 0, s.Size())
}

// TestStackString() ensures that the string representation of the stack includes
// the elements and the "Stack:" prefix.
func TestStackString(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	str := s.String()
	assert.Contains(t, str, "1")
	assert.Contains(t, str, "2")
	assert.Contains(t, str, "Stack:")
}

// TestStackStringExact() checks that the exact string representation of the stack
// matches the expected format.
func TestStackStringExact(t *testing.T) {
	s := NewStack[int]()
	s.Push(10)
	s.Push(20)
	s.Push(30)
	expected := "Stack: [10 20 30]"
	actual := s.String()
	if actual != expected {
		t.Errorf("String() = %q; want %q", actual, expected)
	}
}

// TestStackTopDoesNotRemoveElement() ensures that calling Top() does not alter
// the stack or remove the top element.
func TestStackTopDoesNotRemoveElement(t *testing.T) {
	s := NewStack[int]()
	s.Push(5)
	s.Top()
	assert.Equal(t, 1, s.Size())
}

// TestStackClearEmptyStack() verifies that calling Clear() on an empty stack
// leaves it empty and without errors.
func TestStackClearEmptyStack(t *testing.T) {
	s := NewStack[int]()
	s.Clear()
	assert.True(t, s.IsEmpty())
}

// TestStackIsEmpty() checks that IsEmpty() returns true for an empty stack and
// false after pushing an element.
func TestStackIsEmpty(t *testing.T) {
	s := NewStack[int]()
	assert.True(t, s.IsEmpty())
	s.Push(1)
	assert.False(t, s.IsEmpty())
}

// TestStackPop() verifies that Top() removes and returns the top element in LIFO
// order, and returns an error when called on an empty stack.
func TestStackPop(t *testing.T) {
	s := NewStack[int]()
	s.Push(10)
	s.Push(20)
	val, err := s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 20, val)
	assert.Equal(t, 1, s.Size())
	val, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 10, val)
	assert.True(t, s.IsEmpty())
	_, err = s.Pop()
	assert.EqualError(t, err, "stack empty")
}

// TestStackTopAfterClear() ensures that calling Top() after clearing the stack
// returns an error.
func TestStackTopAfterClear(t *testing.T) {
	s := NewStack[int]()
	s.Push(42)
	s.Clear()
	_, err := s.Top()
	assert.Error(t, err)
}

// TestStackPopAfterClear() ensures that calling Pop() after clearing the stack
// returns an error.
func TestStackPopAfterClear(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Clear()
	_, err := s.Pop()
	assert.Error(t, err)
}

// TestStackWithStruct() verifies that the stack works with custom struct types
// and that Top() returns the correct value.
func TestStackWithStruct(t *testing.T) {
	type Point struct{ X, Y int }
	s := NewStack[Point]()
	s.Push(Point{1, 2})
	top, err := s.Top()
	assert.NoError(t, err)
	assert.Equal(t, Point{1, 2}, top)
}
