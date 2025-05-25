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
	"errors"
	"fmt"
	"strings"
)

// SinglyLinkedList[T comparable] represents a singly linked list that stores
// values of a generic type T. It maintains pointers to the head and tail nodes, as
// well as the current size of the list.
type SinglyLinkedList[T comparable] struct {
	head *SinglyLinkedNode[T]
	tail *SinglyLinkedNode[T]
	size int
}

// NewSinglyLinkedList[T comparable]() creates and returns a new empty singly
// linked list.
//
// Returns:
//   - A pointer to the newly created SinglyLinkedList.
func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

// Head() returns the first node of the list or nil if the list is empty.
//
// Returns:
//   - A pointer to the head of the list.
func (l *SinglyLinkedList[T]) Head() *SinglyLinkedNode[T] {
	return l.head
}

// Tail() returns the last node of the list or nil if the list is empty.
//
// Returns:
//   - A pointer to the tail of the list.
func (l *SinglyLinkedList[T]) Tail() *SinglyLinkedNode[T] {
	return l.tail
}

// Size() returns the number of elements in the list.
//
// Returns:
//   - An integer representing the total number of elements.
func (l *SinglyLinkedList[T]) Size() int {
	return l.size
}

// IsEmpty() checks if the list is empty.
//
// Returns:
//   - true if the list has no elements.
//   - false if the list has at least one element.
func (l *SinglyLinkedList[T]) IsEmpty() bool {
	return l.Size() == 0
}

// Clear() removes all elements from the list.
func (l *SinglyLinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Prepend() inserts a new element at the beginning of the list.
//
// Parameters:
//   - data: The value to insert at the beginning of the list.
func (l *SinglyLinkedList[T]) Prepend(data T) {
	newNode := NewSinglyLinkedNode(data)
	if l.IsEmpty() {
		l.tail = newNode
	} else {
		newNode.SetNext(l.Head())
	}
	l.head = newNode
	l.size++
}

// Append inserts a new element at the end of the list.
//
// Parameters:
//   - data: The value to insert at the end of the list.
func (l *SinglyLinkedList[T]) Append(data T) {
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
func (l *SinglyLinkedList[T]) Find(data T) *SinglyLinkedNode[T] {
	for current := l.Head(); current != nil; current = current.Next() {
		if current.Data() == data {
			return current
		}
	}
	return nil
}

// RemoveFirst() removes the first element from the list.
func (l *SinglyLinkedList[T]) RemoveFirst() {
	if l.IsEmpty() {
		return
	}
	l.head = l.Head().Next()
	if l.head == nil {
		l.tail = nil
	}
	l.size--
}

// RemoveLast() removes the last element from the list.
func (l *SinglyLinkedList[T]) RemoveLast() {
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

// Remove deletes the first occurence of the specified data from the list.
//
// Parameters:
//   - data: The value to remove from the list.
func (l *SinglyLinkedList[T]) Remove(data T) {
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
func (l *SinglyLinkedList[T]) String() string {
	if l.IsEmpty() {
		return "SinglyLinkedList: []"
	}
	var parts []string
	l.ForEach(func(value T) { parts = append(parts, fmt.Sprintf("[%v]", value)) })
	return "SinglyLinkedList: " + strings.Join(parts, " → ")
}

// ForEach() iterates over each element in the list and appliesa given function.
//
// Parameters:
//   - f: A function that takes a value of type T. It is applied to each element in
//     the list in order.
func (l *SinglyLinkedList[T]) ForEach(f func(T)) {
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
func (l *SinglyLinkedList[T]) InsertAt(index int, data T) error {
	if index < 0 || index > l.Size() {
		return errors.New("index out of bounds")
	}
	if index == 0 {
		l.Prepend(data)
		return nil
	}
	if index == l.Size() {
		l.Append(data)
		return nil
	}
	newNode := NewSinglyLinkedNode(data)
	prev := l.Head()
	for i := 1; i < index; i++ {
		prev = prev.Next()
	}
	newNode.SetNext(prev.Next())
	prev.SetNext(newNode)
	l.size++
	return nil
}

// RemoveAll() removes all occurrences of the specified data from the list.
//
// Parameters:
//   - data: The value to remove from the list.
func (l *SinglyLinkedList[T]) RemoveAll(data T) {
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
func (l *SinglyLinkedList[T]) Reverse() {
	if l.IsEmpty() || l.Size() == 1 {
		return
	}
	var prev *SinglyLinkedNode[T]
	current := l.Head()
	l.tail = l.Head()
	for current != nil {
		next := current.Next()
		current.SetNext(prev)
		prev = current
		current = next
	}
	l.head = prev
}
