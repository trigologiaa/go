// Package singlylinkedlistany provides a generic implementation of a singly linked
// list data structure. It allows the storage of values of any type, providing
// operations for manipulating the list such as adding, removing, searching, and
// iterating over the elements.
//
// This package is useful for operations that require sequential data storage with
// dynamic size, such as constructing a list of items, appending or prepending
// elements, searching for specific values, or removing elements from a list.
//
// Included features:
//   - Create a new singly linked list.
//   - Add elements to the beginning or end of the list.
//   - Find elements in the list.
//   - Remove elements by value or by position.
//   - Check if the list is empty or retrieve the size of the list.
//   - Clear the entire list.
//   - Iterate over the elements in the list.
//   - Reverse the order of elements in the list.
//   - Get a string representation of the list.
//
// Most methods handle cases where the list is empty and return nil or no-op
// accordingly. Methods like 'InsertAt()' and 'RemoveAll()' ensure safe list
// manipulation without causing invalid memory access.
package singlylinkedlistany

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewLinkedList() verifies that a newly created singly linked list is not nil,
// its size is 0, and both the head and tail are nil.
func TestNewLinkedList(t *testing.T) {
	list := NewSinglyLinkedList()
	assert.NotNil(t, list)
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

// TestLinkedListPrependOnEmptyList() tests the Prepend method on an empty list,
// ensuring the element becomes both the head and tail of the list.
func TestLinkedListPrependOnEmptyList(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Prepend("1")
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())
}

// TestLinkedListPrependOnNonEmptyList() tests the Prepend method on a non-empty
// list, ensuring the element becomes the new head, and the tail remains the same.
func TestLinkedListPrependOnNonEmptyList(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append("1")
	list.Prepend("2")
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "2", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())
}

// TestLinkedListAppendOnEmptyList() tests the Append method on an empty list,
// ensuring the element becomes both the head and tail of the list.
func TestLinkedListAppendOnEmptyList(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append("1")
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "1", list.Tail().Data())
}

// TestLinkedListAppendOnNonEmptyList() tests the Append method on a non-empty
// list, ensuring the element is added to the tail of the list.
func TestLinkedListAppendOnNonEmptyList(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append("1")
	list.Append("2")
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "1", list.Head().Data())
	assert.Equal(t, "2", list.Tail().Data())
}

// TestLinkedListClear() tests the Clear method, ensuring it empties the list and
// resets both the head and tail to nil.
func TestLinkedListClear(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Clear()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

// TestLinkedListIsEmpty() tests the IsEmpty method, confirming it returns true for
// an empty list and false for a non-empty list.
func TestLinkedListIsEmpty(t *testing.T) {
	list := NewSinglyLinkedList()
	assert.True(t, list.IsEmpty())
	list.Append(1)
	assert.False(t, list.IsEmpty())
}

// TestLinkedListSize() tests the Size method, ensuring it accurately reports the
// number of elements in the list.
func TestLinkedListSize(t *testing.T) {
	list := NewSinglyLinkedList()
	assert.Equal(t, 0, list.Size())
	list.Append(1)
	assert.Equal(t, 1, list.Size())
	list.Append("Hello")
	assert.Equal(t, 2, list.Size())
}

// TestLinkedListHead() tests the Head method, verifying it returns the first
// element in the list.
func TestLinkedListHead(t *testing.T) {
	list := NewSinglyLinkedList()
	assert.Nil(t, list.Head())
	list.Append(1)
	assert.Equal(t, 1, list.Head().Data())
	list.Append("Hello")
	assert.Equal(t, 1, list.Head().Data())
}

// TestLinkedListTail() tests the Tail method, verifying it returns the last
// element in the list.
func TestLinkedListTail(t *testing.T) {
	list := NewSinglyLinkedList()
	assert.Nil(t, list.Tail())
	list.Append(1)
	assert.Equal(t, 1, list.Tail().Data())
	list.Append("Hello")
	assert.Equal(t, "Hello", list.Tail().Data())
}

// TestLinkedListFind() tests the Find method, confirming it returns the correct
// node containing the value and returns nil for a non-existent value.
func TestLinkedListFind(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(1)
	list.Append("Hello")
	list.Append(3.14)
	nodo := list.Find("Hello")
	assert.NotNil(t, nodo)
	assert.Equal(t, "Hello", nodo.Data())
	nodo = list.Find("NotFound")
	assert.Nil(t, nodo)
}

// TestLinkedListRemoveFirst() tests the RemoveFirst method, ensuring that it
// removes the first element and updates the head of the list accordingly.
func TestLinkedListRemoveFirst(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(1)
	list.Append("Hello")
	list.Append(3.14)
	list.RemoveFirst()
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "Hello", list.Head().Data())
	assert.Equal(t, 3.14, list.Tail().Data())
	list.RemoveFirst()
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 3.14, list.Head().Data())
	assert.Equal(t, 3.14, list.Tail().Data())
	list.RemoveFirst()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

