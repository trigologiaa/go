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

// TestNewSinglyLinkedNode() verifies that a newly created node has the expected
// data and its Next pointer is nil.
func TestNewSinglyLinkedNode(t *testing.T) {
	node := NewSinglyLinkedNode(10)
	assert.Equal(t, 10, node.Data(), "The value of the node should be 10")
	assert.Nil(t, node.Next(), "The next node should be nil")
}

// TestSetData() tests the SetData method, ensuring the data of a node can be
// updated correctly.
func TestSetData(t *testing.T) {
	node := NewSinglyLinkedNode(5)
	node.SetData(20)
	assert.Equal(t, 20, node.Data(), "The value of the node should be 20")
}

// TestSetNext() verifies that the SetNext method correctly updates the Next
// pointer of a node.
func TestSetNext(t *testing.T) {
	node1 := NewSinglyLinkedNode(10)
	node2 := NewSinglyLinkedNode(20)
	node1.SetNext(node2)
	assert.Equal(t, node2, node1.Next(), "The next node should be node2")
}

// TestHasNext() checks whether the HasNext method returns the correct value when a
// node has a next node or not.
func TestHasNext(t *testing.T) {
	node1 := NewSinglyLinkedNode(5)
	assert.False(t, node1.HasNext(), "node1 should not have next")
	node2 := NewSinglyLinkedNode(10)
	node1.SetNext(node2)
	assert.True(t, node1.HasNext(), "node1 should have the following")
}

// TestSetNextNil() tests the case when the Next pointer of a node is set to nil.
func TestSetNextNil(t *testing.T) {
	node := NewSinglyLinkedNode(100)
	node.SetNext(nil)
	assert.Nil(t, node.Next(), "The next node should be nil")
}

// TestHasNextWhenNextIsNil() ensures that HasNext returns false when a node has no
// next node.
func TestDataAfterSetNext(t *testing.T) {
	node1 := NewSinglyLinkedNode(1)
	node2 := NewSinglyLinkedNode(2)
	node1.SetNext(node2)
	assert.Equal(t, 1, node1.Data(), "The value of node1 should be 1")
	assert.Equal(t, 2, node1.Next().Data(), "The value of node2 should be 2")
}

// TestNewLinkedList() verifies that a newly created singly linked list is not nil,
// its size is 0, and both the head and tail are nil.
func TestHasNextWhenNextIsNil(t *testing.T) {
	node := NewSinglyLinkedNode(50)
	assert.False(t, node.HasNext(), "The node should not have next")
}
