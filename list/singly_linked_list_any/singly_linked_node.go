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

// SinglyLinkedNode represents a node in a singly linked list. It stores a value of
// type any and a pointer to the next node in the list.
type SinglyLinkedNode struct {
	data any
	next *SinglyLinkedNode
}

// NewSinglyLinkedNode() creates and returns a new node for a singly linked list
// with the provided data value. The next pointer is initialized to nil.
//
// Parameters:
//   - data: The value to store in the node.
//
// Returns:
//   - A pointer to a new SinglyLinkedNode containing the specified data.
func NewSinglyLinkedNode(data any) *SinglyLinkedNode {
	return &SinglyLinkedNode{data: data}
}

// SetData() sets the data stored in the node.
//
// Parameters:
//   - data: The new value to store in the node.
func (n *SinglyLinkedNode) SetData(data any) {
	n.data = data
}

// Data() returns the value stored in the node.
//
// Returns:
//   - The data stored in the node.
func (n *SinglyLinkedNode) Data() any {
	return n.data
}

// SetNext() sets the next pointer of the current node to the specified node.
//
// Parameters:
//   - newNext: A pointer to the node that should follow the current node.
func (n *SinglyLinkedNode) SetNext(newNext *SinglyLinkedNode) {
	n.next = newNext
}

// Next() returns the next node in the list.
//
// Returns:
//   - A pointer to the next SinglyLinkedNode or nil if this is the last node.
func (n *SinglyLinkedNode) Next() *SinglyLinkedNode {
	return n.next
}

// HasNext() checks whether the current node has a next node.
//
// Returns:
//   - true if the next node is not nil.
//   - false if the next node is nil.
func (n *SinglyLinkedNode) HasNext() bool {
	return n.Next() != nil
}