// TestLinkedListRemoveLast() tests the RemoveLast method, ensuring it removes the
// last element and updates the tail of the list accordingly.
func TestLinkedListRemoveLast(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(1)
	list.Append("Hello")
	list.Append(3.14)
	list.RemoveLast()
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, "Hello", list.Tail().Data())
	list.RemoveLast()
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 1, list.Tail().Data())
	list.RemoveLast()
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

// TestLinkedListRemove() tests the Remove method, confirming that it removes the
// first occurrence of the specified element from the list.
func TestLinkedListRemove(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(1)
	list.Append("Hello")
	list.Append(3.14)
	list.Remove("Hello")
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 3.14, list.Tail().Data())
	list.Remove(1)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 3.14, list.Head().Data())
	assert.Equal(t, 3.14, list.Tail().Data())
	list.Remove(3.14)
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.Head())
	assert.Nil(t, list.Tail())
}

// TestLinkedListRemoveFirstOnEmpty() tests the RemoveFirst method on an empty
// list, confirming that it does nothing and the size remains 0.
func TestLinkedListRemoveFirstOnEmpty(t *testing.T) {
	list := NewSinglyLinkedList()
	list.RemoveFirst()
	assert.Equal(t, 0, list.Size())
}

// TestLinkedListRemoveLastOnEmpty() tests the RemoveLast method on an empty list,
// confirming that it does nothing and the size remains 0.
func TestLinkedListRemoveLastOnEmpty(t *testing.T) {
	list := NewSinglyLinkedList()
	list.RemoveLast()
	assert.Equal(t, 0, list.Size())
}

// TestLinkedListStringOnEmpty() tests the String method on an empty list, ensuring
// the string representation is correct.
func TestLinkedListStringOnEmpty(t *testing.T) {
	list := NewSinglyLinkedList()
	assert.Equal(t, "SinglyLinkedList: []", list.String())
}

// TestLinkedListString() tests the String method on a non-empty list, ensuring the
// string representation matches the elements in the list.
func TestLinkedListString(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(1)
	list.Append("Hello")
	list.Append(3.14)
	assert.Equal(t, "SinglyLinkedList: [1] → [Hello] → [3.14]", list.String())
}

// TestLinkedListInsertAt() tests the InsertAt method, verifying that it inserts an
// element at the specified index, shifting other elements accordingly.
func TestLinkedListInsertAt(t *testing.T) {
	list := NewSinglyLinkedList()
	list.InsertAt(0, 1)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	list.InsertAt(1, 3.14)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, 3.14, list.Tail().Data())
	list.InsertAt(1, "Hello")
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, "Hello", list.Head().Next().Data())
}

