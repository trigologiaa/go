// Package singlylinkedlist provides a generic implementation of a singly linked
// list. A singly linked list is a linear data structure shere each element points
// to the next, with the last element pointing to nil.
//
// This package supports a variety of operations such as adding elements to the
// list, removing elements, and traversing through the list. It also provides
// functionality to check the size of the list and whether it is empty.
//
// Included features:
//   - Create a new singly linked list.
//   - Add elements to the beginning or end of the list.
//   - Search for an element.
//   - Remove elements by value or position.
//   - Check if the list is empty.
//   - Retrieve the first (head) and last (tail) elements.
//   - Reverse the list's order.
//   - Clear all elements from the list.
//   - Iterate over the list and apply a function to each element.
//   - Insert elements at a specified index.
//   - Remove all occurrences of a value from the list.
//
// Most methods handle cases where the list is empty and return nil or no-op
// accordingly. Methods like 'InsertAt()' and 'RemoveAll()' ensure safe list
// manipulation without causing invalid memory access.
package singlylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewSinglyLinkedNode() verifies that a newly created node has the expected
// data and its Next pointer is nil.
func TestNewSinglyLinkedNode(t *testing.T) {
	node := NewSinglyLinkedNode(42)
	assert.Equal(t, 42, node.Data(), "The value of the node should be 42")
	assert.Nil(t, node.Next(), "The next node should be nil")
}

// TestSetData() tests the SetData method, ensuring the data of a node can be
// updated correctly.
func TestSetData(t *testing.T) {
	node := NewSinglyLinkedNode(10)
	node.SetData(20)
	assert.Equal(t, 20, node.Data(), "The value of the node should be 20 after updating it")
}

// TestSetNextAndHasNext() verifies that the SetNext method correctly updates the
// next pointer of a node and that HasNext behaves as expected.
func TestSetNextAndHasNext(t *testing.T) {
	node1 := NewSinglyLinkedNode(1)
	node2 := NewSinglyLinkedNode(2)
	node1.SetNext(node2)
	assert.Equal(t, node2, node1.Next(), "The next node of node1 should be node2")
	assert.True(t, node1.HasNext(), "node1 should have a next node")
	assert.False(t, node2.HasNext(), "node2 should not have a next node")
}

// TestNextAndHasNextOnLastNode() ensures that Next() returns nil and HasNext()
// returns false for the last node in the list.
func TestNextAndHasNextOnLastNode(t *testing.T) {
	node := NewSinglyLinkedNode(100)
	assert.Nil(t, node.Next(), "The next node should be nil")
	assert.False(t, node.HasNext(), "The node should not have a next node")
}
