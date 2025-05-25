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

// SinglyLinkedNode[T comparable] represents a node in a singly linked list. It
// stores a value of type T and a pointer to the next node in the list.
type SinglyLinkedNode[T comparable] struct {
	data T
	next *SinglyLinkedNode[T]
}

// NewSinglyLinkedNode[T comparable](data) creates and returns a new node for a
// singly linked list with the provided data value. The next pointer is initialized
// to nil.
//
// Parameters:
//   - data: The value to store in the node.
//
// Returns:
//   - A pointer to a new SinglyLinkedNode[T comparable] containing the specified
//     data.
func NewSinglyLinkedNode[T comparable](data T) *SinglyLinkedNode[T] {
	return &SinglyLinkedNode[T]{data: data}
}

// SetData() sets the data stored in the node.
//
// Parameters:
//   - data: The new value to store in the node.
func (n *SinglyLinkedNode[T]) SetData(data T) {
	n.data = data
}

// Data() returns the value stored in the node.
//
// Returns:
//   - The data of type T stored in the node.
func (n *SinglyLinkedNode[T]) Data() T {
	return n.data
}

// SetNext() sets the next pointer of the current node to the specified node.
//
// Parameters:
//   - newNext: A pointer to the node that should follow the current node.
func (n *SinglyLinkedNode[T]) SetNext(newNext *SinglyLinkedNode[T]) {
	n.next = newNext
}

// Next() returns the next node in the list.
//
// Returns:
//   - A pointer to the next SinglyLinkedNode or nil if this is the last node.
func (n *SinglyLinkedNode[T]) Next() *SinglyLinkedNode[T] {
	return n.next
}

// HasNext() checks whether the current node has a next node.
//
// Returns:
//   - true if the next node is not nil.
//   - false if the next node is nil.
func (n *SinglyLinkedNode[T]) HasNext() bool {
	return n.next != nil
}
