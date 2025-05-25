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
	"errors"
	"fmt"
	"strings"
)

// SinglyLinkedList represents a singly linked list that stores values of any type.
// It maintains pointers to the head and tail nodes, as well as the current size of
// the list.
type SinglyLinkedList struct {
	head *SinglyLinkedNode
	tail *SinglyLinkedNode
	size int
}

// NewSinglyLinkedList() creates and returns a new empty singly linked list.
//
// Returns:
//
//   - A pointer to the newly created SinglyLinkedList.
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// Head() returns the first node of the list or nil if the list is empty.
//
// Returns:
//   - A pointer to the head of the list.
func (l *SinglyLinkedList) Head() *SinglyLinkedNode {
	return l.head
}

// Tail() returns the last node of the list or nil if the list is empty.
//
// Returns:
//   - A pointer to the tail of the list.
func (l *SinglyLinkedList) Tail() *SinglyLinkedNode {
	return l.tail
}

// Size() returns the number of elements in the list.
//
// Returns:
//   - An integer representing the total number of elements.
func (l *SinglyLinkedList) Size() int {
	return l.size
}

// IsEmpty() checks if the list is empty.
//
// Returns:
//   - true if the list has no elements.
//   - false if the list has at least one element.
func (l *SinglyLinkedList) IsEmpty() bool {
	return l.Size() == 0
}

// Clear() removes all elements from the list.
func (l *SinglyLinkedList) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Prepend() inserts a new element at the beginning of the list.
//
// Parameters:
//   - data: The value to insert at the beginning of the list.
func (l *SinglyLinkedList) Prepend(data any) {
	newNode := NewSinglyLinkedNode(data)
	if l.IsEmpty() {
		l.tail = newNode
	} else {
		newNode.SetNext(l.Head())
	}
	l.head = newNode
	l.size++
}

// Append() inserts a new element at the end of the list.
//
// Parameters:
//   - data: The value to insert at the end of the list.
func (l *SinglyLinkedList) Append(data any) {
	newNode := NewSinglyLinkedNode(data)
	if l.IsEmpty() {
		l.head = newNode
	} else {
		l.Tail().SetNext(newNode)
	}
	l.tail = newNode
	l.size++
}

// Find() searches for the first node containing the specified data.
//
// Parameters:
//   - data: The value to search for.
//
// Returns:
//   - A pointer to the node containing the data, or nil if not found.
func (l *SinglyLinkedList) Find(data any) *SinglyLinkedNode {
	for current := l.Head(); current != nil; current = current.Next() {
		if current.Data() == data {
			return current
		}
	}
	return nil
}

// RemoveFirst() removes the first element from the list.
func (l *SinglyLinkedList) RemoveFirst() {
	if l.IsEmpty() {
		return
	}
	l.head = l.Head().Next()
	if l.Head() == nil {
		l.tail = nil
	}
	l.size--
}

// RemoveLast() removes the last element from the list.
func (l *SinglyLinkedList) RemoveLast() {
	if l.IsEmpty() {
		return
	}
	if l.Size() == 1 {
		l.head = nil
		l.tail = nil
	} else {
		current := l.Head()
		for current.Next() != l.Tail() {
			current = current.Next()
		}
		current.SetNext(nil)
		l.tail = current
	}
	l.size--
}

// Remove() deletes the first occurrence of the specified data from the list.
//
// Parameters:
//   - data: The value to remove from the list.
func (l *SinglyLinkedList) Remove(data any) {
	if l.IsEmpty() {
		return
	}
	if l.Head().Data() == data {
		l.RemoveFirst()
		return
	}
	prev := l.Head()
	current := l.Head().Next()
	for current != nil {
		if current.Data() == data {
			prev.SetNext(current.Next())
			if current == l.Tail() {
				l.tail = prev
			}
			l.size--
			return
		}
		prev = current
		current = current.Next()
	}
}

// String() returns a string representation of the list.
//
// Returns:
//   - A formatted string showing the sequence of node values.
func (l *SinglyLinkedList) String() string {
	if l.IsEmpty() {
		return "SinglyLinkedList: []"
	}
	var parts []string
	l.ForEach(func(value any) { parts = append(parts, fmt.Sprintf("[%v]", value)) })
	return "SinglyLinkedList: " + strings.Join(parts, " → ")
}

// ForEach() iterates over each element in the list and applies a given function.
//
// Parameters:
//   - f: A function that takes a value of type any. It is applied to each element
//     in the list in order.
func (l *SinglyLinkedList) ForEach(f func(any)) {
	for current := l.Head(); current != nil; current = current.Next() {
		f(current.Data())
	}
}

// InsertAt() inserts a new element at the specified index in the list.
//
// Parameters:
//   - index: The zero-based position where the new element will be inserted.
//   - data: The value to insert.
//
// Returns:
//   - An error occurs if the index is invalid, otherwise, nil.
func (l *SinglyLinkedList) InsertAt(index int, data any) error {
	if index < 0 || index > l.Size() {
		return errors.New("index out of bounds")
	}
	if index == 0 {
		l.Prepend(data)
		return nil
	}
	newNode := NewSinglyLinkedNode(data)
	prev := l.Head()
	for i := 1; i < index; i++ {
		prev = prev.Next()
	}
	newNode.SetNext(prev.Next())
	prev.SetNext(newNode)
	if newNode.Next() == nil {
		l.tail = newNode
	}
	l.size++
	return nil
}

// RemoveAll() removes all occurrences of the specified data from the list.
//
// Parameters:
//   - data: The value to remove from the list.
func (l *SinglyLinkedList) RemoveAll(data any) {
	if l.IsEmpty() {
		return
	}
	for l.Head() != nil && l.Head().Data() == data {
		l.RemoveFirst()
	}
	if l.IsEmpty() {
		return
	}
	prev := l.Head()
	current := l.Head().Next()
	for current != nil {
		if current.Data() == data {
			prev.SetNext(current.Next())
			l.size--
			if current == l.Tail() {
				l.tail = prev
			}
			current = prev.Next()
		} else {
			prev = current
			current = current.Next()
		}
	}
}

// Reverse() reverses the order of elements in the list.
func (l *SinglyLinkedList) Reverse() {
	var prev *SinglyLinkedNode
	current := l.Head()
	var next *SinglyLinkedNode
	for current != nil {
		next = current.Next()
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
	if l.Head() != nil {
		l.tail = l.Head()
		for l.Tail().Next() != nil {
			l.tail = l.Tail().Next()
		}
	}
}