// TestLinkedListInsertAtOutOfBounds() tests the InsertAt method when an invalid
// index is provided (negative or out of bounds), ensuring that the list is not
// modified.
func TestLinkedListInsertAtOutOfBounds(t *testing.T) {
	list := NewSinglyLinkedList()
	list.InsertAt(-1, 1)
	assert.Equal(t, 0, list.Size())
	list.InsertAt(2, 1)
	assert.Equal(t, 0, list.Size())
}

// TestLinkedListRemoveAll() tests the RemoveAll method, ensuring that all
// instances of a specific value are removed from the list.
func TestLinkedListRemoveAll(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(1)
	list.Append("Hello")
	list.Append(1)
	list.Append(3.14)
	list.Append(1)
	list.RemoveAll(1)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, "Hello", list.Head().Data())
	assert.Equal(t, 3.14, list.Tail().Data())
}

// TestLinkedListRemoveAllNotFound() tests the RemoveAll method when the value to
// be removed is not found in the list, ensuring that the list remains unchanged.
func TestLinkedListRemoveAllNotFound(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(1)
	list.Append("Hello")
	list.RemoveAll("NotFound")
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.Head().Data())
	assert.Equal(t, "Hello", list.Tail().Data())
}

// TestLinkedListReverse() tests the Reverse method, ensuring that the order of
// elements is reversed.
func TestLinkedListReverse(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(1)
	list.Append("Hello")
	list.Append(3.14)
	list.Reverse()
	assert.Equal(t, 3.14, list.Head().Data())
	assert.Equal(t, 1, list.Tail().Data())
	assert.Equal(t, "SinglyLinkedList: [3.14] → [Hello] → [1]", list.String())
}

// TestLinkedListReverseEmptyAndSingle() tests the Reverse method on an empty list
// and a list with a single element, confirming no changes occur for these cases.
func TestLinkedListReverseEmptyAndSingle(t *testing.T) {
	empty := NewSinglyLinkedList()
	empty.Reverse()
	assert.Nil(t, empty.Head())
	assert.Nil(t, empty.Tail())
	single := NewSinglyLinkedList()
	single.Append(1)
	single.Reverse()
	assert.Equal(t, 1, single.Head().Data())
	assert.Equal(t, 1, single.Tail().Data())
}

// TestLinkedListForEach() tests the ForEach method, ensuring that it correctly
// applies a function to each element in the list.
func TestLinkedListForEach(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(1)
	list.Append("Hello")
	list.Append(3.14)
	var result []any
	list.ForEach(func(value any) { result = append(result, value) })
	assert.Equal(t, []any{1, "Hello", 3.14}, result)
}

// TestLinkedListLargeScaleAppend() tests appending a large number of elements
// (10000) to the list, ensuring the list's size and tail/head are correct.
func TestLinkedListLargeScaleAppend(t *testing.T) {
	list := NewSinglyLinkedList()
	for i := range 10000 {
		list.Append(i)
	}
	assert.Equal(t, 10000, list.Size())
	assert.Equal(t, 0, list.Head().Data())
	assert.Equal(t, 9999, list.Tail().Data())
}

// TestLinkedListNodeLinks() tests the links between nodes in the list, ensuring
// that the `Next` pointer for each node is correct.
func TestLinkedListNodeLinks(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(1)
	list.Append("Hello")
	list.Append(3.14)
	firstNode := list.Head()
	secondNode := firstNode.Next()
	thirdNode := secondNode.Next()
	assert.Equal(t, "Hello", secondNode.Data())
	assert.Equal(t, 3.14, thirdNode.Data())
	assert.Nil(t, thirdNode.Next())
}

// TestLinkedListInsertAtBounds() tests the InsertAt method for inserting an
// element at the bounds of the list (first and last position), ensuring it does
// not cause errors.
func TestLinkedListInsertAtBounds(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	err := list.InsertAt(0, "Start")
	assert.NoError(t, err)
	assert.Equal(t, "Start", list.Head().Data())
	err = list.InsertAt(4, "End")
	assert.NoError(t, err)
	assert.Equal(t, "End", list.Tail().Data())
}
